import React from "react";
import { Admin, Resource, ListGuesser } from "react-admin";
import { fetchJson } from "./utils/fetch";
import jsonServerProvider from "ra-data-json-server";
import "./App.css";

const httpClient = (url, options = {}) => {
    if (!options.headers) {
        options.headers = new Headers({ Accept: "application/json" });
    }
    // const token = localStorage.getItem("token");
    options.headers.set(
        "Authorization",
        `Bearer ${process.env.REACT_APP_BEARER_TOKEN}` // TODO(Dan): Only use for testing purposes)
    );
    return fetchJson(url, options);
};

const dataProvider = jsonServerProvider(
    // TODO(Dan): Workaround to use cors-anywhere to prevent CORS errors (remove before prod)
    "https://cors-anywhere.herokuapp.com/redeem-trust.herokuapp.com/v1",
    httpClient
);

const App = () => (
    <Admin dataProvider={dataProvider}>
        <Resource name="hosts" list={ListGuesser} />
        <Resource name="links" list={ListGuesser} />
    </Admin>
);

export default App;
