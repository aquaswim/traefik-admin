import React, { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router";
import { Alert, Card } from "react-bootstrap";
import ServiceForm from "../../components/forms/ServiceForm";
import api from "../../lib/api.js";
import {useMUpdateService, useQGetServicesByID} from "../../lib/query.js";
import { useQueryClient } from "@tanstack/react-query";

function ServiceEdit() {
  const { id } = useParams();
  const navigate = useNavigate();

  const queryClient = useQueryClient();

  const {data: service, isLoading: fetchLoading, error: fetchError} = useQGetServicesByID(id)
  const { mutate, error: mutateError } = useMUpdateService(queryClient);

  const handleSubmit = (formData) => {
    mutate({ ...formData, id: id }, {
      onSuccess: () => {
        navigate("/config/services");
      },
    });
  };

  const handleCancel = () => {
    navigate("/config/services");
  };

  if (fetchLoading) {
    return <div>Loading...</div>;
  }

  if (fetchError) {
    return <Alert variant="danger">{fetchError}</Alert>;
  }

  if (mutateError) {
    return <Alert variant="danger">{mutateError}</Alert>;
  }

  return (
    <Card>
      <Card.Header>
        <h4>Edit Service: {service.id}</h4>
      </Card.Header>
      <Card.Body>
        <ServiceForm
          initialValues={service}
          onSubmit={handleSubmit}
          onCancel={handleCancel}
          isEditForm={true}
        />
      </Card.Body>
    </Card>
  );
}

export default ServiceEdit;
