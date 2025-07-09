import React, { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router";
import { Alert, Card } from "react-bootstrap";
import ServiceForm from "../../components/forms/ServiceForm";
import api from "../../lib/api.js";
import { useMUpdateService } from "../../lib/query.js";
import { useQueryClient } from "@tanstack/react-query";

function ServiceEdit() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [service, setService] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const queryClient = useQueryClient();

  const { mutateAsync } = useMUpdateService(queryClient);

  useEffect(() => {
    const fetchData = async () => {
      try {
        setLoading(true);
        setError(null);
        const data = await api.getServiceByID(id);
        setService(data);
      } catch (err) {
        setError(err);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, [id]);

  const handleSubmit = (formData) => {
    setLoading(true);
    mutateAsync({ ...formData, id: id })
      .then(() => {
        navigate("/config/services");
      })
      .catch((err) => {
        setError(err);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  const handleCancel = () => {
    navigate("/config/services");
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
