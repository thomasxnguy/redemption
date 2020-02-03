import React, { Component } from "react";
import { Col, Row, Alert } from "reactstrap";
import { CodesTable } from "../Shared/CodesTable";
import { httpClient } from "../../../utils/axios";
import ReactLoading from "react-loading";

class AllCodes extends Component {
  constructor(props) {
    super(props);
    this.state = {
      codes: [],
      loading: true,
      error: false
    };
  }

  async componentDidMount() {
    try {
      httpClient.defaults.headers.common[
        "Authorization"
      ] = `Bearer ${localStorage.getItem("bearerToken")}`;
      const resp = await httpClient.get("/v1/links");
      let data = resp.data;
      this.setState({ codes: data, loading: false });
    } catch (err) {
      this.setState({ loading: false, error: true });
    }
  }

  render() {
    const { loading, codes, error } = this.state;
    return (
      <div className="animated fadeIn">
        {error ? (
          <Alert color="danger">
            There was an error. Did you set the API_AUTH_TOKEN correctly?
          </Alert>
        ) : null}
        {!error && loading ? (
          <Row>
            <Col>
              <ReactLoading type="bars" color="#acb4bc" />
            </Col>
          </Row>
        ) : (
          <CodesTable codes={codes} />
        )}
      </div>
    );
  }
}

export default AllCodes;
