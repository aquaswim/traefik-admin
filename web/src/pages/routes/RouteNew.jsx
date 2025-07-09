import React from "react";
import { useNavigate } from "react-router";
import { Card } from "react-bootstrap";
import RouteForm from "../../components/forms/RouteForm";
import { useMCreateRoute } from "../../lib/query.js";
import { useQueryClient } from "@tanstack/react-query";

function RouteNew() {
  const navigate = useNavigate();

  const { mutate } = useMCreateRoute(useQueryClient());

  const handleSubmit = (formData) => {
    // Create the route using the API
    mutate(formData, {
      onSuccess: () => {
        navigate("/config/routers");
      },
    });
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
          isEditForm={false}
        />
      </Card.Body>
    </Card>
  );
}

export default RouteNew;
