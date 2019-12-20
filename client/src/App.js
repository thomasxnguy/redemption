import React from "react";
import { Admin, Resource } from "react-admin";
import dataProvider from "./utils/dataProvider";
import { HostList, HostCreate } from "./models/hosts";
import { LinkList, LinkEdit, LinkCreate } from "./models/links";
import Dashboard from "./components/dashboard";
import "./App.css";

const App = () => (
    <Admin dataProvider={dataProvider} dashboard={Dashboard}>
        <Resource name="hosts" list={HostList} create={HostCreate} />
        <Resource
            name="links"
            list={LinkList}
            edit={LinkEdit}
            create={LinkCreate}
        />
    </Admin>
);

export default App;
