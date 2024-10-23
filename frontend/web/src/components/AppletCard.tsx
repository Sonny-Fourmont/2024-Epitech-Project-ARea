import React from 'react';
import Image from 'next/image';
import { ProviderColors } from '@/utils/providers';

interface AppletCardProps {
    title: string;
    isConnected: boolean;
    provider: string;
    icon: string;
}

const AppletCard: React.FC<AppletCardProps> = ({ title, isConnected, provider, icon }) => {
    return (
        <div className={`p-4 rounded-lg ${isConnected ? ProviderColors[provider] : ProviderColors['default']}`} style={{ width: '100%' }}>
            <div className="mb-4">
                <Image
                    src={icon}
                    alt={`${title} icon`}
                    width={24}
                    height={24}
                    className="w-6 h-6"
                />
            </div>
            <h3 className="text-lg font-bold text-white mb-2">{title}</h3>
            <div className="mt-4">
                <p className="text-xs text-white">{isConnected ? 'Connected' : 'Not Connected'}</p>
            </div>
        </div>
    );
};

export default AppletCard;
