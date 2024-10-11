import React from 'react';
import OAuthButton from './OAuthButton';
import handleGoogleLogin from '@/viewmodels/GoogleLogin';
import handleMicrosoftLogin from '@/viewmodels/MicrosoftLogin';
import handleGithubLogin from '@/viewmodels/GithubLogin';
import handleFacebookLogin from '@/viewmodels/FacebookLogin';

const OAuthButtons: React.FC = () => {
    return (
        <div className="flex space-x-4">
            <OAuthButton provider="google" onClick={handleGoogleLogin} />
            <OAuthButton provider="microsoft" onClick={handleMicrosoftLogin} />
            <OAuthButton provider="facebook" onClick={handleFacebookLogin} />
            <OAuthButton provider="github" onClick={handleGithubLogin} />
        </div>
    );
};

export default OAuthButtons;