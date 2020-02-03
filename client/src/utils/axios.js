import axios from "axios";

const httpClient = axios.create({
  timeout: 10000,
  headers: {
    "Content-Type": "application/json"
  }
});
httpClient.interceptors.request.use(
  config => {
    config.headers.Authorization = `Bearer ${localStorage.getItem(
      "bearerToken"
    )}`;
    return config;
  },
  error => {
    return console.error(error);
  }
);

export { httpClient };
