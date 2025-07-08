import React from 'react';
import {Container, Nav, Navbar, NavDropdown} from "react-bootstrap";

function AppNavbar() {
    return (
        <Navbar expand="lg" className="bg-body-tertiary" fixed="top">
            <Container>
                <Navbar.Brand href="#/">Traefik Admin</Navbar.Brand>
                <Navbar.Toggle aria-controls="basic-navbar-nav" />
                <Navbar.Collapse id="basic-navbar-nav">
                    <Nav className="me-auto">
                        <Nav.Link href="#/stats">Stats</Nav.Link>
                        <NavDropdown title="Config" id="basic-nav-dropdown">
                            <NavDropdown.Item href="#/config/services">
                                Services
                            </NavDropdown.Item>
                            <NavDropdown.Item href="#/config/routers">Routers</NavDropdown.Item>
                            <NavDropdown.Divider />
                            <NavDropdown.Item href="#/config/middlewares">
                                Middleware
                            </NavDropdown.Item>
                        </NavDropdown>
                    </Nav>
                </Navbar.Collapse>
            </Container>
        </Navbar>
    );
}

export default AppNavbar;