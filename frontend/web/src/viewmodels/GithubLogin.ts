import { BACKEND_URL } from 'config';
import { useEffect } from 'react';
import { useRouter } from 'next/router';

export default function useGithubLogin() {
    const router = useRouter();

    const handleGithubLogin = () => {
        window.location.href = `${BACKEND_URL}/github/login`;
    }

    // const handleGithubCallback = async () => {
    //     const urlParams = new URLSearchParams(window.location.search);
    //     const code = urlParams.get('code');
    //     if (code) {
    //         try {
    //             const response = await fetch(`${BACKEND_URL}/github/?code=${code}`);
    //             const data = await response.json();
    //             if (data.token) {
    //                 localStorage.setItem('github_token', data.token);
    //                 console.log('Token stored successfully:', data.token);
    //                 router.push('/login');
    //             } else {
    //                 console.error('Failed to retrieve token:', data.error);
    //             }
    //         } catch (error) {
    //             console.error('Error during token retrieval:', error);
    //         }
    //     }
    // }

    return (
        handleGithubLogin
    )
}
