import axios from "axios";
import { API_URL } from "../config";
export const graphRequest = async (graphT) => axios.post(`${API_URL}/admin/graficar`, graphT)
export const addBookRequest = async (data) => axios.post(`${API_URL}/tutor/agregar-arbolB`, data)
export const searchBookRequest = async (carnet) => axios.post(`${API_URL}/tutor/obtener-libros`, carnet)
export const addPubsRequest = async (data) => axios.post(`${API_URL}/tutor/agregar-contenido`, data)
export const getBooksRequest = async () => axios.get(`${API_URL}/obtener-tlibros`)
export const acceptBookRequest = async (data) => axios.post(`${API_URL}/admin/aceptar-libro`,data)
export const getCoursesRequest = async (data) => axios.post(`${API_URL}/estudiante/buscar-estudiante`,data)
export const getBooksStudentsRequest = async (data) => axios.post(`${API_URL}/estudiante/buscar-libros`,data)
export const getBooksAStudentsRequest = async (data) => axios.post(`${API_URL}/estudiante/buscar-libros-aprobados`,data)