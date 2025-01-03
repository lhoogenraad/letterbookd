import axios, { InternalAxiosRequestConfig } from "axios"; 

const instance = axios.create({
	baseURL : 'https://shelfd-server-a8ec075348eb.herokuapp.com/api/',
	// baseURL : 'http://localhost:8000/api/',
	headers: {
		"Content-Type": "application/json",
		timeout : 10000,
	}, 
});

/**
 * Set up axios instance to use localStorage.authToken as bearer oauth
 * token if it exists.
 */
instance.interceptors.request.use( (config: InternalAxiosRequestConfig): InternalAxiosRequestConfig => {
	const authToken = localStorage.getItem("authToken");
	if (authToken) {
		config.headers.Authorization = `Bearer ${authToken}`;
	}
	return config;
});

export default instance;

