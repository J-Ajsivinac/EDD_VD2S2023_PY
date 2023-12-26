import PropTypes from 'prop-types';
import { RiBook2Fill } from "react-icons/ri";
import Tag from './Tag';
function CardCourse({ code, name }) {
    return (
        <div className="flex basis-[46%] flex-row bg-panel-dark p-6 items-center gap-5 rounded-md">
            <RiBook2Fill size={38} color='#8c82f7' />
            <div className='flex flex-col gap-2 w-full '>
                <Tag number={code} />
                <span className='text-lg font-semibold'>{name}</span>
            </div>
        </div>
    )
}

export default CardCourse

CardCourse.propTypes = {
    code: PropTypes.node,
    name: PropTypes.node,
};