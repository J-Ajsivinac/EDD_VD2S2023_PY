import PropTypes from 'prop-types';
import { ImBook } from "react-icons/im";
function ItemBooks({ title, state }) {
    return (
        <div className='w-full bg-sub-dark flex justify-between py-3 px-4 rounded-md'>
            <div className='flex items-center gap-4'>
                <ImBook size={20} />
                <span>{title}</span>
            </div>
            <span className='border px-4 py-1 rounded-md'>{state}</span>
        </div>
    )
}

export default ItemBooks

ItemBooks.propTypes = {
    title: PropTypes.node,
    state: PropTypes.node,
};