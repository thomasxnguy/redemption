import React from "react";
import {
    List,
    Datagrid,
    TextField,
    BooleanField,
    DateField,
    ArrayField,
    SingleFieldList,
    EditButton,
    TextInput,
    Edit,
    SimpleForm,
    BooleanInput,
    DateInput,
    NumberInput,
    ShowButton
} from "react-admin";
import { Chip, Button } from "@material-ui/core";
import { FileCopy } from "@material-ui/icons";

const AssetField = ({ record = {}, source }) => {
    const asset = `${record.token_id} - ${record.amount}`;
    return <Chip label={asset} variant="outlined" color="secondary" />;
};

const CopyLinkField = ({ record = {}, source }) => {
    const { link } = record;

    const copyToClipboard = event => {
        navigator.clipboard.writeText(link).then(
            function() {
                console.log("Link copied to clipboard!");
            },
            function(err) {
                console.error("Error: could not copy link to clipboard ", err);
            }
        );
    };

    return (
        <Button
            variant="contained"
            color="primary"
            endIcon={<FileCopy />}
            onClick={copyToClipboard}
        >
            Copy
        </Button>
    );
};

export const LinkList = props => (
    <List {...props}>
        <Datagrid>
            <BooleanField source="valid" />
            <BooleanField source="asset.used" label="Redeemed?" />
            <TextField source="code" />
            <TextField source="provider" />
            <DateField source="expiration_date" />
            <ArrayField source="asset.assets" label="Value">
                <SingleFieldList>
                    <AssetField source="token_id" />
                </SingleFieldList>
            </ArrayField>
            <CopyLinkField source="link" />
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
