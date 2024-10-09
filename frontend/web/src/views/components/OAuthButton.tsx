import React from 'react';
import Image from 'next/image';

type OAuthButtonProps = {
    provider: 'google' | 'microsoft' | 'github' | 'facebook';
    onClick: () => void;
    className?: string;
};

const providerStyles = {
    google: 'bg-red-100 hover:bg-red-200 text-red-700',
    microsoft: 'bg-blue-100 hover:bg-blue-200 text-blue-700',
    github: 'bg-gray-100 hover:bg-gray-200 text-gray-700',
    facebook: 'bg-blue-100 hover:bg-blue-200 text-blue-600',
};

const providerLogos = {
    google: 'images/logos/google.svg',
    microsoft: 'images/logos/microsoft.svg',
    github: 'images/logos/github.svg',
    facebook: 'images/logos/facebook.svg',
};

const OAuthButton: React.FC<OAuthButtonProps> = ({ provider, onClick, className = '' }) => {
    const providerClass = providerStyles[provider];
    const providerLogo = providerLogos[provider];

    return (
        <button
            onClick={onClick}
            className={`p-2 text-white rounded flex items-center justify-center ${providerClass} ${className}`}
        >
            <Image
                src={providerLogo}
                alt={`${provider} logo`}
                width={24}
                height={24}
            />
        </button>
    );
};

export default OAuthButton;