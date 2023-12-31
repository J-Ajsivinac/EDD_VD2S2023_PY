import { Navbar } from '../components/Navbar'
import { ContainerMain } from '../components/ContainerMain'
import ItemBooks from '../components/ItemBooks'
import Uploader from '../components/uploader'
import { addBookRequest, searchBookRequest } from "../api/peticiones";
import { useEffect, useState } from 'react';

function TutorBooks() {
    const [books, setBooks] = useState([])
    const carnetV = {
        carnet: parseInt(localStorage.getItem('carnet')),
    }
    const agregarLibro = async (data) => {
        console.log(data)
        try {
            const res = await addBookRequest(data)
            console.log(res)
            obtenerLibros()
        } catch (error) {
            console.log(error)
        }
    }

    const obtenerLibros = async () => {
        try {
            const res = await searchBookRequest(carnetV)
            // console.log(res.data.libros)
            setBooks(res.data.libros)
        } catch (error) {
            console.log(error)
        }

    }

    useEffect(() => {
        obtenerLibros()
    }, [])

    return (
        <div className='flex h-screen bg-bg-dark'>
            <Navbar></Navbar>
            <ContainerMain>
                <div className='flex w-full h-full mt-4 items-center gap-5 flex-col '>
                    <div className='flex items-start justify-center w-2/3 py-4 px-6 rounded-lg flex-col gap-4 text-white bg-panel-dark'>
                        <h2 className='font-bold text-xl'>Cargar Libros</h2>
                        <Uploader height='30' extension='.pdf' onUpload={agregarLibro} iscontent={true} />
                    </div>
                    <div className='flex items-start justify-center w-2/3 py-4 px-6 rounded-lg flex-col gap-5 text-white bg-panel-dark'>
                        <h2 className='font-bold text-xl'>Listado de Libros</h2>
                        <div className='flex flex-col gap-4 w-full rounded-b-lg p'>
                            {/* <ItemBooks title='Libro 1' state='Aceptado'></ItemBooks>
                            <ItemBooks title='Libro 2' state='Pendiente'></ItemBooks>
                            <ItemBooks title='Libro 2' state='Rechazado'></ItemBooks> */}
                            {books === null || books.length === 0 ? <h2 className='text-white font-bold text-center'>No hay Libros registrados</h2> :
                                books.map((book, index) => {
                                    return <ItemBooks title={book.Nombre} state={book.Estado} key={index} />
                                })
                            }
                        </div>
                    </div>
                </div>
            </ContainerMain>
        </div>
    )
}

export default TutorBooks