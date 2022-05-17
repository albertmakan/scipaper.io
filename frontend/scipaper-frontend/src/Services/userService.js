import httpClient from "./httpClient";

export const register = (regForm) => {
  return Promise.resolve(httpClient.post(`/users/register`, regForm));
};

export const login = (loginForm) => {
  return Promise.resolve(httpClient.post(`/users/auth`, loginForm));
};
