import { Navbar } from '../components/Navbar'
import { ContainerMain } from '../components/ContainerMain'
import Uploader from '../components/uploader'
import axios from 'axios';
import { API_URL } from "../config";

function LoadStudents() {

    const handleUploadStudents = async (file) => {
        try {
            const formData = new FormData();
            formData.append('file', file);
            const resp = await axios.post(`${API_URL}/upload`, formData, {
                headers: {
                    'Content-Type': 'multipart/form-data'
                }
            });
            console.log(resp)
        } catch (e) {
            console.log(e)
        }
    }

    return (
        <div className='flex h-screen bg-bg-dark'>
            <Navbar></Navbar>
            <ContainerMain>
                <div className='flex w-full h-full mt-3 items-center flex-col gap-4 '>
                    <div className='flex items-start justify-center w-2/3 py-4 px-6 rounded-lg flex-col gap-4 text-white bg-panel-dark'>
                        <h2 className='font-bold text-xl'>Cargar Estudiantes</h2>
                        <Uploader height={"30"} onUpload={handleUploadStudents} />
                    </div>
                    <div className='flex flex-col w-2/3 bg-panel-dark py-4 px-6  rounded-lg'>
                        <h2 className='text-white font-bold text-lg'>Lista de Estudiantes</h2>
                        <div className='flex w-full px-5 mt-5 text-white'>
                            <table className='w-full'>
                                <thead>
                                    <tr>
                                        <th>Columna 1</th>
                                        <th>Columna 2</th>
                                        <th>Columna 3</th>
                                        <th>Columna 4</th>
                                        <th>Columna 5</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    <tr>
                                        <td>Fila 1, Celda 1</td>
                                        <td>Fila 1, Celda 2</td>
                                        <td>Fila 1, Celda 3</td>
                                        <td>Fila 1, Celda 4</td>
                                        <td>Fila 1, Celda 5</td>
                                    </tr>
                                    <tr>
                                        <td>Fila 2, Celda 1</td>
                                        <td>Fila 2, Celda 2</td>
                                        <td>Fila 2, Celda 3</td>
                                        <td>Fila 2, Celda 4</td>
                                        <td>Fila 2, Celda 5</td>
                                    </tr>
                                    <tr>
                                        <td>Fila 3, Celda 1</td>
                                        <td>Fila 3, Celda 2</td>
                                        <td>Fila 3, Celda 3</td>
                                        <td>Fila 3, Celda 4</td>
                                        <td>Fila 3, Celda 5</td>
                                    </tr>
                                    <tr>
                                        <td>Fila 4, Celda 1</td>
                                        <td>Fila 4, Celda 2</td>
                                        <td>Fila 4, Celda 3</td>
                                        <td>Fila 4, Celda 4</td>
                                        <td>Fila 4, Celda 5</td>
                                    </tr>
                                    <tr>
                                        <td>Fila 5, Celda 1</td>
                                        <td>Fila 5, Celda 2</td>
                                        <td>Fila 5, Celda 3</td>
                                        <td>Fila 5, Celda 4</td>
                                        <td>Fila 5, Celda 5</td>
                                    </tr>
                                </tbody>
                            </table>
                        </div>
                    </div>
                </div>
            </ContainerMain>
        </div>
    )
}

export default LoadStudents