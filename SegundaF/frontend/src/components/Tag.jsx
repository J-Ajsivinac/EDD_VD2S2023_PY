import PropTypes from 'prop-types';

function Tag({ number }) {
    return (
        <span className='px-3 py-1 border-2 border-[#F7E683] rounded-md w-fit text-[#F7E683] font-light'>{number}</span>
    )
}

export default Tag

Tag.propTypes = {
    number: PropTypes.node,
};