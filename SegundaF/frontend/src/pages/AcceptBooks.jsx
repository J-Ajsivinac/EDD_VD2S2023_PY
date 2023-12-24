import { Navbar } from '../components/Navbar'
import { ContainerMain } from '../components/ContainerMain'

function AcceptBooks() {
    return (
        <div className='flex h-screen bg-bg-dark'>
            <Navbar></Navbar>
            <ContainerMain>
                <div className='flex w-full h-3/5 justify-center items-center flex-col'>
                    <h2 className='text-white w-2/3 px-6 font-bold text-lg'>Aceptar Libros</h2>
                    <div className='flex items-center justify-center w-2/3 py-4 px-6 rounded-lg flex-col gap-7 text-white'>
                        <div className='w-full flex flex-row items-center'>
                            <input type="text" className='basis-3/4 h-10' name="" id="" />
                            <button type="submit" className='flex-1'>Aceptar</button>
                            <button type="submit" className='flex-1'>Rechazar</button>
                        </div>
                        <button type="submit">Finalizar</button>
                    </div>
                </div>
            </ContainerMain>
        </div>
    )
}

export default AcceptBooks