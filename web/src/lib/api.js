const GET = "GET";
const POST = "POST";
const PUT = "PUT";
const DELETE = "DELETE";

const api = {
  getServices: () =>
    fetch("/api/services", { method: GET }).then((res) => res.json()),
  createService: (data) =>
    fetch("/api/services", {
      method: POST,
      body: JSON.stringify(data),
      headers: { "Content-Type": "application/json" },
    }).then((res) => res.json()),
  getServiceByID: (id) =>
    fetch(`/api/services/${id}`, { method: GET }).then((res) => res.json()),
  updateServiceByID: (id, data) =>
    fetch(`/api/services/${id}`, {
      method: PUT,
      body: JSON.stringify(data),
      headers: { "Content-Type": "application/json" },
    }).then((res) => res.json()),
  deleteServiceByID: (id) =>
    fetch(`/api/services/${id}`, { method: DELETE }).then((res) => res.json()),

  // Routes API
  getRoutes: () =>
    fetch("/api/routes", { method: GET }).then((res) => res.json()),
  createRoute: (data) =>
    fetch("/api/routes", {
      method: POST,
      body: JSON.stringify(data),
      headers: { "Content-Type": "application/json" },
    }).then((res) => res.json()),
  getRouteByID: (id) =>
    fetch(`/api/routes/${id}`, { method: GET }).then((res) => res.json()),
  updateRouteByID: (id, data) =>
    fetch(`/api/routes/${id}`, {
      method: PUT,
      body: JSON.stringify(data),
      headers: { "Content-Type": "application/json" },
    }).then((res) => res.json()),
  deleteRouteByID: (id) =>
    fetch(`/api/routes/${id}`, { method: DELETE }).then((res) => res.json()),

  getTraefikConfig: (type) =>
    fetch(`/api/traefik-config/${type}`, { method: GET }).then((res) =>
      res.text(),
    ),
};

export default api;
