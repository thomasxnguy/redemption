import React, { useState } from "react";
import {
  Button,
  Card,
  CardBody,
  CardHeader,
  Col,
  Row,
  CardFooter,
  Alert
} from "reactstrap";
import { AmountsEntry } from "./AmountsEntry";
import { Formik, Form, FieldArray, useFormikContext } from "formik";
import { httpClient } from "../../../utils/axios";
import { CodesTable } from "../Shared/CodesTable";
import { FormInput } from "../../../forms/input";
import * as Yup from "yup";
import ReactLoading from "react-loading";

const TotalAmount = () => {
  const { values } = useFormikContext();
  const { quantity } = values;
  return values.amounts.map(({ token_id, amount }) => {
    return amount !== undefined ? (
      <div key={token_id} className="mb-2">
        <span>
          {amount * quantity}
          <small> {token_id}</small>
        </span>
      </div>
    ) : null;
  });
};

const CreateCodes = props => {
  const [codes, setCodes] = useState([]);
  const [loading, setLoading] = useState(false);
  const bearerToken = localStorage.getItem("bearerToken");
  const noAuthAPI = bearerToken == "" || bearerToken == null ? true : false;

  return (
    <div className="animated fadeIn">
      {noAuthAPI && (
        <Alert color="danger">
          There was an error. Did you set the API_AUTH_TOKEN correctly?
        </Alert>
      )}
      <Formik
        initialValues={{
          quantity: 1,
          amounts: [
            {
              token_id: "BNB",
              amount: 1
            },
            {
              token_id: "BUSD-BD1",
              amount: 1
            }
          ],
          coin: 714, // TODO(Dan): Replace with variable
          provider: window.location.host
        }}
        validationSchema={Yup.object({
          quantity: Yup.number()
            .positive("Minimum of 1 link")
            .required("Required"),
          amounts: Yup.array().of(
            Yup.object({
              token_id: Yup.string().required("Required"),
              amount: Yup.number()
                .typeError("Must be a number")
                .positive(1, "Minimum of 1 coin")
                .required("Required")
            })
          )
        })}
        onSubmit={async (values, { setSubmitting }) => {
          const { coin, provider, quantity, amounts } = values;
          setLoading(true);
          let resp = await httpClient.post("/v1/links/create", {
            asset: {
              coin,
              assets: amounts
            },
            link_count: quantity,
            provider
          });
          setCodes(resp.data);
          setLoading(false);
          setSubmitting(false);
        }}
      >
        <div>
          <Row>
            <Col xs="12" sm="8">
              <Form>
                <Card>
                  <CardHeader>
                    <strong>Number of Codes</strong>
                  </CardHeader>
                  <CardBody>
                    <Row>
                      <Col xs="12">
                        <FormInput
                          label="Quantity"
                          name="quantity"
                          type="number"
                        />
                      </Col>
                    </Row>
                  </CardBody>
                </Card>
                <Card>
                  <CardHeader>
                    <strong>Amount per Code</strong>
                  </CardHeader>
                  <CardBody>
                    <FieldArray
                      name="amounts"
                      validateOnChange={false}
                      render={AmountsEntry}
                    />
                  </CardBody>
                  <CardFooter>
                    <Button type="submit" size="sm" color="primary">
                      <i className="fa fa-gift"></i> Generate
                    </Button>
                  </CardFooter>
                </Card>
              </Form>
            </Col>
            <Col xs="12" sm="4">
              <Card>
                <CardHeader>Total</CardHeader>
                <CardBody>
                  <TotalAmount />
                </CardBody>
              </Card>
            </Col>
          </Row>
          {loading && <ReactLoading type="bars" color="#acb4bc" />}
          {codes.length !== 0 && <CodesTable codes={codes} />}
        </div>
      </Formik>
    </div>
  );
};

export default CreateCodes;
