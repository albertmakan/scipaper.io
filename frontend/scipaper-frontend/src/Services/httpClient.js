import axios from "axios";
import errorHandler from "./errorHandler";

const httpClient = axios.create({
  baseURL: "http://localhost:9095",
  headers: {
    Authorization: `Bearer ${localStorage.getItem("Token")}`,
  },
});

httpClient.interceptors.response.use(
  (response) => {
    return response.data;
  },
  (error) => {
    errorHandler(error.response);
    return Promise.reject(error);
  }
);

export default httpClient;
