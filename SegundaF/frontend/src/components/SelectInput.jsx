import { useState, useEffect } from "react";
import PropTypes from 'prop-types';
// import './scroll.css'
import { FaChevronDown } from "react-icons/fa6";
import { RiBookFill } from "react-icons/ri";

export function SelectInput({ options, placeHolder, onSelectOption, value }) {

    const [openselect, setOpenSelect] = useState(false);
    const [selectedOption, setSelectedOption] = useState(value);

    useEffect(() => {
        setSelectedOption(value);
    }, [value]);

    function selectvalue(option) {
        setSelectedOption(option);
        setOpenSelect(false);
        onSelectOption(option);
    }

    function openOption() {
        setOpenSelect(true);
    }

    return (
        <div className="flex w-full">
            <input
                value={selectedOption ? selectedOption.nombre : ''}
                onClick={openOption}
                onBlur={() => {
                    setOpenSelect(false);
                }}
                id="league"
                type="text"
                className="h-10 px-2 py-4 rounded-s-md w-[92%] bg-sub-dark placeholder:text-gray-400 text-white outline-none"
                placeholder={placeHolder}
                readOnly
            />
            <div
                tabIndex={0}
                onBlur={() => {
                    setOpenSelect(false);
                }}
                className="flex items-center justify-center w-[8%] bg-sub-dark rounded-e-md h-10 "
                onClick={openOption}
            >
                <FaChevronDown size={22} color="#fff" />
            </div>
            <div className={openselect ? "absolute w-full py-2 visible translate-y-12 max-h-52 overflow-visible opacity-100 bg-sub-dark/75 backdrop-blur-sm z-10 px-2 overflow-y-auto rounded-md border border-white/30" : "absolute w-full max-h-28 bg-yellow-700 rounded-sm top-full hidden -translate-y-10 overflow-hidden overflow-y-auto transition text-yellow-500"}>
                {options.map((item, index) => (
                    <li onMouseDown={() => selectvalue(item)} key={index} className="flex items-center gap-3 z-20 list-none text-[#a0a2a9] font-semibold hover:text-white px-2 py-1 rounded-md transition-transform hover:transition-all ease-in-out duration-150">
                        <RiBookFill />
                        {item.nombre}
                    </li>
                ))}
            </div >
        </div>
    );
}

SelectInput.propTypes = {
    options: PropTypes.array.isRequired,
    placeHolder: PropTypes.node.isRequired,
    onSelectOption: PropTypes.func.isRequired,
    value: PropTypes.node,
};