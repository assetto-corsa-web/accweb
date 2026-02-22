import axios from 'axios'

const BASE_URL = import.meta.env.VITE_BACKEND_BASE_URL || window.location.origin + '/api'

const http = axios.create({
  baseURL: BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
})

export function setAuthToken(token: string | null): void {
  if (token) {
    http.defaults.headers.common['Authorization'] = `Bearer ${token}`;
  } else {
    delete http.defaults.headers.common['Authorization'];
  }
}

export default http
