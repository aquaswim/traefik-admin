import { useMutation, useQuery } from "@tanstack/react-query";
import api from "./api.js";

export const useQListServices = () =>
  useQuery({
    queryKey: ["services"],
    queryFn: () => {
      return api.getServices();
    },
    initialData: [],
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
