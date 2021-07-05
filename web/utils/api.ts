import axios from "axios";

export const api = axios.create({ baseURL: String(process.env.BASE_API_URL) });
