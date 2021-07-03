import axios from 'axios';

const axiosInstance = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL,
});

axiosInstance.defaults.headers.post['Content-Type'] = 'application/json';

export default axiosInstance;
