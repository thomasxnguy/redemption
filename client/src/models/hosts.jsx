import React from "react";
import {
    List,
    Datagrid,
    TextField,
    SimpleForm,
    TextInput,
    NumberInput,
    Create
} from "react-admin";

export const HostList = props => (
    <List {...props}>
        <Datagrid>
            <TextField source="id" label="coin" />
            <TextField source="host" />
        </Datagrid>
    </List>
);

export const HostCreate = props => (
    <Create {...props}>
        <SimpleForm>
            <NumberInput source="coin" />
            <TextInput source="host" />
        </SimpleForm>
    </Create>
);
