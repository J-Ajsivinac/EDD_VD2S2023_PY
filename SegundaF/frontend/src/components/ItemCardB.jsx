import PropTypes from 'prop-types';
import { ImBook } from "react-icons/im";
import ModalPDF from '../components/ModalPDF';
function ItemCardB({ title, content }) {
    console.log(content)
    return (
        <div className="w-full bg-sub-dark flex justify-between py-4 px-4 rounded-md">
            <div className='flex items-center gap-4'>
                <ImBook size={20} />
                <span>{title}</span>
            </div>
            <ModalPDF name={title} content={content} />
        </div>
    )
}

export default ItemCardB
ItemCardB.propTypes = {
    title: PropTypes.node,
    content: PropTypes.node,
};