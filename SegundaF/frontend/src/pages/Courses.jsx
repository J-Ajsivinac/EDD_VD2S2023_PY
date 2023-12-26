import { Navbar } from '../components/Navbar'
import { ContainerMain } from '../components/ContainerMain'
import CardCourse from '../components/CardCourse'
function Courses() {
    return (
        <div className='flex h-screen bg-bg-dark'>
            <Navbar></Navbar>
            <ContainerMain>
                <div className='flex w-full h-full mt-4 items-center gap-5 flex-col '>
                    <div className='flex flex-row  items-stretch  w-2/3 rounded-lg gap-4 flex-wrap text-white '>
                        <CardCourse code='0772' name=' Introducción a la Programación y Computación 2' />
                        <CardCourse code='0771' name='Introduccción....' />
                        <CardCourse code='0999' name='Introduccción...fff.' />
                    </div>
                </div>
            </ContainerMain>
        </div>
    )
}

export default Courses