import React from "react";
import { useQTraefikConfig } from "../lib/query.js";
import { Button, Card, Col, Row } from "react-bootstrap";

function HomePage() {
  const { data: yamlConfig, refetch, isLoading } = useQTraefikConfig("yaml");

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <Card>
      <Card.Header>
        <h2>Current Config</h2>
      </Card.Header>
      <Card.Body>
        <pre>{yamlConfig}</pre>
      </Card.Body>
      <Card.Footer>
        <Button onClick={() => refetch()}>Refresh</Button>
      </Card.Footer>
    </Card>
  );
}

export default HomePage;
