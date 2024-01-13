import { Navbar } from '../components/Navbar'
import { ContainerMain } from '../components/ContainerMain'
import CardBooks from '../components/CardBooks'
import { getBooksAStudentsRequest } from '../api/peticiones'
import { useEffect, useState } from 'react'

function StudentBooks() {
    // const librosU = {"data":[]}
    const [librosU, setLibrosU] = useState({})

    useEffect(() => {
        obtenerLibros()
    }, [])
    const obtenerLibros = async () => {
        var arreglo = localStorage.getItem('cursos').split(',')
        const data = {
            codigo: arreglo,
        }
        console.log(data)
        try {
            const res = await getBooksAStudentsRequest(data)
            // setData(res.data)
            setLibrosU((prevLibrosU) => {
                const nuevoLibrosU = { ...prevLibrosU };
                res.data.libros.forEach((estudiante) => {
                    if (estudiante.Carnet in nuevoLibrosU && estudiante.Libros != null) {
                        nuevoLibrosU[estudiante.Curso].push(estudiante.Libros);
                    } else {
                        if (estudiante.Libros != null) {
                            nuevoLibrosU[estudiante.Curso] = [estudiante.Libros];

                        }
                    }
                });
                console.log("----", nuevoLibrosU);
                return nuevoLibrosU;
            });
        } catch (error) {
            console.log(error)
        }
    }
    return (
        <div className='flex h-screen bg-bg-dark'>
            <Navbar></Navbar>
            <ContainerMain>
                <div className='flex w-full h-full mt-4 items-center gap-5 flex-col '>
                    <div className='flex items-start justify-center w-2/3 rounded-lg flex-col gap-4 text-white'>

                        {
                            Object.keys(librosU).length === 0 ? <h2 className='text-white font-bold text-center'>No hay Libros registrados</h2> :
                                Object.keys(librosU).map((Curso, i) => (
                                    <CardBooks code={Curso} data={librosU[Curso][0]} key={i} />
                                ))
                        }
                    </div>
                </div>
            </ContainerMain>
        </div>
    )
}

export default StudentBooks