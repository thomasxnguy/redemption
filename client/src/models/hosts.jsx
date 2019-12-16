import React from "react";
import { List, Datagrid, TextField, NumberField } from "react-admin";

export const HostList = props => (
    <List {...props}>
        <Datagrid rowClick="edit">
            <NumberField source="coin" />
            <TextField source="host" />
        </Datagrid>
    </List>
);
