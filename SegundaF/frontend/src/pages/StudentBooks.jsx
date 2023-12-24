import { Navbar } from '../components/Navbar'
import { ContainerMain } from '../components/ContainerMain'
import CardBooks from '../components/CardBooks'

function StudentBooks() {
    return (
        <div className='flex h-screen bg-bg-dark'>
            <Navbar></Navbar>
            <ContainerMain>
                <div className='flex w-full h-full mt-4 items-center gap-5 flex-col '>
                    <div className='flex items-start justify-center w-2/3 rounded-lg flex-col gap-4 text-white'>
                        <CardBooks code='123456' title='Introduccción....' />
                        <CardBooks code='222' title='Introduccción....' />
                    </div>
                </div>
            </ContainerMain>
        </div>
    )
}

export default StudentBooks