import React from 'react';
import Image from 'next/image';

type OAuthButtonProps = {
  provider: 'google' | 'microsoft' | 'github';
  onClick: () => void;
  className?: string;
};

const providerStyles = {
  google: 'bg-red-500 hover:bg-red-600',
  microsoft: 'bg-blue-700 hover:bg-blue-800',
  github: 'bg-gray-800 hover:bg-gray-900',
};

const providerLogos = {
  google: 'images/logos/google.svg',
  microsoft: 'images/logos/microsoft.svg',
  github: 'images/logos/github.svg',
};

const OAuthButton: React.FC<OAuthButtonProps> = ({ provider, onClick, className = '' }) => {
  const providerClass = providerStyles[provider];
  const providerLogo = providerLogos[provider];

  return (
    <button
      onClick={onClick}
      className={`p-2 text-white rounded flex items-center justify-center ${providerClass} ${className}`}
    >
      <Image src={providerLogo} alt={`${provider} logo`} width={24} height={24} />
    </button>
  );
};

export default OAuthButton;