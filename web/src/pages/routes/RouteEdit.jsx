import React, { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router";
import { Card, Alert } from "react-bootstrap";
import RouteForm from "../../components/forms/RouteForm";

// Mock data for demonstration - in a real app, this would come from an API
const mockRoutes = [
  {
    id: "route1",
    type: "http",
    rule: "Host(`example.com`) && Path(`/api`)",
    service: "service1",
  },
  {
    id: "route2",
    type: "tcp",
    rule: "HostSNI(`example.com`)",
    service: "service2",
  },
  {
    id: "route3",
    type: "udp",
    rule: "HostSNI(`udp-example.com`)",
    service: "service3",
  },
];

function RouteEdit() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [route, setRoute] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    // Simulate API call to fetch route by ID
    setLoading(true);
    setError(null);

    // In a real application, this would be an API call
    setTimeout(() => {
      const foundRoute = mockRoutes.find((r) => r.id === id);
      if (foundRoute) {
        setRoute(foundRoute);
        setLoading(false);
      } else {
        setError(`Route with ID "${id}" not found`);
        setLoading(false);
      }
    }, 500); // Simulate network delay
  }, [id]);

  const handleSubmit = (formData) => {
    // In a real application, this would be an API call to update the route
    console.log("Updating route:", formData);

    // Simulate successful update
    alert(`Route "${formData.id}" updated successfully!`);
    navigate("/config/routers");
  };

  const handleCancel = () => {
    navigate("/config/routers");
  };

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <Alert variant="danger">{error}</Alert>;
  }

  return (
    <Card>
      <Card.Header>
        <h4>Edit Route: {route.id}</h4>
      </Card.Header>
      <Card.Body>
        <RouteForm
          initialValues={route}
          onSubmit={handleSubmit}
          onCancel={handleCancel}
          isEditForm={true}
        />
      </Card.Body>
    </Card>
  );
}

export default RouteEdit;
