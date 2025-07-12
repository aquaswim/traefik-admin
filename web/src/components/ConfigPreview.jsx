import React from 'react';
import {Accordion, Button} from "react-bootstrap";
import {useQTraefikConfig} from "../lib/query.js";

function ConfigPreview() {
    const { data: yamlConfig, refetch, isLoading } = useQTraefikConfig("yaml");

    return (
        <Accordion>
            <Accordion.Item eventKey="0">
                <Accordion.Header>Config Preview</Accordion.Header>
                <Accordion.Body>
                    {isLoading && <div>Loading...</div>}
                    <pre>
                        {yamlConfig}
                    </pre>
                    <Button onClick={() => refetch()}>Refresh</Button>
                </Accordion.Body>
            </Accordion.Item>
        </Accordion>
    );
}

export default ConfigPreview;