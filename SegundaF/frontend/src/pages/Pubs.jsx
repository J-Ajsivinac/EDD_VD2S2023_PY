import { Navbar } from '../components/Navbar'
import { ContainerMain } from '../components/ContainerMain'
import { CardPub } from '../components/CardPub'
import { getBooksStudentsRequest } from '../api/peticiones'
import { useEffect, useState } from 'react'

function Pubs() {
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
            const res = await getBooksStudentsRequest(data)
            console.log('---.--', res.data)
            // setData(res.data)
            setLibrosU((prevLibrosU) => {
                const nuevoLibrosU = { ...prevLibrosU };
                res.data.libros.forEach((estudiante) => {
                    if (estudiante.Publicaciones != null) {
                        if (estudiante.Carnet in nuevoLibrosU) {
                            nuevoLibrosU[estudiante.Curso].push(estudiante.Publicaciones);
                        } else {
                            nuevoLibrosU[estudiante.Curso] = [estudiante.Publicaciones];
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
                    <div className='flex flex-col items-start  w-2/3 rounded-lg gap-4 flex-wrap text-white '>
                        {
                            Object.keys(librosU).length === 0 ? <h2 className='text-white font-bold text-center'>No hay Publicaciones registradas</h2> : Object.keys(librosU).map((Curso, i) => (
                                <CardPub code={Curso} content={librosU[Curso][0]} key={i} />
                            ))
                        }
                    </div>
                </div>
            </ContainerMain>
        </div>
    )
}

export default Pubs