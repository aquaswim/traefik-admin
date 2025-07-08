import React from 'react';
import {Container} from "react-bootstrap";
import AppNavbar from "./AppNavbar.jsx";
import {Outlet} from "react-router";

function PanelLayout() {
    return (
        <>
            <AppNavbar />
            <Container as="main" className={"main"}>
                <Outlet/>
            </Container>
        </>
    );
}

export default PanelLayout;