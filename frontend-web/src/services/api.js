import axios from "axios";
import authService from "./AuthService";

const API_BASE_URL = "http://localhost:3000";

const api = axios.create({
  baseURL: API_BASE_URL,
});

api.interceptors.request.use(
  (config) => {
    const authData = authService.getCurrentUser();
    if (authData && authData.token) {
      config.headers["Authorization"] = `Bearer ${authData.token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export default api;
