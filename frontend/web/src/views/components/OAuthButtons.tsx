import React from 'react';
import OAuthButton from './OAuthButton';

const OAuthButtons: React.FC = () => {
  const handleGoogleLogin = () => {
    // Handle Google login logic
  };

  const handleMicrosoftLogin = () => {
    // Handle Microsoft login logic
  };

  const handleGithubLogin = () => {
    // Handle GitHub login logic
  };

  return (
    <div className="flex space-x-4">
      <OAuthButton provider="google" onClick={handleGoogleLogin} />
      <OAuthButton provider="microsoft" onClick={handleMicrosoftLogin} />
      <OAuthButton provider="github" onClick={handleGithubLogin} />
    </div>
  );
};

export default OAuthButtons;