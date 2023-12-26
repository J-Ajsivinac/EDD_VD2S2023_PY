import { Navbar } from '../components/Navbar'
import { ContainerMain } from '../components/ContainerMain'
import Uploader from '../components/uploader'

function LoadTutor() {
    return (
        <div className='flex h-screen bg-bg-dark'>
            <Navbar></Navbar>
            <ContainerMain>
                <div className='flex w-full h-full justify-center items-center'>
                    <div className='flex items-start justify-center w-2/3 py-4 px-6 rounded-lg flex-col gap-4 text-white bg-panel-dark'>
                        <h2 className='font-bold text-xl'>Cargar Tutores</h2>
                        <Uploader height={"60"} extension=".csv"></Uploader>
                    </div>
                </div>
            </ContainerMain>
        </div>
    )
}

export default LoadTutor