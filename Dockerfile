FROM node:22-alpine AS builder-web
WORKDIR /app
COPY web/package.json web/package-lock.json ./
RUN npm ci
COPY web/ ./
RUN npm run build

FROM golang:1.24-alpine AS builder-go
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY --from=builder-web /app/dist/ web/dist/
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /traefik-admin ./cmd/traefik-admin

FROM alpine

RUN apk add --no-cache tzdata
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
RUN mkdir /config && chown appuser:appgroup /config
USER appuser

ENV APP_ADDRESS=:3000
ENV DB_PATH=/config

VOLUME /config

ENV TZ=Etc/UTC

COPY --from=builder-go /traefik-admin /usr/local/bin/traefik-admin

EXPOSE 3000

ENTRYPOINT ["traefik-admin"]