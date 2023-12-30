import PropTypes from 'prop-types';

function TagState({ state }) {
    if (state === 'Aceptado') {
        return (
            <span className='flex justify-center min-w-28 py-1 border-2 border-[#4ef59d] bg-[#4ef59d]/10 rounded-md w-fit text-[#4ef59d]'>{state}</span>
        )
    } else if (state === 'Rechazado') {
        return (
            <span className='flex justify-center min-w-28 py-1 border-2 border-[#f54e5d] bg-[#f54e5d]/10 rounded-md w-fit text-[#f54e5d]'>{state}</span>
        )
    } else if (state === 'Pendiente') {
        return (
            <span className='flex justify-center min-w-28 py-1 border-2 border-[#f5b64e] bg-[#f5b64e]/10 rounded-md w-fit text-[#f5b64e]'>{state}</span>
        )

    }
}

export default TagState
TagState.propTypes = {
    state: PropTypes.node,
};