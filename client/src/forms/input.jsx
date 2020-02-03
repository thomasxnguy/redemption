import React from "react";
import { FormGroup, FormFeedback, Input } from "reactstrap";
import { useField } from "formik";

export const FormInput = ({ label, ...props }) => {
  let [field, meta, helpers] = useField(props);
  const { setTouched } = helpers;
  const { onChange } = field;
  const { touched, error } = meta;
  field.onChange = (...args) => {
    setTouched(true);
    onChange(...args);
  };
  const inputInvalid = touched && error !== undefined;
  return (
    <FormGroup>
      {label ? <label htmlFor={props.id || props.name}>{label}</label> : null}
      <Input
        invalid={inputInvalid}
        className="text-input"
        onChange={args => {
          setTouched(true);
          field.onChange(...args);
        }}
        {...field}
        {...props}
      />
      {touched && error ? (
        <FormFeedback valid={!inputInvalid} className="error">
          {error}
        </FormFeedback>
      ) : null}
    </FormGroup>
  );
};
