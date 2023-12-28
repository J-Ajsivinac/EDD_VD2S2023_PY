import { Navbar } from '../components/Navbar'
import { ContainerMain } from '../components/ContainerMain'
import Uploader from '../components/uploader'
import { Toaster, toast } from 'sonner';
import axios from 'axios';
import { API_URL } from "../config";

function LoadCourses() {

    const handleUploadCourses = async (file) => {
        try {
            const formData = new FormData();
            formData.append('file', file);
            const resp = await axios.post(`${API_URL}/cargarCursos`, formData, {
                headers: {
                    'Content-Type': 'multipart/form-data'
                }
            });
            toast.success(`${resp.data.message}`, { duration: 2000 })

        } catch (e) {
            toast.error(`${e.data.error}`, { duration: 2000 })
        }
    }
    return (
        <div className='flex h-screen bg-bg-dark'>
            <Navbar></Navbar>
            <ContainerMain>
                <div className='flex w-full h-full justify-center items-center'>
                    <div className='flex items-start justify-center w-2/3 py-4 px-6 rounded-lg flex-col gap-4 text-white bg-panel-dark'>
                        <h2 className='font-bold text-xl'>Cargar Cursos</h2>
                        <Uploader height={"60"} extension=".json" onUpload={handleUploadCourses} />
                    </div>
                </div>
            </ContainerMain>
            <Toaster position="top-center" richColors theme="dark" />
        </div>
    )
}

export default LoadCourses