import React from "react";
import { Button, Col, Label, Row } from "reactstrap";
import { FormInput } from "../../../forms/input";

export const AmountsEntry = arrayHelpers => {
  const { amounts } = arrayHelpers.form.values;
  return (
    <div>
      <Row>
        <Col xs="4">
          <Label>Coin ID</Label>
        </Col>
        <Col xs="5">
          <Label>Amount</Label>
        </Col>
        <Col xs="3"></Col>
      </Row>
      {amounts && amounts.length > 0 ? (
        amounts.map((amount, index) => {
          return (
            <Row key={index}>
              <Col xs="4">
                <FormInput name={`amounts.${index}.token_id`} />
              </Col>
              <Col xs="5">
                <FormInput name={`amounts.${index}.amount`} />
              </Col>
              <Col xs="3" className="d-flex flex-row-reverse">
                <div>
                  <Button
                    className="mr-1"
                    onClick={() => arrayHelpers.remove(index)}
                  >
                    <i className="icon-minus"></i>
                  </Button>
                  <Button
                    className=""
                    onClick={() =>
                      arrayHelpers.insert(index + 1, {
                        token_id: "",
                        amount: undefined
                      })
                    }
                  >
                    <i className="icon-plus"></i>
                  </Button>
                </div>
              </Col>
            </Row>
          );
        })
      ) : (
        <Button
          onClick={() => arrayHelpers.push({ token_id: "BNB", amount: "1" })}
        >
          <i className="icon-plus"></i>
        </Button>
      )}
    </div>
  );
};
