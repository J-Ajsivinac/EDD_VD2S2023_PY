import PropTypes from 'prop-types';
import { createContext, useContext, useEffect, useState } from "react";
import { loginRequest } from "../api/auth";


const AuthContext = createContext();

export const useAuth = () => {
    const context = useContext(AuthContext);
    if (!context) throw new Error("useAuth must be used within a AuthProvider");
    return context;
};

export const AuthProvider = ({ children }) => {
    const [mode, setMode] = useState(false);
    const [errors, setErrors] = useState(null)

    const signin = async (user) => {
        try {
            const res = await loginRequest(user)
            // console.log(res)
            setMode(res.data.mode)
            localStorage.setItem('mode', res.data.mode);
            localStorage.setItem('carnet', res.data.carnet);
        } catch (error) {
            setErrors(error.response.data.message)
        }
    }

    useEffect(() => {
        const storedMode = localStorage.getItem('mode');
        if (storedMode) {
            setMode(storedMode);
        }
    }, [])

    const logout = () => {
        setMode(null)
        localStorage.removeItem('mode');
        localStorage.removeItem('carnet');
    }

    return (
        <AuthContext.Provider
            value={{
                mode,
                signin,
                errors,
                logout,
            }}
        >
            {children}
        </AuthContext.Provider>
    );
};

export default AuthContext;

AuthProvider.propTypes = {
    children: PropTypes.node.isRequired,
};