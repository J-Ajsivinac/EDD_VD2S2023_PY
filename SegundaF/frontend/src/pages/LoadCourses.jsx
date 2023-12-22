import { Navbar } from '../components/Navbar'
import { ContainerMain } from '../components/ContainerMain'

function LoadCourses() {
    return (
        <div className='flex h-screen bg-bg-dark'>
            <Navbar></Navbar>
            <ContainerMain>
                <div className='flex w-full h-full justify-center items-center'>
                    <div className='flex items-start justify-center w-2/3 py-4 px-6 rounded-lg flex-col gap-4 text-white bg-panel-dark'>
                        <h2 className='font-bold text-xl'>Cargar Cursos</h2>
                        <div className='w-full h-60 bg-sub-dark rounded-lg'>

                        </div>
                        <button type="submit">Subir</button>
                    </div>
                </div>
            </ContainerMain>
        </div>
    )
}

export default LoadCourses