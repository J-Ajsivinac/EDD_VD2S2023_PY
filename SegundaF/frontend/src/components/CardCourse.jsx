import PropTypes from 'prop-types';
import { RiBook2Fill } from "react-icons/ri";
function CardCourse({ code, name }) {
    return (
        <div className="flex flex-grow flex-col bg-panel-dark px-6 py-10 items-center gap-5 rounded-md">
            <RiBook2Fill size={38} color='#8c82f7' />
            <div className='flex flex-col gap-2 w-full justify-center items-center '>
                <p className='text-xl'>{code}</p>
                <p className='text-xl'>{name}</p>
            </div>
        </div>
    )
}

export default CardCourse

CardCourse.propTypes = {
    code: PropTypes.node,
    name: PropTypes.string,
};