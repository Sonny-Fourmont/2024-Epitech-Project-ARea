'use client'

import { useEffect, useState } from 'react';
import Navbar from '@/components/Navbar';
import Footer from '@/components/Footer';
import { BACKEND_URL } from 'config';

const MyProfilePage = () => {
    const [userData, setUserData] = useState(null);

    useEffect(() => {
        const fetchUserData = async () => {
            try {
                const response = await fetch(`${BACKEND_URL}/user`, {
                    headers: {
                        'Accept': 'application/json',
                    },
                });
                if (!response.ok) {
                    throw new Error('Network response was not ok');
                }

                const data = await response.json();
                setUserData(data);
            } catch (error) {
                console.error('Error fetching user data:', error);
            }
        };

        fetchUserData();
    }, []);

    return (
        <div>
            <Navbar />
            <ProfileSection userData={userData} />
            <Footer />
        </div>
    );
};

const ProfileSection = ({ userData }) => (
    <div className="px-8">
        <section className="flex flex-col items-center justify-center py-20 bg-hero-pattern bg-cover bg-center text-black">
            <h1 className="text-5xl font-bold mb-4">My Profile</h1>
            <p className="text-xl mb-8">Manage your account settings.</p>
            {userData && (
                <div>
                    <p>Name: {userData.name}</p>
                    <p>Email: {userData.email}</p>
                </div>
            )}
        </section>
    </div>
);

export default MyProfilePage;