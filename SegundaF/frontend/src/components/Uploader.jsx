import { LuUploadCloud, LuTrash2, LuFileText } from "react-icons/lu";
import { useState } from "react";
import PropTypes from 'prop-types';

function Uploader({ height, onUpload }) {
    const [fileName, setFileName] = useState('No se ha seleccionado un archivo')
    const h = {
        "60": "h-[240px]",
        "30": "h-[125px]"
    }

    const handleSubmit = (e) => {
        e.preventDefault();
        const file = e.target[0].files[0];
        if (file) {
            onUpload(file);
        }
    }

    return (
        <div className="w-full">
            <form onSubmit={handleSubmit} action="" className="w-ful" >
                <div className={`flex flex-col  border-2 border-dashed border-[#485773] ${h[height]} cursor-pointer rounded-lg hover:border-[#418cff] transition-transform hover:transition-all ease-in-out duration-150`}
                    onClick={() => document.querySelector('.input-f').click()}>
                    <input type="file" accept=".csv" className="input-f" hidden
                        onChange={({ target: { files } }) => {
                            files[0] && setFileName(files[0].name)
                        }} />
                    <div className="w-full h-full flex flex-col items-center justify-center gap-2">
                        <LuUploadCloud size={70} color="#9ea2ab" />
                        <p className="text-text-gray-1">Clic para eligir el archivo</p>
                    </div>
                </div>

                <div className="w-full mt-3">
                    <section className="w-full flex justify-between items-center px-4 py-3 rounded-md bg-sub-dark">
                        <LuFileText size={25} color="#9ea2ab" />
                        <span className="flex items-center gap-3">
                            {fileName}
                            <LuTrash2 size={20} color="#f54e5d"
                                className="cursor-pointer"
                                onMouseDown={() => {
                                    setFileName('No se ha seleccionado un archivo')
                                    const inputFile = document.querySelector('.input-f');
                                    if (inputFile) {
                                        inputFile.value = '';
                                    }
                                }} />
                        </span>
                    </section>
                </div>
                <div className="w-full flex justify-end mt-4">
                    <button type="submit">Subir</button>
                </div>
            </form>
        </div>
    )
}

export default Uploader

Uploader.propTypes = {
    height: PropTypes.node,
    onUpload: PropTypes.func
};