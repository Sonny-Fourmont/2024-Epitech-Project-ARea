import React from 'react';
import Image from 'next/image';
import { ProviderColors } from '@/utils/providers';
import { ProviderLogos } from '@/utils/providers';

interface AppletCardProps {
    title: string,
    provider: string,
    if: string,
    ifType: string,
    isConnected: boolean,
    that: string,
    thatType: string,
}

const AppletCard: React.FC<AppletCardProps> = ({ title, isConnected, provider, if: ifCondition, ifType, that, thatType }) => {
    return (
        <div className={`p-4 rounded-lg ${isConnected ? ProviderColors[provider] : ProviderColors['default']}`} style={{ width: '100%' }}>
            <div className="mb-4">
                <Image
                    src={ProviderLogos[provider]}
                    alt={`${title} icon`}
                    width={24}
                    height={24}
                    className="w-6 h-6"
                />
            </div>
            <h3 className="text-lg font-bold text-white mb-2">{title}</h3>
            <div className="mt-4">
                <p className="text-xs text-white">{isConnected ? 'Connected' : 'Not Connected'}</p>
                <p className="text-xs text-white">If: {ifCondition} ({ifType})</p>
                <p className="text-xs text-white">That: {that} ({thatType})</p>
            </div>
        </div>
    );
};

export default AppletCard;
