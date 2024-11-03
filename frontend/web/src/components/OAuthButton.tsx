import React from 'react';
import Image from 'next/image';
import { ProviderColors, ProviderLogos } from '@/utils/providers';

type OAuthButtonProps = {
    provider: string;
    onClick: () => void;
    className?: string;
};

const OAuthButton: React.FC<OAuthButtonProps> = ({ provider, onClick, className = '' }) => {
    const providerClass = ProviderColors[provider];
    const providerLogo = ProviderLogos[provider];

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
