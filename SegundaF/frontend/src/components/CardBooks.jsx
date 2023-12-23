import PropTypes from 'prop-types';
import ItemCardB from './ItemCardB';
function CardBooks({ code, title }) {
    return (
        <div className="flex flex-col w-full p-4 rounded-lg gap-2 bg-panel-dark">
            <div className='flex gap-2 px-3'>
                <span>{code}</span>
                <h2>{title}</h2>
            </div>
            <ItemCardB title='Introduccción....' />
        </div>
    )
}

export default CardBooks

CardBooks.propTypes = {
    code: PropTypes.node,
    title: PropTypes.node,
};