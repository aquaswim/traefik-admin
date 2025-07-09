import React from "react";
import { useParams, useNavigate } from "react-router";
import { Card, Alert } from "react-bootstrap";
import RouteForm from "../../components/forms/RouteForm";
import { useQGetRouteByID, useMUpdateRoute } from "../../lib/query.js";
import { useQueryClient } from "@tanstack/react-query";

function RouteEdit() {
  const { id } = useParams();
  const navigate = useNavigate();
  const queryClient = useQueryClient();

  const { data: route, isLoading: loading, error } = useQGetRouteByID(id);
  const { mutate: updateRoute, error: mutateError } =
    useMUpdateRoute(queryClient);

  const handleSubmit = (formData) => {
    updateRoute(
      { ...formData },
      {
        onSuccess: () => {
          navigate("/config/routers");
        },
      },
    );
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

  if (mutateError) {
    return <Alert variant="danger">{mutateError}</Alert>;
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
