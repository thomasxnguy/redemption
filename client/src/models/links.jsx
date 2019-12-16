import React from "react";
import {
    List,
    Datagrid,
    TextField,
    NumberField,
    BooleanField,
    DateField
} from "react-admin";

export const LinkList = props => (
    <List {...props}>
        <Datagrid rowClick="edit">
            <TextField source="link" />
            <TextField source="code" />
            <TextField source="provider" />
            <BooleanField source="valid" />
            <DateField source="expiration_date" />
            <DateField source="created_date" />
            <NumberField source="asset.coin" />
        </Datagrid>
    </List>
);
