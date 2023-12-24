import axios from "axios";
import { API_URL } from "../config";
export const loginRequest = async (user)=> axios.post(`${API_URL}/login/`, user)