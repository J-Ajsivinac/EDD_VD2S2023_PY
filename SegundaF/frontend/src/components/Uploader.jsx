import { LuUploadCloud, LuTrash2, LuFileText, LuFileCheck2 } from "react-icons/lu";
import { useState } from "react";
import PropTypes from 'prop-types';

function Uploader({ height, onUpload, extension, iscontent = false }) {
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
            restart();
        }
    }

    const handleSubmit2 = (e) => {
        e.preventDefault();
        const file = e.target[0].files[0];
        const reader = new FileReader();
        console.log(file)
        reader.onloadend = async (e) => {
            const text = (e.target.result);
            const fileName = file.name;
            const fileNameS = fileName.replace(/\.[^/.]+$/, "");

            const data = {
                carnet: parseInt(localStorage.getItem('carnet')),
                nombre: fileNameS,
                contenido: text
            }
            onUpload(data);
            restart();
        };
        reader.onerror = (error) => {
            console.error("Error al leer el archivo:", error);
        };

        reader.readAsDataURL(file);
    }

    const restart = () => {
        setFileName('No se ha seleccionado un archivo')
        const inputFile = document.querySelector('.input-f');
        if (inputFile) {
            inputFile.value = '';
        }
    }


    return (
        <div className="w-full">

            <form onSubmit={!iscontent ? handleSubmit : handleSubmit2} action="" className="w-ful" >
                <div className={`flex flex-col  border-2 border-dashed border-[#485773] ${h[height]} cursor-pointer rounded-lg hover:border-[#418cff] hover:bg-blue-800/10 transition-transform hover:transition-all ease-in-out duration-150`}
                    onClick={() => document.querySelector('.input-f').click()}>
                    <input type="file" accept={extension} className="input-f" hidden
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
                        {
                            fileName === 'No se ha seleccionado un archivo' ? <LuFileText size={22} color="#9ea2ab" /> :
                                <LuFileCheck2 size={22} color="#4ef59d" />
                        }
                        <span className="flex items-center gap-3 font-semibold">
                            {fileName}
                            {
                                fileName === 'No se ha seleccionado un archivo' ? <LuTrash2 size={22} color="#9ea2ab" onClick={restart} /> :
                                    <LuTrash2 size={22} color="#ff4d4d" onClick={restart} />
                            }
                        </span>
                    </section>
                </div>
                <div className="w-full flex justify-end mt-4">
                    <button className="px-10  bg-btn-primary hover:bg-btn-primary-hover text-white py-2 rounded-md font-semibold" type="submit">Subir</button>
                </div>
            </form>
        </div>
    )
}

export default Uploader

Uploader.propTypes = {
    height: PropTypes.node,
    extension: PropTypes.node,
    onUpload: PropTypes.func,
    iscontent: PropTypes.bool,
};