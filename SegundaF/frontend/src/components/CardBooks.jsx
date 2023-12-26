import PropTypes from 'prop-types';
import ItemCardB from './ItemCardB';
import Tag from './Tag';
function CardBooks({ code, title }) {
    return (
        <div className="flex flex-col w-full px-4 py-6 rounded-lg gap-5 bg-panel-dark">
            <div className='flex items-center gap-4 px-3'>
                <Tag number={code} />
                <h2 className='text-lg font-semibold'>{title}</h2>
            </div>
            <div className='flex flex-col items-center gap-4 px-3'>
                <ItemCardB title='Introduccción....' />
                <ItemCardB title='Introduccción....' />
            </div>
        </div>
    )
}

export default CardBooks

CardBooks.propTypes = {
    code: PropTypes.node,
    title: PropTypes.node,
};