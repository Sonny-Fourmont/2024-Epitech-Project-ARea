import React from 'react';
import OAuthButton from './OAuthButton';
import useGoogleLogin from '@/viewmodels/GoogleLogin';
import useMicrosoftLogin from '@/viewmodels/MicrosoftLogin';
import useGithubLogin from '@/viewmodels/GithubLogin';

const OAuthButtons: React.FC = () => {
    const handleGoogleLogin = useGoogleLogin();
    const handleMicrosoftLogin = useMicrosoftLogin();
    const handleGithubLogin = useGithubLogin();
    return (
        <div className="flex space-x-4">
            <OAuthButton provider="google" onClick={handleGoogleLogin} />
            <OAuthButton provider="microsoft" onClick={handleMicrosoftLogin} />
            <OAuthButton provider="github" onClick={handleGithubLogin} />
        </div>
    );
};

export default OAuthButtons;