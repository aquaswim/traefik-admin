import React, { useState, useEffect } from 'react';
import { Table, Button, Card, Stack } from 'react-bootstrap';
import { Link, useNavigate } from 'react-router';

// Mock data for demonstration
const mockRoutes = [
    {
        id: 'route1',
        type: 'http',
        rule: 'Host(`example.com`) && Path(`/api`)',
        service: 'service1'
    },
    {
        id: 'route2',
        type: 'tcp',
        rule: 'HostSNI(`example.com`)',
        service: 'service2'
    },
    {
        id: 'route3',
        type: 'udp',
        rule: 'HostSNI(`udp-example.com`)',
        service: 'service3'
    }
];

function RouteList() {
    const [routes, setRoutes] = useState([]);
    const navigate = useNavigate();

    // Simulate fetching data
    useEffect(() => {
        // In a real application, this would be an API call
        setRoutes(mockRoutes);
    }, []);

    const handleDelete = (id) => {
        // In a real application, this would be an API call
        if (window.confirm(`Are you sure you want to delete route "${id}"?`)) {
            setRoutes(routes.filter(route => route.id !== id));
        }
    };

    return (
        <Card>
            <Card.Header>
                <Stack direction="horizontal" gap={3}>
                    <div className="me-auto">
                        <h4 className="mb-0">Routes</h4>
                    </div>
                    <Link to="/config/routers/new">
                        <Button variant="primary">Create New Route</Button>
                    </Link>
                </Stack>
            </Card.Header>
            <Card.Body>
                {routes.length === 0 ? (
                    <p className="text-center">No routes found.</p>
                ) : (
                    <Table striped bordered hover responsive>
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>Type</th>
                                <th>Rule</th>
                                <th>Service</th>
                                <th>Actions</th>
                            </tr>
                        </thead>
                        <tbody>
                            {routes.map((route) => (
                                <tr key={route.id}>
                                    <td>{route.id}</td>
                                    <td>{route.type.toUpperCase()}</td>
                                    <td>{route.rule}</td>
                                    <td>{route.service}</td>
                                    <td>
                                        <Stack direction="horizontal" gap={2}>
                                            <Button 
                                                variant="outline-primary" 
                                                size="sm"
                                                onClick={() => navigate(`/config/routers/${route.id}/edit`)}
                                            >
                                                Edit
                                            </Button>
                                            <Button 
                                                variant="outline-danger" 
                                                size="sm"
                                                onClick={() => handleDelete(route.id)}
                                            >
                                                Delete
                                            </Button>
                                        </Stack>
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </Table>
                )}
            </Card.Body>
        </Card>
    );
}

export default RouteList;
