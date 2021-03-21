import axios from "axios";
import AuthService from '../services/authService';

const api = axios.create({
    baseURL: 'http://localhost:3000/api/',
});

export default new AuthService(api);
