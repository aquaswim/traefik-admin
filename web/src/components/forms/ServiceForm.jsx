import React, { useState, useEffect } from "react";
import { Form, Button, Row, Col, Stack } from "react-bootstrap";

function ServiceForm({
  initialValues = {},
  onSubmit,
  onCancel,
  isEditForm = false,
}) {
  const [formData, setFormData] = useState({
    id: "",
    servers: [""],
    type: "http",
    ...initialValues,
  });

  // Reset form when initialValues change
  useEffect(() => {
    if (initialValues) {
      setFormData({
        id: "",
        servers: [""],
        type: "http",
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

  const handleServerChange = (index, value) => {
    const updatedServers = [...formData.servers];
    updatedServers[index] = value;
    setFormData({
      ...formData,
      servers: updatedServers,
    });
  };

  const addServer = () => {
    setFormData({
      ...formData,
      servers: [...formData.servers, ""],
    });
  };

  const removeServer = (index) => {
    const updatedServers = [...formData.servers];
    updatedServers.splice(index, 1);
    setFormData({
      ...formData,
      servers: updatedServers.length ? updatedServers : [""],
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
        <Form.Label>Servers</Form.Label>
        {formData.servers.map((server, index) => (
          <Row key={index} className="mb-2">
            <Col>
              <Form.Control
                type="text"
                value={server}
                onChange={(e) => handleServerChange(index, e.target.value)}
                required
              />
            </Col>
            <Col xs="auto">
              <Button
                variant="danger"
                onClick={() => removeServer(index)}
                disabled={formData.servers.length === 1}
              >
                Remove
              </Button>
            </Col>
          </Row>
        ))}
        <Button variant="secondary" onClick={addServer} className="mt-2">
          Add Server
        </Button>
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

export default ServiceForm;
