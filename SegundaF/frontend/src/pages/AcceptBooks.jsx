import { Navbar } from '../components/Navbar'
import { ContainerMain } from '../components/ContainerMain'
import { SelectInput } from '../components/SelectInput'
import { acceptBookRequest, getBooksRequest } from '../api/peticiones'
import { useEffect, useState } from 'react';

function AcceptBooks() {
    const [selectedBook, setSelectedBook] = useState('')

    useEffect(() => {
        obtenerLibros()
    }, [selectedBook])

    const librosU = []
    const obtenerLibros = async () => {
        try {
            const res = await getBooksRequest()
            console.log(res.data)
            // setData(res.data)
            res.data.data.forEach(estudiante => {
                estudiante.Libros.forEach(libro => {
                    const libroExistente = librosU.find(libroU => libroU.nombre === libro.Nombre)

                    if (!libroExistente && libro.Estado === 'Pendiente') {
                        librosU.push({
                            carnet: estudiante.Carnet,
                            nombre: libro.Nombre,
                        });
                    }
                })
            });
            console.log(librosU)
        } catch (error) {
            console.log(error)
        }
    }

    const handleSelectBook = (value) => {
        if (value) {
            setSelectedBook(value)
        }
    }

    const acceptBook = async () => {
        const data = {
            carnet: selectedBook.carnet,
            nombre: selectedBook.nombre,
            estado: 'Aceptado'
        }
        try {
            if (selectedBook === '') {
                return
            }
            const res = await acceptBookRequest(data)
            console.log(res)
            setSelectedBook(null)
            obtenerLibros()
        } catch (error) {
            console.log(error)
        }
    }

    const rejectBooks = async () => {
        const data = {
            carnet: selectedBook.carnet,
            nombre: selectedBook.nombre,
            estado: 'Rechazado'
        }
        try {
            if (selectedBook === '') {
                return
            }
            const res = await acceptBookRequest(data)
            console.log(res)
            setSelectedBook(null)
            obtenerLibros()
        } catch (error) {
            console.log(error)
        }
    }

    return (
        <div className='flex h-screen bg-bg-dark'>
            <Navbar />
            <ContainerMain>
                <div className='flex w-full h-3/5 justify-center items-center flex-col '>
                    <div className='w-2/3 bg-panel-dark py-4 px-6 rounded-lg'>
                        <h2 className='text-white full font-bold text-xl mb-2'>Aceptar Libros</h2>
                        <div className='flex items-center justify-center w-full py-4 rounded-lg flex-col gap-7 text-white'>
                            <div className='w-full flex flex-row items-center gap-4'>
                                {/* <input type="text" className='basis-3/4 h-10' name="" id="" /> */}
                                <div className='relative basis-3/4 h-10'>
                                    <SelectInput options={librosU} placeHolder='Selecciona un libro' onSelectOption={handleSelectBook} value={selectedBook} />
                                </div>
                                <button onClick={acceptBook} className='flex-1 bg-btn-green hover:bg-btn-green-hover h-10 rounded-md font-semibold transition-transform hover:transition-all ease-in-out duration-150'>Aceptar</button>
                                <button onClick={rejectBooks} className='flex-1 bg-btn-red hover:bg-btn-red-hover h-10 rounded-md font-semibold transition-transform hover:transition-all ease-in-out duration-150'>Rechazar</button>
                            </div>
                            <button type="" className='bg-btn-primary hover:bg-btn-primary-hover px-8 py-2 rounded font-semibold transition-transform hover:transition-all ease-in-out duration-150'>Finalizar</button>
                        </div>
                    </div>
                </div>
            </ContainerMain>
        </div>
    )
}

export default AcceptBooks