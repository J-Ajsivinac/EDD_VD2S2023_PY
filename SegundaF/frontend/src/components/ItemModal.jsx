import { Link } from "react-router-dom"
import PropTypes from 'prop-types';

function ItemModal({ text, link }) {
    return (
        <Link to={link} className="px-4 py-2 text-[#a0a2a9] rounded-md hover:text-white transition-transform hover:transition-all ease-in-out duration-150 cursor-pointer">{text}</Link>
    )
}

export default ItemModal

ItemModal.propTypes = {
    text: PropTypes.node,
    link: PropTypes.node,
};