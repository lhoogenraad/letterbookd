import axios, { InternalAxiosRequestConfig } from "axios"; 

const instance = axios.create({
	baseURL : 'http://shelfd.alwaysdata.net/api/',
	headers: {
		"Content-Type": "application/json",
		timeout : 1000,
	}, 
});

/**
 * Set up axios instance to use localStorage.authToken as bearer oauth
 * token if it exists.
 */
instance.interceptors.request.use( (config: InternalAxiosRequestConfig): InternalAxiosRequestConfig => {
	const authToken = process.env.PERM_TOKEN;
	if (authToken) {
		config.headers.Authorization = `Bearer ${authToken}`;
	}
	return config;
});

export default instance;

