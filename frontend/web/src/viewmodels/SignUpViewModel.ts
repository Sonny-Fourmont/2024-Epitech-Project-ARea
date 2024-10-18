import { useState, useEffect } from 'react';
import { BACKEND_URL } from 'config';
import LoginPage from '@/pages/login';

export default function useSignUpViewModel() {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [confirmPassword, setConfirmPassword] = useState('');
    const [firstName, setFirstName] = useState('');
    const [lastName, setLastName] = useState('');
    const [errors, setErrors] = useState({ email: '', password: '', confirmPassword: '', firstName: '', lastName: '' });

    const validateEmail = (email: string) => {
        const regex = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$/;
        return regex.test(email);
    }

    const validatePassword = (password: string) => {
        return password.length >= 8;
    }

    const validateConfirmPassword = (password: string, confirmPassword: string) => {
        return password === confirmPassword;
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

    useEffect(() => {
        const timer = setTimeout(() => {
            if (confirmPassword && !validateConfirmPassword(password, confirmPassword)) {
                setErrors(prev => ({ ...prev, confirmPassword: 'Passwords do not match' }));
            } else {
                setErrors(prev => ({ ...prev, confirmPassword: '' }));
            }
        }, 500);

        return () => clearTimeout(timer);
    }, [confirmPassword, password]);

    const registerUser = async () => {
        const response = await fetch(`${BACKEND_URL}/users/register`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                username: `${firstName} ${lastName}`,
                password: `${password}`,
                email: `${email}`,
            }),
            mode: 'no-cors'
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.message || 'Failed to register user');
        }

        const responseText = await response.text();
        if (responseText) {
            return JSON.parse(responseText);
        } else {
            throw new Error('Empty response from server');
        }
    };

    const handleSignUp = async () => {
        if (validateEmail(email) && validatePassword(password) && validateConfirmPassword(password, confirmPassword)) {
            try {
                const userData = await registerUser();

                console.log(`User register successfully:`, userData);
                LoginPage();
            } catch (error) {
                console.error(' Registration error:', error);
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
        confirmPassword,
        setConfirmPassword,
        firstName,
        setFirstName,
        lastName,
        setLastName,
        handleSignUp,
        errors
    };
}
