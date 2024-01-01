import { useState } from "react";
import PropTypes from 'prop-types';
import { LuEye } from "react-icons/lu";

export default function Modal({ name, content }) {
    const [modal, setModal] = useState(false);

    const toggleModal = () => {
        setModal(!modal);
    };

    document.body.classList.toggle('active-modal')

    return (
        <>
            <button onClick={toggleModal} className='flex items-center gap-2 bg-btn-primary px-3 py-1 rounded-md'><LuEye size={28} /><span className='font-semibold'>Ver</span></button>
            {modal && (
                <div className="fixed w-full top-0 left-0 right-0 bottom-0 h-full z-10">
                    <div onClick={toggleModal} className="fixed top-0 left-0 right-0 bottom-0 bg-bg-dark/20 backdrop-blur-md"></div>
                    <div className="w-11/12 h-5/6 absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 bg-sub-dark py-4 px-6 rounded-md">
                        <iframe src={content} className="w-full h-full"></iframe>
                    </div>
                </div >
            )
            }
        </>
    )
}

Modal.propTypes = {
    name: PropTypes.node,
    content: PropTypes.node,
};