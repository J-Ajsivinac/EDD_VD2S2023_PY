import PropTypes from 'prop-types';
import { RiBook2Fill } from "react-icons/ri";
function CardCourse({ code, name }) {
    return (
        <div className="flex basis-[46%] flex-row bg-panel-dark p-6 items-center gap-4 rounded-md">
            <RiBook2Fill size={36} color='#8c82f7' />
            <div className='flex flex-col gap-2 w-full '>
                <span className='px-3 py-1 border rounded-md w-fit'>{code}</span>
                <span>{name}</span>
            </div>
        </div>
    )
}

export default CardCourse

CardCourse.propTypes = {
    code: PropTypes.node,
    name: PropTypes.node,
};