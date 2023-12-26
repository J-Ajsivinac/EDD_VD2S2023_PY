import PropTypes from 'prop-types';
import { ImBook } from "react-icons/im";
import { LuEye } from "react-icons/lu";
function ItemCardB({ title }) {
    return (
        <div className="w-full bg-sub-dark flex justify-between py-4 px-4 rounded-md">
            <div className='flex items-center gap-4'>
                <ImBook size={20} />
                <span>{title}</span>
            </div>
            <button className='flex items-center gap-2 bg-btn-primary px-3 py-1 rounded-md'><LuEye size={28} /><span className='font-semibold'>Ver</span></button>
        </div>
    )
}

export default ItemCardB
ItemCardB.propTypes = {
    title: PropTypes.node,
};