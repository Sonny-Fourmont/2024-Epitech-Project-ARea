import React, { useState } from 'react';
import Image from 'next/image';

interface InputProps {
    label: string;
    type: 'text' | 'password' | 'email';
    value: string;
    onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
    placeholder?: string;
    showPasswordToggle?: boolean;
    error?: string;
}

const Input: React.FC<InputProps> = ({
    label,
    type,
    value,
    onChange,
    placeholder = '',
    showPasswordToggle = false,
    error,
}) => {
    const [showPassword, setShowPassword] = useState(false);

    const togglePasswordVisibility = () => {
        setShowPassword(!showPassword);
    };

    return (
        <div className="relative w-full mb-4">
            <label className="absolute -top-3 left-2 text-xs p-1 #ededed" style={{ color: '#6C6A67' }}>
                {label}
            </label>
            <input
                type={type === 'password' && showPassword ? 'text' : type}
                value={value}
                onChange={onChange}
                placeholder={placeholder}
                className={`p-2 border ${error ? 'border-red-500' : 'border-gray-300'} rounded w-full placeholder-placeholder mt-2`}
            />
            {error && <p className="text-red-600 text-sm mt-1">{error}</p>}
            {type === 'password' && showPasswordToggle && (
                <div
                    className="absolute inset-y-0 right-2 flex items-center cursor-pointer"
                    onClick={togglePasswordVisibility}
                >
                    {showPassword ? (
                        <Image
                            src="/images/logos/eye-on.svg"
                            alt="Hide password"
                            width={20}
                            height={20}
                        />
                    ) : (
                        <Image
                            src="/images/logos/eye-off.svg"
                            alt="Show password"
                            width={20}
                            height={20}
                        />
                    )}
                </div>
            )}
        </div>
    );
};

export default Input;
