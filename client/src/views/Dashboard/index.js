import React, { Component } from "react";
import {
  Card,
  CardBody,
  CardHeader,
  Col,
  Row,
  Form,
  FormGroup,
  FormText,
  Input,
  CardFooter,
  Button
} from "reactstrap";

class Dashboard extends Component {
  constructor(props) {
    super(props);
    this.state = {
      bearerToken: ""
    };
  }

  componentDidMount() {
    let bearerToken = localStorage.getItem("bearerToken");
    if (bearerToken == null) bearerToken = "";
    this.setState({ bearerToken });
  }

  handleChange = event => {
    this.setState({ bearerToken: event.target.value });
  };

  handleSubmit = event => {
    event.preventDefault();
    const { bearerToken } = this.state;
    localStorage.setItem("bearerToken", bearerToken);
  };

  handleClear = event => {
    event.preventDefault();
    localStorage.setItem("bearerToken", "");
    this.setState({ bearerToken: "" });
  };

  render() {
    const { bearerToken } = this.state;
    return (
      <div className="animated fadeIn">
        <Row>
          <Col xs="12" sm="12" lg="12">
            <Card>
              <CardHeader>
                <strong>API Auth Token</strong>
                <small> See info</small>
              </CardHeader>
              <Form onSubmit={this.handleSubmit} className="was-validated">
                <CardBody className="pb-0">
                  <FormGroup>
                    <Input
                      type="text"
                      id="apiAuthToken"
                      onChange={this.handleChange}
                      value={bearerToken}
                      placeholder="API Auth Token"
                      required
                    />
                    <FormText>
                      See your{" "}
                      <a href="https://devcenter.heroku.com/articles/config-vars">
                        Heroku Config Vars
                      </a>{" "}
                      for AUTH_API_TOKEN
                    </FormText>
                  </FormGroup>
                </CardBody>
                <CardFooter>
                  <Button type="submit" size="sm" color="primary">
                    <i className="fa fa-dot-circle-o"></i> Save to LocalStorage
                  </Button>
                  <Button
                    type="reset"
                    onClick={this.handleClear}
                    size="sm"
                    color="danger"
                  >
                    <i className="fa fa-ban"></i> Clear
                  </Button>
                </CardFooter>
              </Form>
            </Card>
          </Col>
        </Row>
      </div>
    );
  }
}

export default Dashboard;
