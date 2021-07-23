import axios from "axios";

/**
 * Axios http instance to be used gloablly to make http calls
 */
const http = axios.create({
  baseURL: process.env.REACT_APP_API_BASE_URL,
  timeout: 1000,
});

export default http;
