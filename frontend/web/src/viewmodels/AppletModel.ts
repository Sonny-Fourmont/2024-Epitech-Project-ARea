import { useState, useEffect } from 'react';
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

            const appletsData = await response.json();
            setApplets(appletsData);
        };

        fetchApplets();
    }, []);

    return {
        applets,
    };
}
