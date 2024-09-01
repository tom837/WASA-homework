import axios from "axios";

// Set the base URL for your backend API
axios.defaults.baseURL = 'http://localhost:3000';

const instance = axios.create({
	baseURL: __API_URL__,
	timeout: 1000 * 5
});



export default axios;