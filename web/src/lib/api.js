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
};

export default api;
