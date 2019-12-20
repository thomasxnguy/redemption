import { fetchUtils } from "react-admin";
import { stringify } from "query-string";
import { fetchJson } from "./fetch";

const apiUrl = "/v1"; // TODO(Dan): Replace with dynamic URL (get URL)

/**
 * Helper method to format paths (since our API does not follow REST conventions)
 * @param {*} resource
 */
const apiPath = resource => {
    switch (resource) {
        case "links":
            return "link";
        case "hosts":
            return "hosts";
        default:
            console.error("This path is not a resource URL");
            return;
    }
};

/**
 * Helper method to format paths (since our API does not follow REST conventions)
 * @param {*} resource
 */
const resourceId = resource => {
    switch (resource) {
        case "links":
            return "code";
        case "hosts":
            return "coin";
        default:
            console.error("This path is not a resource URL");
            return;
    }
};

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
        const id = resourceId(resource);

        return httpClient(url).then(({ headers, json }) => ({
            data: json.map((resource, i) => ({
                ...resource,
                id: resource[id]
            })),
            total: parseInt(
                headers
                    .get("content-range")
                    .split("/")
                    .pop(),
                10
            )
        }));
    },

    getOne: (resource, params) => {
        const id = resourceId(resource);
        return httpClient(`${apiUrl}/${apiPath(resource)}/${params.id}`).then(
            ({ json }) => ({
                data: { ...json, id: json[id] }
            }),
            err => {
                console.error(err);
            }
        );
    },

    getMany: (resource, params) => {
        const query = {
            filter: JSON.stringify({ id: params.ids })
        };
        const url = `${apiUrl}/${resource}?${stringify(query)}`;
        const id = resourceId(resource);
        return httpClient(url).then(({ json }) => ({
            data: json.map((resource, i) => ({ ...resource, id: resource[id] }))
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

    update: (resource, params) => {
        const method = resource === "links" ? "POST" : "PUT";
        const id = resourceId(resource);
        return httpClient(`${apiUrl}/${apiPath(resource)}/${params.id}`, {
            method: method,
            body: JSON.stringify(params.data)
        }).then(
            ({ json }) => {
                return { data: { ...json, id: json[id] } };
            },
            err => console.error(err)
        );
    },

    updateMany: (resource, params) => {
        const query = {
            filter: JSON.stringify({ id: params.ids })
        };
        return httpClient(`${apiUrl}/${resource}?${stringify(query)}`, {
            method: "PUT",
            body: JSON.stringify(params.data)
        }).then(({ json }) => ({ data: json }));
    },

    create: (resource, params) => {
        const createPath = resource === "links" ? `/create` : ""; // Note: Because of peculiarity of API paths
        const method = resource === "hosts" ? "PUT" : "POST"; // Note: Because peculiarity of API
        const body = resource === "hosts" ? [params.data] : params.data; // NOTE: Because peculiarity of API
        return httpClient(`${apiUrl}/${resource}${createPath}`, {
            method: method,
            body: JSON.stringify(body)
        }).then(
            ({ json }) => {
                const id =
                    resource === "links"
                        ? json[0].code // Note: Hackaround to get 1st link's code
                        : Math.round(Math.random() * 9999999999); // Note: Hackaround to assign unique ID for new host
                return {
                    data: { ...params.data, id }
                };
            },
            err => console.error(err)
        );
    },

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
