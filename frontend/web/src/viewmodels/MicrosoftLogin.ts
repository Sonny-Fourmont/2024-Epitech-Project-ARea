import { BACKEND_URL } from 'config';

export default function useMicrosoftLogin() {
    const handleMicrosoftLogin = () => {
        window.location.href = `${BACKEND_URL}/microsoft/login`;
    };

    return (
        handleMicrosoftLogin
    )
}
