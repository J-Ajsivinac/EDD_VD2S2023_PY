import PropTypes from 'prop-types';

export function Card({ children }) {
    return <div className='flex bg-panel-dark max-w-md w-full p-10 rounded-md flex-col gap-2'>{children}</div>
}

Card.propTypes = {
    children: PropTypes.node.isRequired,
};