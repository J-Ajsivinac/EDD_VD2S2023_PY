import PropTypes from 'prop-types';
import { ImBook } from "react-icons/im";

function ItemCardB({ title }) {
    return (
        <div className="w-full bg-sub-dark flex justify-between py-3 px-4 rounded-md">
            <div className='flex  items-center gap-4'>
                <ImBook size={20} />
                <span>{title}</span>
            </div>
            <button>Ver</button>
        </div>
    )
}

export default ItemCardB
ItemCardB.propTypes = {
    title: PropTypes.node,
};