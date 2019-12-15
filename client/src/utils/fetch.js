// TODO(Dan): Fork of fetchUtils to allow for CORS-compliant testing

class HttpError extends Error {
    constructor(message, status, body) {
        super(message);
        this.name = this.constructor.name;
        if (typeof Error.captureStackTrace === "function") {
            Error.captureStackTrace(this, this.constructor);
        } else {
            this.stack = new Error(message).stack;
        }
        this.stack = new Error().stack;
    }
}

const createHeadersFromOptions = options => {
    const requestHeaders = options.headers || {
        Accept: "application/json"
    };
    if (
        !requestHeaders.has("Content-Type") &&
        !(options && (!options.method || options.method === "GET")) &&
        !(options && options.body)
    ) {
        requestHeaders.set("Content-Type", "application/json");
    }
    if (options.user && options.user.authenticated && options.user.token) {
        requestHeaders.set("Authorization", options.user.token);
    }
    return requestHeaders;
};

export const fetchJson = (url, options) => {
    console.log("fetchJson function");
    const requestHeaders = createHeadersFromOptions(options);

    return fetch(url, { ...options, headers: requestHeaders })
        .then(response =>
            response.text().then(text => ({
                status: response.status,
                statusText: response.statusText,
                headers: new Map(), // TODO(Dan): Workaround to add x-total-count
                body: text
            }))
        )
        .then(({ status, statusText, headers, body }) => {
            let json;
            try {
                json = JSON.parse(body);
            } catch (e) {
                // not json, no big deal
            }
            console.log(json);
            headers.set("x-total-count", json.length.toString());
            if (status < 200 || status >= 300) {
                return Promise.reject(
                    new HttpError(
                        (json && json.message) || statusText,
                        status,
                        json
                    )
                );
            }
            return Promise.resolve({ status, headers, body, json });
        });
};
