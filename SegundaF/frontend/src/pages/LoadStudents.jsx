import { Navbar } from '../components/Navbar'
import { ContainerMain } from '../components/ContainerMain'

function LoadStudents() {
    return (
        <div className='flex h-screen bg-bg-dark'>
            <Navbar></Navbar>
            <ContainerMain>
                <div className='flex w-full h-full mt-4 items-center flex-col '>
                    <div className='flex flex-col w-2/3 bg-panel-dark px-2 py-4'>
                        <div className='flex justify-between w-full bg-panel-dark items-center px-5'>
                            <h2 className='text-white font-bold text-lg'>Cargar Estudiantes</h2>
                            <button className='bg-btn-primary hover:bg-btn-primary-hover text-white font-bold py-2 px-7 rounded-md transition-transform hover:transition-all ease-in-out duration-150'>Cargar</button>
                        </div>
                        <div className='flex w-full px-5 mt-5 text-white'>
                            <table className='w-full'>
                                <thead>
                                    <tr>
                                        <th>Columna 1</th>
                                        <th>Columna 2</th>
                                        <th>Columna 3</th>
                                        <th>Columna 4</th>
                                        <th>Columna 5</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    <tr>
                                        <td>Fila 1, Celda 1</td>
                                        <td>Fila 1, Celda 2</td>
                                        <td>Fila 1, Celda 3</td>
                                        <td>Fila 1, Celda 4</td>
                                        <td>Fila 1, Celda 5</td>
                                    </tr>
                                    <tr>
                                        <td>Fila 2, Celda 1</td>
                                        <td>Fila 2, Celda 2</td>
                                        <td>Fila 2, Celda 3</td>
                                        <td>Fila 2, Celda 4</td>
                                        <td>Fila 2, Celda 5</td>
                                    </tr>
                                    <tr>
                                        <td>Fila 3, Celda 1</td>
                                        <td>Fila 3, Celda 2</td>
                                        <td>Fila 3, Celda 3</td>
                                        <td>Fila 3, Celda 4</td>
                                        <td>Fila 3, Celda 5</td>
                                    </tr>
                                    <tr>
                                        <td>Fila 4, Celda 1</td>
                                        <td>Fila 4, Celda 2</td>
                                        <td>Fila 4, Celda 3</td>
                                        <td>Fila 4, Celda 4</td>
                                        <td>Fila 4, Celda 5</td>
                                    </tr>
                                    <tr>
                                        <td>Fila 5, Celda 1</td>
                                        <td>Fila 5, Celda 2</td>
                                        <td>Fila 5, Celda 3</td>
                                        <td>Fila 5, Celda 4</td>
                                        <td>Fila 5, Celda 5</td>
                                    </tr>
                                </tbody>
                            </table>
                        </div>
                    </div>
                </div>
            </ContainerMain>
        </div>
    )
}

export default LoadStudents