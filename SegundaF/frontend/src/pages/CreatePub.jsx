import { Navbar } from '../components/Navbar'
import { ContainerMain } from '../components/ContainerMain'
import { useForm } from 'react-hook-form'
import { LuUser2 } from "react-icons/lu";

function CreatePub() {
    const { register, formState: { errors } } = useForm()
    return (
        <div className='flex h-screen bg-bg-dark'>
            <Navbar></Navbar>
            <ContainerMain>
                <div className='flex w-full h-full items-center justify-center gap-5 flex-col '>
                    <form className='w-2/3 flex flex-col bg-panel-dark py-4 text-white rounded-md px-4  gap-4'>
                        <h2 className='font-bold text-xl'>Crear Publicación</h2>
                        <div className='flex w-full flex-row items-center gap-4'>
                            {/* <UserIcon userName={user.nombre} className="w-fit"></UserIcon> */}
                            <div className='p-3 bg-sub-dark rounded-full'>
                                <LuUser2 size={25} />
                            </div>
                            <textarea className=' pt-2 px-2 resize-none w-full bg-sub-dark rounded-md h-36 text-white outline-none' placeholder='Comentario'
                                {...register("contenido", { required: true })} />
                        </div>
                        {errors.contenido && (<p className='text-red-400 px-12 my-2'>Contenido del comentario requerido</p>)}
                        <div className='w-full flex justify-end'>
                            <button type='submit' className='px-8  bg-blue-800 hover:bg-blue-900 text-white py-2 rounded-md '>Publicar</button>
                        </div>
                    </form>
                </div>
            </ContainerMain>
        </div>
    )
}

export default CreatePub