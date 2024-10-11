import { BACKEND_URL } from 'config';

export default function useGithubLogin() {
    const handleGithubLogin = () => {
        window.location.href = `${BACKEND_URL}/github/login`;
    }
    return (
        handleGithubLogin
    )
}
