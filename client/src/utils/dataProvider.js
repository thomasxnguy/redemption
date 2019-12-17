import { fetchUtils } from "react-admin";
import { stringify } from "query-string";
import { fetchJson } from "./fetch";

const apiUrl =
    "https://cors-anywhere.herokuapp.com/redeem-trust.herokuapp.com/v1"; // TODO(Dan): Replace with dynamic URL (get URL)

const httpClient = (url, options = {}) => {
    if (!options.headers) {
        options.headers = new Headers({ Accept: "application/json" });
    }
    const token = localStorage.getItem("bearerToken");
    options.headers.set(
        "Authorization",
        `Bearer ${token}` // TODO(Dan): Only use for testing purposes)
    );
    return fetchJson(url, options);
};

export default {
    getList: (resource, params) => {
        const { page, perPage } = params.pagination;
        const { field, order } = params.sort;
        const query = {
            sort: JSON.stringify([field, order]),
            range: JSON.stringify([(page - 1) * perPage, page * perPage - 1]),
            filter: JSON.stringify(params.filter)
        };
        const url = `${apiUrl}/${resource}?${stringify(query)}`;

        return httpClient(url).then(({ headers, json }) => ({
            data: json.map((resource, i) => ({ ...resource, id: i })), // TODO(Dan): API should have an id field for each entry
            total: parseInt(
                headers
                    .get("content-range")
                    .split("/")
                    .pop(),
                10
            )
        }));
    },

    getOne: (resource, params) =>
        httpClient(`${apiUrl}/${resource}/${params.id}`).then(({ json }) => ({
            data: { ...json, id: 1 } // TODO(Dan): API should have an id field for each entry
        })),

    getMany: (resource, params) => {
        const query = {
            filter: JSON.stringify({ id: params.ids })
        };
        const url = `${apiUrl}/${resource}?${stringify(query)}`;
        return httpClient(url).then(({ json }) => ({
            data: json.map((resource, i) => ({ ...resource, id: i })) // TODO(Dan): API should have an id field for each entry
        }));
    },

    getManyReference: (resource, params) => {
        const { page, perPage } = params.pagination;
        const { field, order } = params.sort;
        const query = {
            sort: JSON.stringify([field, order]),
            range: JSON.stringify([(page - 1) * perPage, page * perPage - 1]),
            filter: JSON.stringify({
                ...params.filter,
                [params.target]: params.id
            })
        };
        const url = `${apiUrl}/${resource}?${stringify(query)}`;

        return httpClient(url).then(({ headers, json }) => ({
            data: json.map((resource, i) => ({ ...resource, id: i })), // TODO(Dan): API should have an id field for each entry
            total: parseInt(
                headers
                    .get("content-range")
                    .split("/")
                    .pop(),
                10
            )
        }));
    },

    update: (resource, params) =>
        httpClient(`${apiUrl}/${resource}/${params.id}`, {
            method: "PUT",
            body: JSON.stringify(params.data)
        }).then(({ json }) => ({ data: { ...json, id: 1 } })), // TODO(Dan): API should have an id field for each entry

    updateMany: (resource, params) => {
        const query = {
            filter: JSON.stringify({ id: params.ids })
        };
        return httpClient(`${apiUrl}/${resource}?${stringify(query)}`, {
            method: "PUT",
            body: JSON.stringify(params.data)
        }).then(({ json }) => ({ data: json }));
    },

    create: (resource, params) =>
        httpClient(`${apiUrl}/${resource}`, {
            method: "POST",
            body: JSON.stringify(params.data)
        }).then(({ json }) => ({
            data: { ...params.data, id: json.id } // TODO(Dan): API should return json.id
        })),

    delete: (resource, params) =>
        httpClient(`${apiUrl}/${resource}/${params.id}`, {
            method: "DELETE"
        }).then(({ json }) => ({ data: json })),

    deleteMany: (resource, params) => {
        const query = {
            filter: JSON.stringify({ id: params.ids })
        };
        return httpClient(`${apiUrl}/${resource}?${stringify(query)}`, {
            method: "DELETE",
            body: JSON.stringify(params.data)
        }).then(({ json }) => ({ data: json }));
    }
};
