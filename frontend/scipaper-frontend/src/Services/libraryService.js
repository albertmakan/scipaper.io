import httpClient from "./httpClient";

export const getAll = () => {
  return Promise.resolve(
    httpClient.get(`/library/search`, { params: { query: "" } })
  );
};
