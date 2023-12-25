import { Link } from 'react-router-dom'
import { LuSchool, LuBookOpenCheck, LuLogOut, LuFile, LuClipboard } from "react-icons/lu";
import PropTypes from 'prop-types';
import Modal from './Modal'
export function Navbar() {
    const opcionesAdminC = [
        { text: 'Cursos', link: '/admin/load/courses' },
        { text: 'Tutores', link: '/admin/load/courses' },
        { text: 'Estudiantes', link: '/admin/load/students' },
    ]

    const opcionesAdminR = [
        { text: 'Arbol B', link: '/admin/report/ArbolB' },
        { text: 'Grafo', link: '/admin/report/Grafo' },
        { text: 'Merkle', link: '/admin/report/Merkle' },
    ]

    return (
        <nav className="fixed top-0 w-full bg-panel-dark flex justify-between py-4 px-10 text-white items-center">
            <Link to="/student/courses" className="text-xl font-bold flex flex-row items-center gap-2">
                <LuSchool size={30} />
                <span> ProjectUp</span>
            </Link>

            <ul className="relative flex gap-x-3 items-center text-[#d6d6d6] font-semibold">
                <Modal opciones={opcionesAdminC} name="Cursos" position={66} Icon={LuFile} />
                <Link to="/admin/accept" className='flex flex-row items-center gap-2 px-4 py-3 hover:bg-alt-dark rounded-lg hover:text-white cursor-pointer'><LuBookOpenCheck size={20} />Libros</Link>
                <Modal opciones={opcionesAdminR} name="Reportes" position={15} Icon={LuClipboard} />
                <Link to="/" className='flex flex-row items-center gap-2 px-3 py-3 hover:bg-alt-dark rounded-lg hover:text-white cursor-pointer'><LuLogOut size={20} /></Link>

            </ul>
        </nav>
    )
}

Navbar.propTypes = {
    mode: PropTypes.node,
};