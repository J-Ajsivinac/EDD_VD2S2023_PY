import { Navbar } from '../components/Navbar'
import { ContainerMain } from '../components/ContainerMain'
import CardCourse from '../components/CardCourse'
import { getCoursesRequest } from '../api/peticiones'
import { useEffect, useState } from 'react'
function Courses() {
    const [courses, setCourses] = useState([])
    // const [nameCourses, setNameCourses] = useState([])
    const obtenerCursos = async () => {
        const data = {
            carnet: parseInt(localStorage.getItem('carnet')),
        }
        try {
            const res = await getCoursesRequest(data)
            setCourses(res.data.cursos)
            console.log("---", res.data.cursos)
            // setNameCourses(res.data.nombre)
            localStorage.setItem('cursos', res.data.cursos)
        } catch (error) {
            console.log(error)
        }
    }

    useEffect(() => {
        obtenerCursos()
    }, [])

    return (
        <div className='flex h-screen bg-bg-dark'>
            <Navbar></Navbar>
            <ContainerMain>
                <div className='flex w-full h-full mt-4 items-center gap-5 flex-col '>
                    <div className='flex flex-row  items-stretch  w-2/3 rounded-lg gap-4 flex-wrap text-white '>
                        <h2 className='text-white text-xl font-bold'>Listado de Cursos</h2>
                    </div>
                    <div className='flex flex-row  items-stretch  w-2/3 rounded-lg gap-4 flex-wrap text-white '>
                        {
                            courses === null || courses.length === 0 ? <h2 className='text-white font-bold text-center'>No hay Cursos registrados</h2> :
                                courses.map((course, index) => {
                                    if (course !== '') {
                                        return <CardCourse code={course} key={index} />
                                    }
                                })
                        }
                    </div>
                </div>
            </ContainerMain>
        </div>
    )
}

export default Courses