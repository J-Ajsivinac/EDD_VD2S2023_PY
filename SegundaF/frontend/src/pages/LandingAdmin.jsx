import { Navbar } from '../components/Navbar'
import { ContainerMain } from '../components/ContainerMain'
import { useAuth } from "../context/authContext";
import { useEffect } from 'react';
import { LuCopyPlus, LuCheckSquare, LuImage } from "react-icons/lu";

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
                <div className='flex items-center justify-center w-2/3 py-2 px-6 rounded-lg flex-row
                gap-1 text-white bg-sub-dark'>
                    <div className='flex items-center text-center justify-center w-1/3 py-4 px-6 rounded-lg flex-col gap-4 text-white bg-sub-dark'>
                        <div className='bg-[#fc9000]/15 p-3 rounded-md'>
                            <LuCopyPlus size={30} color='#fc9000' />
                        </div>
                        Registra Alumnos, Tutores y Cursos sin Esfuerzo
                    </div>
                    <div className='flex items-center text-center justify-center w-1/3 py-4 px-6 rounded-lg flex-col gap-4 text-white bg-sub-dark'>
                        <div className='bg-[#f5016f]/15 p-3 rounded-md'>
                            <LuCheckSquare size={30} color='#f5016f' />
                        </div>
                        Acepta libros fácilmente
                    </div>
                    <div className='flex items-center text-center justify-center w-1/3 py-4 px-6 rounded-lg flex-col gap-4 text-white bg-sub-dark'>
                        <div className='bg-[#08d989]/15 p-3 rounded-md'>
                            <LuImage size={30} color='#08d989' />
                        </div>
                        Obtén informes detallados
                    </div>
                </div>
            </ContainerMain>
        </div>
    )
}

export default LandingAdmin