import React, { useState, useEffect } from 'react';
import { Table, Button, Card, Stack } from 'react-bootstrap';
import { Link, useNavigate } from 'react-router';

// Mock data for demonstration
const mockServices = [
    {
        id: 'service1',
        servers: ['server1.example.com', 'server2.example.com'],
        type: 'http'
    },
    {
        id: 'service2',
        servers: ['192.168.1.1:8080'],
        type: 'tcp'
    },
    {
        id: 'service3',
        servers: ['udp-server.example.com:53'],
        type: 'udp'
    }
];

function ServiceList() {
    const [services, setServices] = useState([]);
    const navigate = useNavigate();

    // Simulate fetching data
    useEffect(() => {
        // In a real application, this would be an API call
        setServices(mockServices);
    }, []);

    const handleDelete = (id) => {
        // In a real application, this would be an API call
        if (window.confirm(`Are you sure you want to delete service "${id}"?`)) {
            setServices(services.filter(service => service.id !== id));
        }
    };

    return (
        <Card>
            <Card.Header>
                <Stack direction="horizontal" gap={3}>
                    <div className="me-auto">
                        <h4 className="mb-0">Services</h4>
                    </div>
                    <Link to="/config/services/new">
                        <Button variant="primary">Create New Service</Button>
                    </Link>
                </Stack>
            </Card.Header>
            <Card.Body>
                {services.length === 0 ? (
                    <p className="text-center">No services found.</p>
                ) : (
                    <Table striped bordered hover responsive>
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>Type</th>
                                <th>Servers</th>
                                <th>Actions</th>
                            </tr>
                        </thead>
                        <tbody>
                            {services.map((service) => (
                                <tr key={service.id}>
                                    <td>{service.id}</td>
                                    <td>{service.type.toUpperCase()}</td>
                                    <td>
                                        <ul className="mb-0 ps-3">
                                            {service.servers.map((server, index) => (
                                                <li key={index}>{server}</li>
                                            ))}
                                        </ul>
                                    </td>
                                    <td>
                                        <Stack direction="horizontal" gap={2}>
                                            <Button 
                                                variant="outline-primary" 
                                                size="sm"
                                                onClick={() => navigate(`/config/services/${service.id}/edit`)}
                                            >
                                                Edit
                                            </Button>
                                            <Button 
                                                variant="outline-danger" 
                                                size="sm"
                                                onClick={() => handleDelete(service.id)}
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

export default ServiceList;
