import React, { useState, useEffect } from "react";
import { Form, Button, Stack } from "react-bootstrap";

function RouteForm({
  initialValues = {},
  onSubmit,
  onCancel,
  isEditForm = false,
}) {
  const [formData, setFormData] = useState({
    id: "",
    type: "http",
    rule: "",
    service: "",
    ...initialValues,
  });

  // Reset form when initialValues change
  useEffect(() => {
    if (initialValues) {
      setFormData({
        id: "",
        type: "http",
        rule: "",
        service: "",
        ...initialValues,
      });
    }
  }, [initialValues]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit(formData);
  };

  return (
    <Form onSubmit={handleSubmit}>
      <Form.Group className="mb-3">
        <Form.Label>ID</Form.Label>
        <Form.Control
          type="text"
          name="id"
          value={formData.id}
          onChange={handleChange}
          required
          disabled={isEditForm}
        />
      </Form.Group>

      <Form.Group className="mb-3">
        <Form.Label>Type</Form.Label>
        <Form.Select
          name="type"
          value={formData.type}
          onChange={handleChange}
          required
        >
          <option value="http">HTTP</option>
          <option value="tcp">TCP</option>
          <option value="udp">UDP</option>
        </Form.Select>
      </Form.Group>

      <Form.Group className="mb-3">
        <Form.Label>Rule</Form.Label>
        <Form.Control
          type="text"
          name="rule"
          value={formData.rule}
          onChange={handleChange}
          required
        />
        <Form.Text className="text-muted">
          For HTTP: "Host(`example.com`) && Path(`/api`)" | For TCP/UDP:
          "HostSNI(`example.com`)"
        </Form.Text>
      </Form.Group>

      <Form.Group className="mb-3">
        <Form.Label>Service</Form.Label>
        <Form.Control
          type="text"
          name="service"
          value={formData.service}
          onChange={handleChange}
          required
        />
        <Form.Text className="text-muted">
          The name of the service this route should direct to
        </Form.Text>
      </Form.Group>

      <Stack direction="horizontal" gap={2} className="mt-4">
        <Button variant="primary" type="submit">
          Save
        </Button>
        {onCancel && (
          <Button variant="secondary" onClick={onCancel}>
            Cancel
          </Button>
        )}
      </Stack>
    </Form>
  );
}

export default RouteForm;
