import { BACKEND_URL } from 'config';

export default function useGoogleLogin() {
    const handleGoogleLogin = () => {
        window.location.href = `${BACKEND_URL}/google/login`;
    };

    return (
        handleGoogleLogin
    )
}
