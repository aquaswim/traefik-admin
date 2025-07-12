import React from "react";
import ConfigPreview from "../components/ConfigPreview.jsx";
import {Badge, Card, Col, Row, Spinner} from "react-bootstrap";
import {useQListRoutes, useQListServices} from "../lib/query.js";

function HomePage() {
    const {data: services, isLoading: isServiceLoading} = useQListServices()
    const {data: routes, isLoading: isRouteLoading} = useQListRoutes()
    const servicesCount = services.length;
    const routersCount = routes.length;

    return (
        <>
            <Row className="mb-4">
                <Col>
                    <Card className="text-center h-100">
                        <Card.Header>Services</Card.Header>
                        <Card.Body>
                            {isServiceLoading && (
                                <Spinner animation="border" role="status">
                                    <span className="visually-hidden">Loading...</span>
                                </Spinner>
                            )}
                            {!isServiceLoading && (
                                <Card.Title>
                                    <h1><Badge bg="primary">{servicesCount}</Badge></h1>
                                </Card.Title>
                            )}
                        </Card.Body>
                    </Card>
                </Col>
                <Col>
                    <Card className="text-center h-100">
                        <Card.Header>Routers</Card.Header>
                        <Card.Body>
                            {isRouteLoading && (
                                <Spinner animation="border" role="status">
                                    <span className="visually-hidden">Loading...</span>
                                </Spinner>
                            )}
                            {!isRouteLoading && (
                                <Card.Title>
                                    <h1><Badge bg="success">{routersCount}</Badge></h1>
                                </Card.Title>
                            )}
                        </Card.Body>
                    </Card>
                </Col>
            </Row>
            <Row>
                <Col>
                    <ConfigPreview/>
                </Col>
            </Row>
        </>
    );
}

export default HomePage;
