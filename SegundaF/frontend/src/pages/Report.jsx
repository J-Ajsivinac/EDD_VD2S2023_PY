import { Navbar } from '../components/Navbar'
import { ContainerMain } from '../components/ContainerMain'
import { useParams } from 'react-router-dom'

function Report() {
    const { graph } = useParams()

    return (
        <div className='flex h-screen bg-bg-dark'>
            <Navbar></Navbar>
            <ContainerMain>
                <div className='flex w-full h-full mt-4 items-center flex-col '>
                    <div className='flex gap-5 flex-col w-2/3 bg-panel-dark px-5 py-6 rounded-md'>
                        <div className='flex justify-between w-full bg-panel-dark items-center'>
                            <h2 className='text-white font-bold text-lg'>Reporte</h2>
                            <div className='flex gap-3 text-white items-center'>
                                <span className='font-medium'>Tipo</span>
                                <span className='px-4 bg-alt-dark py-2 rounded-lg'>{graph}</span>
                            </div>

                        </div>
                        <div className='w-full h-96 bg-sub-dark rounded-md'>
                        </div>
                        <button className='text-white'>Descargar</button>
                    </div>
                </div>
            </ContainerMain>
        </div>
    )
}

export default Report