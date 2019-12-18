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
    NumberInput,
    ShowButton,
    ChipField
} from "react-admin";
import { Chip } from "@material-ui/core";

const AssetField = ({ record = {}, source }) => {
    const asset = `${record.token_id} - ${record.amount}`;
    return <Chip label={asset} variant="outlined" color="secondary" />;
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
                    <AssetField source="token_id" />
                    {/* TODO(Dan): Create custom component to show amount and currency */}
                </SingleFieldList>
            </ArrayField>
            <UrlField source="link" />
            <EditButton />
            <ShowButton />
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
