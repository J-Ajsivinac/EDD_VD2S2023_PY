import axios from "axios";
import { API_URL } from "../config";
export const graphRequest = async (graphT) => axios.post(`${API_URL}/admin/graficar`, graphT)
export const addBookRequest = async (data) => axios.post(`${API_URL}/tutor/agregar-arbolB`, data)
export const searchBookRequest = async (carnet) => axios.post(`${API_URL}/tutor/obtener-libros`, carnet)
export const addPubsRequest = async (data) => axios.post(`${API_URL}/tutor/agregar-contenido`, data)