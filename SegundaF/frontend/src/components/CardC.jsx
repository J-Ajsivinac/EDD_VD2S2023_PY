import PropTypes from 'prop-types';

export function CardC({ children }) {
    return <div className='flex w-3/5 py-4 px-6 rounded-lg flex-col gap-4 text-white'>{children}</div>
}

CardC.propTypes = {
    children: PropTypes.node.isRequired,
};