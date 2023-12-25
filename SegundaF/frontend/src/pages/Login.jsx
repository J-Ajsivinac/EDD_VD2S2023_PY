import { Card } from '../components/Card'
import { FaCircleUser } from "react-icons/fa6";
import { FaKey } from "react-icons/fa6";
import { FaCheck } from "react-icons/fa6";
import { useForm } from 'react-hook-form'
import { Toaster, toast } from 'sonner';
import { useNavigate } from 'react-router-dom'
import { useAuth } from "../context/authContext";
import { useEffect } from "react";

function Login() {
    const { register, handleSubmit, formState: { errors } } = useForm()
    const navigate = useNavigate()
    const { signin, errors: loginErrors, mode } = useAuth();


    const onSubmit = handleSubmit(data => {
        signin(data)
    })

    useEffect(() => {
        if (loginErrors) {
            toast.error(`${loginErrors}`, { duration: 2000 })
        }
    },)


    useEffect(() => {
        if (mode === 'admin') {
            navigate("/admin/index")
        } else if (mode === 'tutor') {
            navigate("/tutor/books")
        } else if (mode === 'user') {
            navigate("/student/courses")
        }
    }, [mode]);

    return (
        <div className='flex h-screen items-center justify-center bg-bg-dark'>
            <Card>
                <h1 className='text-2xl font-bold text-white text-center'>ProjectUp</h1>
                <span className='text-text-gray text-center'>Bienvenido de nuevo</span>
                <form onSubmit={onSubmit} className='flex flex-col gap-5 mt-6'>
                    <div className="flex flex-col">
                        <div className="relative">
                            <div className="absolute flex border border-transparent left-0 top-0 h-full w-12 p-2">
                                <div className="flex items-center justify-center rounded-tl rounded-bl z-10 bg-sub-dark text-text-gray text-lg h-full w-full">
                                    <FaCircleUser size={20} />
                                </div>
                            </div>
                            <input id="carnet"
                                {...register("carnet", { required: true })}
                                name="carnet"
                                type="text"
                                placeholder="Carnet"
                                className="text-sm sm:text-base relative w-full rounded-md border-2 border-sub-dark bg-sub-dark placeholder-gray-400 focus:border-indigo-400 focus:outline-none py-3 pr-2 pl-12 text-white" />
                        </div>
                        {errors.carnet && (<p className='text-red-400  my-1'>El Carnet es requerido</p>)}                        </div>
                    <div className="flex flex-col">
                        <div className="relative">
                            <div className="absolute flex border border-transparent left-0 top-0 h-full w-12 p-2">
                                <div className="flex items-center justify-center rounded-tl rounded-bl z-10 bg-sub-dark text-text-gray text-lg h-full w-full">
                                    <FaKey size={20} />
                                </div>
                            </div>
                            <input id="contrasena"
                                name="contrasena"
                                {...register("contrasena", { required: true })}
                                type="password"
                                placeholder="Contraseña"
                                className="text-sm sm:text-base relative w-full rounded-md border-2 border-sub-dark bg-sub-dark placeholder-gray-400 focus:border-indigo-400 focus:outline-none py-3 pr-2 pl-12 text-white" />
                        </div>
                        {errors.contrasena && (<p className='w-full text-red-400'>La Contraseña es requerida</p>)}
                    </div>
                    <div className="inline-flex items-center">
                        <label className="relative flex items-center pe-3 rounded-full cursor-pointer" htmlFor="check">
                            <input type="checkbox"
                                className="before:content[''] peer relative h-6 w-6 cursor-pointer appearance-none rounded-md border-2 border-blue-gray-200 transition-all before:absolute before:top-2/4 before:left-2/4 before:block before:h-12 before:w-12 before:-translate-y-2/4 before:-translate-x-2/4 before:rounded-full before:bg-blue-gray-500 before:opacity-0 before:transition-opacity checked:border-[#443300] checked:bg-[#f3edb9] checked:hover:before:opacity-100 "
                                id="check"
                                {...register("tutor")} />
                            <span
                                className="absolute text-white transition-opacity opacity-0 pointer-events-none top-2/4 left-2/4 -translate-y-2/4 -translate-x-full peer-checked:opacity-100">
                                <FaCheck size={12} color="#181718" />
                            </span>
                        </label>
                        <label className="mt-px text-[#bcbcbc] cursor-pointer select-none font-bold peer-checked:text-[#F0F0F0]" htmlFor="check">
                            Soy Tutor
                        </label>
                    </div>

                    <button className='bg-btn-primary hover:bg-btn-primary-hover text-white font-bold py-2 px-2 rounded-md mt-2 transition-transform hover:transition-all ease-in-out duration-150'>Iniciar sesión</button>
                </form>
            </Card>
            <Toaster position="top-center" richColors theme="dark" />
        </div>
    )
}

export default Login