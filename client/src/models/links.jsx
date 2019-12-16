import React from "react";
import {
    List,
    Datagrid,
    TextField,
    BooleanField,
    DateField,
    UrlField,
    ArrayField,
    SingleFieldList,
    EditButton,
    TextInput,
    Edit,
    SimpleForm,
    BooleanInput,
    DateInput,
    NumberInput
} from "react-admin";

const ValueField = ({ record = {}, source }) => {
    console.log(source);
    return "hi";
};

export const LinkList = props => (
    <List {...props}>
        <Datagrid rowClick="edit">
            <BooleanField source="valid" />
            <BooleanField source="asset.used" label="Redeemed?" />
            <TextField source="code" />
            <TextField source="provider" />
            <DateField source="expiration_date" />
            <ArrayField source="asset.assets" label="Value">
                <SingleFieldList>
                    <ValueField source="asset.assets" />

                    {/* Create custom component to show amount and currency */}
                </SingleFieldList>
            </ArrayField>
            <UrlField source="link" />
            <EditButton />
        </Datagrid>
    </List>
);

export const LinkEdit = props => (
    <Edit {...props}>
        <SimpleForm>
            <TextInput source="link" />
            <TextInput source="code" />
            <TextInput source="provider" />
            <BooleanInput source="valid" />
            <DateInput source="expiration_date" />
            <DateInput source="created_date" />
            <NumberInput source="asset.coin" />
            {/* <TextInput source="id" /> */}
        </SimpleForm>
    </Edit>
);

export const LinkCreate = props => (
    <Edit {...props}>
        <SimpleForm>
            <TextInput source="link" />
            <TextInput source="code" />
            <TextInput source="provider" />
            <BooleanInput source="valid" />
            <DateInput source="expiration_date" />
            <DateInput source="created_date" />
            <NumberInput source="asset.coin" />
            {/* <TextInput source="id" /> */}
        </SimpleForm>
    </Edit>
);
