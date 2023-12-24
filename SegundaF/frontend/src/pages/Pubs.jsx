import { Navbar } from '../components/Navbar'
import { ContainerMain } from '../components/ContainerMain'
import { CardPub } from '../components/CardPub'

function Pubs() {
    return (
        <div className='flex h-screen bg-bg-dark'>
            <Navbar></Navbar>
            <ContainerMain>
                <div className='flex w-full h-full mt-4 items-center gap-5 flex-col '>
                    <div className='flex flex-col items-start  w-2/3 rounded-lg gap-4 flex-wrap text-white '>
                        <CardPub code='123456' name='Introduccción....' content='Lorem ipsum dolor sit amet c' />
                    </div>
                </div>
            </ContainerMain>
        </div>
    )
}

export default Pubs