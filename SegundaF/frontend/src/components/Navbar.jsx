import { Link } from 'react-router-dom'
import { LuSchool, LuFile, LuBookOpenCheck, LuClipboard } from "react-icons/lu";

export function Navbar() {
    return (
        <nav className="fixed top-0 w-full bg-panel-dark flex justify-between py-4 px-10 text-white items-center">
            <Link to="/" className="text-xl font-bold flex flex-row items-center gap-2">
                <LuSchool size={30} />
                <span> ProjectUp</span>
            </Link>

            <ul className="flex gap-x-3 items-center text-[#d6d6d6] font-semibold">
                <li className='flex flex-row items-center gap-2 px-4 py-3 hover:bg-alt-dark rounded-lg hover:text-white cursor-pointer'><LuFile size={20} />Cargar</li>
                <li className='flex flex-row items-center gap-2 px-4 py-3 hover:bg-alt-dark rounded-lg hover:text-white cursor-pointer'><LuBookOpenCheck size={20} />Libros</li>
                <li className='flex flex-row items-center gap-2 px-4 py-3 hover:bg-alt-dark rounded-lg hover:text-white cursor-pointer'><LuClipboard size={20} />Reportes</li>
            </ul>
        </nav>
    )
}
