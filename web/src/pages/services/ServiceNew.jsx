import React from "react";
import { useNavigate } from "react-router";
import { Card } from "react-bootstrap";
import ServiceForm from "../../components/forms/ServiceForm";
import { useMCreateService } from "../../lib/query.js";
import { useQueryClient } from "@tanstack/react-query";

function ServiceNew() {
  const navigate = useNavigate();

  const { mutate } = useMCreateService(useQueryClient());

  const handleSubmit = (formData) => {
    // In a real application, this would be an API call to create the service
    mutate(formData, {
      onSuccess: () => {
        navigate("/config/services");
      },
    });
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
          isEditForm={false}
        />
      </Card.Body>
    </Card>
  );
}

export default ServiceNew;
