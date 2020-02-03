import React from "react";
import { Card, CardBody, CardHeader, Col, Row, Table } from "reactstrap";
import { SingleRow } from "../Shared/SingleRow";

export const CodesTable = ({ codes }) => {
  return (
    <div className="animated fadeIn">
      <Row>
        <Col>
          <Card>
            <CardHeader>
              <i className="fa fa-gift"></i>Links
            </CardHeader>
            <CardBody>
              <Table hover bordered responsive size="sm">
                <thead>
                  <tr>
                    <th>Valid</th>
                    <th>Status</th>
                    <th>Code</th>
                    <th>Value</th>
                    <th>Link</th>
                    <th>Details</th>
                  </tr>
                </thead>
                <tbody>
                  {codes &&
                    codes.map(code => (
                      <SingleRow key={code.code} code={code}></SingleRow>
                    ))}
                </tbody>
              </Table>
            </CardBody>
          </Card>
        </Col>
      </Row>
    </div>
  );
};
