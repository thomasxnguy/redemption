import React from "react";
import { Admin, Resource, ListGuesser } from "react-admin";
import dataProvider from "./utils/dataProvider";
import { HostList } from "./models/hosts";
import { LinkList } from "./models/links";
import "./App.css";

const App = () => (
    <Admin dataProvider={dataProvider}>
        <Resource name="hosts" list={HostList} />
        <Resource name="links" list={LinkList} />
    </Admin>
);

export default App;
