import { useMutation, useQuery } from "@tanstack/react-query";
import api from "./api.js";

// Services queries and mutations
export const useQListServices = () =>
  useQuery({
    queryKey: ["services"],
    queryFn: () => {
      return api.getServices();
    },
    initialData: [],
  });

export const useQGetServicesByID = (id) =>
    useQuery({
        queryKey: ["services", id],
        queryFn: () => {
            return api.getServiceByID(id);
        },
        enabled: !!id,
    });

export const useMCreateService = (queryClient) =>
  useMutation({
    mutationFn: (data) => {
      return api.createService(data);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["services"] });
    },
  });

export const useMDeleteService = (queryClient) =>
  useMutation({
    mutationFn: async ({ id }) => await api.deleteServiceByID(id),
    onSettled: () => {
      queryClient.invalidateQueries({ queryKey: ["services"] });
    },
  });

export const useMUpdateService = (queryClient) =>
  useMutation({
    mutationFn: ({ id, ...data }) => {
      return api.updateServiceByID(id, data);
    },
    onSettled: (data) => {
      queryClient.invalidateQueries({ queryKey: ["services", data.id] });
    },
  });

// Routes queries and mutations
export const useQListRoutes = () =>
  useQuery({
    queryKey: ["routes"],
    queryFn: () => {
      return api.getRoutes();
    },
    initialData: [],
  });

export const useQGetRouteByID = (id) =>
  useQuery({
    queryKey: ["routes", id],
    queryFn: () => {
      return api.getRouteByID(id);
    },
    enabled: !!id,
  });

export const useMCreateRoute = (queryClient) =>
  useMutation({
    mutationFn: (data) => {
      return api.createRoute(data);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["routes"] });
    },
  });

export const useMDeleteRoute = (queryClient) =>
  useMutation({
    mutationFn: async ({ id }) => await api.deleteRouteByID(id),
    onSettled: () => {
      queryClient.invalidateQueries({ queryKey: ["routes"] });
    },
  });

export const useMUpdateRoute = (queryClient) =>
  useMutation({
    mutationFn: ({ id, ...data }) => {
      return api.updateRouteByID(id, data);
    },
    onSettled: (data) => {
      queryClient.invalidateQueries({ queryKey: ["routes", data.id] });
    },
  });
