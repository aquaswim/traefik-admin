import React, { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router";
import { Card, Alert } from "react-bootstrap";
import ServiceForm from "../../components/forms/ServiceForm";

// Mock data for demonstration - in a real app, this would come from an API
const mockServices = [
  {
    id: "service1",
    servers: ["server1.example.com", "server2.example.com"],
    type: "http",
  },
  {
    id: "service2",
    servers: ["192.168.1.1:8080"],
    type: "tcp",
  },
  {
    id: "service3",
    servers: ["udp-server.example.com:53"],
    type: "udp",
  },
];

function ServiceEdit() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [service, setService] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    // Simulate API call to fetch service by ID
    setLoading(true);
    setError(null);

    // In a real application, this would be an API call
    setTimeout(() => {
      const foundService = mockServices.find((s) => s.id === id);
      if (foundService) {
        setService(foundService);
        setLoading(false);
      } else {
        setError(`Service with ID "${id}" not found`);
        setLoading(false);
      }
    }, 500); // Simulate network delay
  }, [id]);

  const handleSubmit = (formData) => {
    // In a real application, this would be an API call to update the service
    console.log("Updating service:", formData);

    // Simulate successful update
    alert(`Service "${formData.id}" updated successfully!`);
    navigate("/config/services");
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
