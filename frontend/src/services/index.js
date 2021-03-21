import axios from "axios";

const http = axios.create({
    baseURL: 'http://localhost:3000/api/',
    headers: {
        'Access-Control-Allow-Origin': '*',
    }
});

class API {

    login (payload) {
        return http.post('auth/login', payload)
    }
}

export default new API();
