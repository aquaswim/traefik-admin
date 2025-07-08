import React from "react";
import { useNavigate } from "react-router";
import { Card } from "react-bootstrap";
import RouteForm from "../../components/forms/RouteForm";

function RouteNew() {
  const navigate = useNavigate();

  const handleSubmit = (formData) => {
    // In a real application, this would be an API call to create the route
    console.log("Creating new route:", formData);

    // Simulate successful creation
    alert(`Route "${formData.id}" created successfully!`);
    navigate("/config/routers");
  };

  const handleCancel = () => {
    navigate("/config/routers");
  };

  return (
    <Card>
      <Card.Header>
        <h4>Create New Route</h4>
      </Card.Header>
      <Card.Body>
        <RouteForm
          initialValues={{
            id: "",
            type: "http",
            rule: "",
            service: "",
          }}
          onSubmit={handleSubmit}
          onCancel={handleCancel}
        />
      </Card.Body>
    </Card>
  );
}

export default RouteNew;
