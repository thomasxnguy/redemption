import React, { Component } from "react";
import {
  Button,
  Card,
  CardBody,
  CardHeader,
  Col,
  Form,
  FormGroup,
  FormText,
  Input,
  Label,
  Row
} from "reactstrap";

class ViewCode extends Component {
  constructor(props) {
    super(props);

    this.toggle = this.toggle.bind(this);
    this.toggleFade = this.toggleFade.bind(this);
    this.state = {
      collapse: true,
      fadeIn: true,
      timeout: 300
    };
  }

  toggle() {
    this.setState({ collapse: !this.state.collapse });
  }

  toggleFade() {
    this.setState(prevState => {
      return { fadeIn: !prevState };
    });
  }

  render() {
    return (
      <div className="animated fadeIn">
        <Row>
          <Col xs="12" sm="8">
            <Form method="post" className="form-horizontal">
              <Card>
                <CardHeader>
                  <strong>Number of Codes</strong>
                </CardHeader>
                <CardBody>
                  <Row>
                    <Col xs="12">
                      <FormGroup>
                        <Input
                          type="text"
                          id="name"
                          placeholder="e.g. 10"
                          required
                        />
                      </FormGroup>
                    </Col>
                  </Row>
                </CardBody>
              </Card>
              <Card>
                <CardHeader>
                  <Button className="float-right" color="success" size="sm">
                    <i className="icon-plus"></i> Add
                  </Button>
                  <strong>Amount per Code</strong>
                </CardHeader>
                <CardBody>
                  <Row>
                    <Col xs="4">
                      <FormGroup>
                        <Label htmlFor="coin">Coin ID</Label>
                        <Input type="select" name="coin" id="coin">
                          <option>BNB</option>
                          <option>BUSD-BD1</option>
                        </Input>
                      </FormGroup>
                    </Col>
                    <Col xs="8">
                      <FormGroup>
                        <Label htmlFor="amount">Amount</Label>
                        <Input name="amount" id="amount"></Input>
                        <FormText className="help-block">
                          Equivalent: 1000000000
                        </FormText>
                      </FormGroup>
                    </Col>
                  </Row>
                </CardBody>
              </Card>
            </Form>
          </Col>
          <Col xs="12" sm="4">
            <Card>
              <CardHeader>
                <strong>Total</strong>
              </CardHeader>
              <CardBody></CardBody>
            </Card>
          </Col>
        </Row>
      </div>
    );
  }
}

export default ViewCode;
