import PropTypes from 'prop-types';

export function CardPub({ code, name, content }) {
    return (
        <div className='border-2 border-panel-dark flex w-full bg-panel-dark py-4 px-6 rounded-lg flex-col gap-4 text-white  hover:border-border-dark transition-transform hover:transition-all ease-in-out duration-150'>
            <div className="flex w-full flex-row items-center gap-4 justify-between">
                <div className='flex flex-row gap-4 items-center'>
                    <div className='flex flex-row gap-2 items-center'>
                        <span className='px-3 py-1 border rounded-md'>{code}</span>
                        <span className='font-normal'>{name}</span>
                    </div>
                </div>

            </div>
            <div>
                <p>{content}</p>
            </div>
        </div>
    )
}

CardPub.propTypes = {
    code: PropTypes.node.isRequired,
    name: PropTypes.node.isRequired,
    content: PropTypes.node.isRequired,

};