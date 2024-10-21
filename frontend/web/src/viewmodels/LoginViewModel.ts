import { useState } from 'react';
import { BACKEND_URL } from 'config';

export default function useLoginViewModel() {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const loginUser = async () => {
        const response = await fetch(`${BACKEND_URL}/users/register`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                email,
                password
            }),
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.message || 'Failed to login user');
        }
        return response.json();
    }

    const handleLogin = () => {
        try {
            const userData = loginUser();

            console.log(`User logged in successfully:`, userData);
        } catch (error) {
            console.error('Login error:', error);
        }
    };

    return {
        email,
        setEmail,
        password,
        setPassword,
        handleLogin,
    };
}
