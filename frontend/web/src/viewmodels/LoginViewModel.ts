import { useState, useEffect } from 'react';
import { BACKEND_URL } from 'config';
import { useRouter } from 'next/router';

export default function useLoginViewModel() {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [errors, setErrors] = useState({email: '', password: ''})
    const router = useRouter();

    const validateEmail = (email: string) => {
        const regex = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$/;
        return regex.test(email);
    }

    const validatePassword = (password: string) => {
        return password.length >= 8;
    }

    useEffect(() => {
        const timer = setTimeout(() => {
            if (email && !validateEmail(email)) {
                setErrors(prev => ({ ...prev, email: 'Invalid email format' }));
            } else {
                setErrors(prev => ({ ...prev, email: '' }));
            }
        }, 500);

        return () => clearTimeout(timer);
    }, [email]);

    useEffect(() => {
        const timer = setTimeout(() => {
            if (password && !validatePassword(password)) {
                setErrors(prev => ({ ...prev, password: 'Password must be at least 8 characters long' }));
            } else {
                setErrors(prev => ({ ...prev, password: '' }));
            }
        }, 500);

        return () => clearTimeout(timer);
    }, [password]);

    const loginUser = async () => {
        const response = await fetch(`${BACKEND_URL}/users/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                email: `${email}`,
                password: `${password}`
            }),
        });

        if (!response.ok) {
            return null
        }
        return response.json();
    }

    const handleLogin = async () => {
        if (validateEmail(email) && validatePassword(password)) {
            const userData = await loginUser();

            if (!userData) {
                setErrors(prev => ({ ...prev, email: ' ' }));
                setErrors(prev => ({ ...prev, password: 'Wrong email or password' }));
                console.error('Login error');
            } else {
                console.log(`User logged in successfully:`, userData);
                router.push('/home');
            }
        } else {
            console.error('Validation error:', errors);
        }
    };

    return {
        email,
        setEmail,
        password,
        setPassword,
        handleLogin,
        errors
    };
}
