import PropTypes from 'prop-types';

export function ContainerMain({ children }) {
    return <div className='mt-24 w-full flex flex-col items-center gap-4 my-4'>{children}</div>
}

ContainerMain.propTypes = {
    children: PropTypes.node.isRequired,
};