import axios from "axios";

const BASE_API_URL = 'http://localhost:9000/api';

export default axios.create({
    baseURL: BASE_API_URL,
    headers: {
        "Content-type": "application/json"
    }
});
