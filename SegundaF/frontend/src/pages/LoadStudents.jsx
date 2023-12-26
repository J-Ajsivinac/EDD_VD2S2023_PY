import { Navbar } from '../components/Navbar'
import { ContainerMain } from '../components/ContainerMain'
import Uploader from '../components/uploader'
import axios from 'axios';
import { API_URL } from "../config";
import { Toaster, toast } from 'sonner';
import { useEffect, useState } from "react";
import '../scroll.css'

function LoadStudents() {
    const [userData, setUserData] = useState([]);

    const getUsers = async () => {
        try {
            const resp = await axios.get(`${API_URL}/estudiantes`);
            // console.log(resp)
            setUserData(resp.data.data);
        } catch (e) {
            console.log(e);
        }
    }
    useEffect(() => {
        getUsers();
    }, [])

    const handleUploadStudents = async (file) => {
        try {
            const formData = new FormData();
            formData.append('file', file);
            const resp = await axios.post(`${API_URL}/upload`, formData, {
                headers: {
                    'Content-Type': 'multipart/form-data'
                }
            });
            toast.success(`${resp.data.message}`, { duration: 2000 })
            getUsers();

        } catch (e) {
            toast.error(`${e.data.error}`, { duration: 2000 })
        }
    }


    return (
        <div className='flex h-screen bg-bg-dark'>
            <Navbar></Navbar>
            <ContainerMain>
                <div className='flex w-full h-full mt-3 items-center flex-col gap-4 '>
                    <div className='flex items-start justify-center w-2/3 py-4 px-6 rounded-lg flex-col gap-4 text-white bg-panel-dark'>
                        <h2 className='font-bold text-xl'>Cargar Estudiantes</h2>
                        <Uploader height={"30"} onUpload={handleUploadStudents} extension=".csv" />
                    </div>
                    <div className='flex flex-col w-2/3 bg-panel-dark py-4 px-6  rounded-lg'>
                        <h2 className='text-white font-bold text-lg'>Lista de Estudiantes</h2>
                        <div className='flex w-full px-5 mt-5 text-white justify-center'>
                            {
                                userData === null || userData.length === 0 ? <h2 className='text-white font-bold text-center'>No hay estudiantes registrados</h2> :
                                    (
                                        <table className='w-full'>
                                            <thead>
                                                <tr>
                                                    <th>Indice</th>
                                                    <th>Carnet</th>
                                                    <th>Nombre</th>
                                                    <th>Password</th>
                                                    <th>Cursos</th>
                                                </tr>
                                            </thead>
                                            <tbody className=''>
                                                {userData.map((user, index) => {
                                                    return (
                                                        <tr className='' key={index}>
                                                            <td className='py-6'>{user.indice}</td>
                                                            <td>{user.carnet}</td>
                                                            <td>{user.nombre}</td>
                                                            <td className='max-w-80 overflow-x-auto'><span className=''>{user.password}</span></td>
                                                            <td className='px-6'>{user.cursos}</td>
                                                        </tr>
                                                    );
                                                })}
                                            </tbody>
                                        </table>
                                    )
                            }
                        </div>
                    </div>
                </div>
            </ContainerMain>
            <Toaster position="top-center" richColors theme="dark" />
        </div>
    )
}

export default LoadStudents