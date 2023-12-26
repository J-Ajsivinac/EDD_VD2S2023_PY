import { Navbar } from '../components/Navbar'
import { ContainerMain } from '../components/ContainerMain'
import { useAuth } from "../context/authContext";
import { useEffect } from 'react';
function LandingAdmin() {

    const { mode } = useAuth();

    useEffect(() => {
        console.log(mode)
    })

    return (
        <div className='flex h-screen bg-bg-dark'>
            <Navbar></Navbar>
            <ContainerMain>
                <div className='flex items-center justify-center w-2/3 py-4 px-6 rounded-lg flex-col gap-4 text-white h-2/3'>
                    <div className='flex flex-col items-center justify-center gap-8'>
                        <span className='text-5xl font-light mt-2 text-center'>
                            Bienvenido <span className='font-normal text-[#7696f6]'>Administrador</span>
                        </span>
                        <span className='text-center text-[#a9a9a9] text-lg'>Simplifica la gestión educativa con nuestro administrador especializado. Desde la inscripción de alumnos hasta la generación de reportes, hemos cubierto cada aspecto para que puedas centrarte en ofrecer una educación de calidad</span>
                    </div>
                </div>
                <div className='flex items-center justify-center w-2/3 py-4 px-6 rounded-lg flex-row
                gap-1 text-white bg-sub-dark'>
                    <div className='flex items-center text-center justify-center w-1/3 py-4 px-6 rounded-lg flex-col gap-4 text-white bg-sub-dark'>
                        Registra Alumnos, Tutores y Cursos sin Esfuerzo, Carga y gestiona información detallada de manera eficiente.
                    </div>
                    <div className='flex items-center text-center justify-center w-1/3 py-4 px-6 rounded-lg flex-col gap-4 text-white bg-sub-dark'>
                        Acepta libros fácilmente, realiza un seguimiento del inventario y simplifica la gestión de tu biblioteca
                    </div>
                    <div className='flex items-center text-center justify-center w-1/3 py-4 px-6 rounded-lg flex-col gap-4 text-white bg-sub-dark'>
                        Obtén información valiosa con informes detallados sobre el rendimiento académico, asistencia y más
                    </div>
                </div>
            </ContainerMain>
        </div>
    )
}

export default LandingAdmin