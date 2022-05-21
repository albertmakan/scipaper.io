import httpClient from "./httpClient";

export const createPaper = (paper) => {
  return Promise.resolve(httpClient.post(`/papers/create-update`, paper));
};

export const updatePaper = (paper) => {
  return Promise.resolve(httpClient.put(`/papers/create-update`, paper));
};

export const read = (id) => {
  return Promise.resolve(httpClient.get(`/papers/paper/${id}`));
};

export const myPapers = () => {
  return Promise.resolve(httpClient.get(`/papers/my-papers`));
};

export const publish = (id) => {
  return Promise.resolve(httpClient.post(`/papers/publish`, { paperId: id }));
};

export const objectIdToDate = (id) => {
  return new Date(parseInt(id.substring(0, 8), 16) * 1000);
};
