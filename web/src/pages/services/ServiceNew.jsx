import React from "react";
import { useNavigate } from "react-router";
import { Card } from "react-bootstrap";
import ServiceForm from "../../components/forms/ServiceForm";

function ServiceNew() {
  const navigate = useNavigate();

  const handleSubmit = (formData) => {
    // In a real application, this would be an API call to create the service
    console.log("Creating new service:", formData);

    // Simulate successful creation
    alert(`Service "${formData.id}" created successfully!`);
    navigate("/config/services");
  };

  const handleCancel = () => {
    navigate("/config/services");
  };

  return (
    <Card>
      <Card.Header>
        <h4>Create New Service</h4>
      </Card.Header>
      <Card.Body>
        <ServiceForm
          initialValues={{
            id: "",
            servers: [""],
            type: "http",
          }}
          onSubmit={handleSubmit}
          onCancel={handleCancel}
        />
      </Card.Body>
    </Card>
  );
}

export default ServiceNew;
