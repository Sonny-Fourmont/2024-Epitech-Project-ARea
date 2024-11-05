import { useState, useEffect } from 'react';
import { useRouter } from 'next/router';
import { BACKEND_URL } from 'config';

export interface Applet {
    if: string,
    ifType: string,
    isConnected: boolean,
    that: string,
    thatType: string,
}

export default function useAppletModel() {
    const [applets, setApplets] = useState<Applet[]>([]);

    useEffect(() => {
        const fetchApplets = async () => {
            const response = await fetch(`${BACKEND_URL}/applet/`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `${localStorage.getItem('Authorization')}`,
                },
            });

            if (!response.ok) {
                throw new Error('Failed to fetch applets');
            }

            const appletsData = JSON.parse(await response.json())
            setApplets(appletsData.applet_array);
        };

        fetchApplets();
    }, []);

    const createApplet = async (applet: Applet, router: ReturnType<typeof useRouter>) => {
        const response = await fetch(`${BACKEND_URL}/applet/`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `${localStorage.getItem('Authorization')}`,
            },
            body: JSON.stringify(applet),
        });

        if (!response.ok) {
            throw new Error('Failed to create applet');
        }

        const newApplet = await response.json();
        setApplets([...applets, newApplet]);
        router.push('/applets');
    };

    return {
        applets,
        createApplet,
    };
}
