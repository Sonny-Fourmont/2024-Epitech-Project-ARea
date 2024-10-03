import { useState } from 'react';

export default function useLoginViewModel() {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const handleLogin = () => {
        // Handle login logic
    };

    return {
        email,
        setEmail,
        password,
        setPassword,
        handleLogin,
    };
}