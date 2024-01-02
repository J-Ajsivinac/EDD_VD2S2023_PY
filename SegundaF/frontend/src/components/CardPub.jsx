import PropTypes from 'prop-types';
import Tag from './Tag';
import { LuMessageSquare } from "react-icons/lu";
export function CardPub({ code, content }) {
    return (
        <div className='border-2 border-panel-dark flex w-full bg-panel-dark py-4 px-6 rounded-lg flex-col gap-4 text-white  hover:border-border-dark transition-transform hover:transition-all ease-in-out duration-150'>
            <div className="flex w-full flex-row items-center gap-4 justify-between">
                <Tag number={code} />
                <div className='flex flex-row gap-4 items-center'>
                    <div className='flex flex-row gap-4 items-center'>
                        <LuMessageSquare size={28} color='#b0b1b2' />
                    </div>
                </div>

            </div>
            <div className='flex flex-col gap-3'>
                {
                    content.map((item, index) => {
                        return <p className='w-full bg-sub-dark rounded-md py-3 px-4' key={index}>{item}</p>
                    })
                }
            </div>
        </div>
    )
}

CardPub.propTypes = {
    code: PropTypes.node.isRequired,
    content: PropTypes.array.isRequired,
};