import { Navbar } from '../components/Navbar'
import { ContainerMain } from '../components/ContainerMain'
import { SelectInput } from '../components/SelectInput'

function AcceptBooks() {
    const typesPub = [
        {
            id: 1,
            nombre: "Libro 1"
        }, {
            id: 2,
            nombre: "Libro 2"
        },
        {
            id: 3,
            nombre: "Libro 3"
        },
        {
            id: 4,
            nombre: "Libro 4"
        }
    ]
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
                                    <SelectInput options={typesPub} placeHolder='Selecciona un libro' onSelectOption={() => console.log("clic")} value="2" />
                                </div>
                                <button type="submit" className='flex-1 bg-btn-green hover:bg-btn-green-hover h-10 rounded-md font-semibold transition-transform hover:transition-all ease-in-out duration-150'>Aceptar</button>
                                <button type="submit" className='flex-1 bg-btn-red hover:bg-btn-red-hover h-10 rounded-md font-semibold transition-transform hover:transition-all ease-in-out duration-150'>Rechazar</button>
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