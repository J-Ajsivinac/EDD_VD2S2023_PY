import { Navbar } from '../components/Navbar'
import { ContainerMain } from '../components/ContainerMain'
import { useParams } from 'react-router-dom'
import { graphRequest } from "../api/peticiones";
import { useEffect, useState } from "react";

function Report() {
    const { graph } = useParams()
    const [imagen, setImagen] = useState('')

    var titulo;
    if (graph === 'ArbolB') {
        titulo = 'Árbol B'
    } else if (graph === 'Grafo') {
        titulo = 'Grafo'
    } else if (graph === 'Merkle') {
        titulo = 'Árbol de Merkle'
    }

    useEffect(() => {
        const gReport = async () => {
            const peticion = {
                grafica: graph
            }
            console.log(JSON.stringify(peticion))
            try {
                const res = await graphRequest(peticion)
                console.log(res)
                // setImagen(res.data.data)
                // const url = URL.createObjectURL(res.data.graph)
                setImagen("http://localhost:3000/" + res.data.graph)
            } catch (error) {
                console.log(error)
                setImagen("")
            }
        }
        gReport();
    }, [graph])

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
                                <span className='px-4 bg-alt-dark py-2 rounded-lg'>{titulo}</span>
                            </div>

                        </div>
                        <img src={imagen} alt="Reporte de Alumnos" />
                        <button className='text-white'>Descargar</button>
                    </div>
                </div>
            </ContainerMain>
        </div>
    )
}

export default Report