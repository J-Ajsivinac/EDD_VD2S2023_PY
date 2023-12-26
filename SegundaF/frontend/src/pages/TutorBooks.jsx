import { Navbar } from '../components/Navbar'
import { ContainerMain } from '../components/ContainerMain'
import ItemBooks from '../components/ItemBooks'
import Uploader from '../components/uploader'

function TutorBooks() {
    return (
        <div className='flex h-screen bg-bg-dark'>
            <Navbar></Navbar>
            <ContainerMain>
                <div className='flex w-full h-full mt-4 items-center gap-5 flex-col '>
                    <div className='flex items-start justify-center w-2/3 py-4 px-6 rounded-lg flex-col gap-4 text-white bg-panel-dark'>
                        <h2 className='font-bold text-xl'>Cargar Libros</h2>
                        <Uploader height='30' extension='.pdf' />
                    </div>
                    <div className='flex items-start justify-center w-2/3 py-4 px-6 rounded-lg flex-col gap-5 text-white bg-panel-dark'>
                        <h2 className='font-bold text-xl'>Listado de Libros</h2>
                        <div className='flex flex-col gap-4 w-full rounded-b-lg p'>
                            <ItemBooks title='Libro 1' state='Aceptado'></ItemBooks>
                            <ItemBooks title='Libro 2' state='Pendiente'></ItemBooks>
                            <ItemBooks title='Libro 2' state='Rechazado'></ItemBooks>
                        </div>
                    </div>
                </div>
            </ContainerMain>
        </div>
    )
}

export default TutorBooks