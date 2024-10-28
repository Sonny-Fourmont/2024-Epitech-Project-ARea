"use client";

import AppletCard from '../../components/AppletCard';
import Navbar from '@/components/Navbar';

const MyAppletsPage = () => {
    const applets = [
        {
            title: "Google Calendar",
            isConnected: true,
            provider: "google",
            icon: "images/logos/google.svg"
        },
        {
            title: "GitHub",
            isConnected: true,
            provider: "github",
            icon: "images/logos/github.svg"
        },
        {
            title: "Spotify",
            isConnected: true,
            provider: "spotify",
            icon: "images/logos/spotify-g.svg"
        }
    ];

    return (
        <div>
            <Navbar />
            <AppletsSection applets={applets} />
        </div>
    );
};

const AppletsSection = ({ applets }) => (
    <div className="px-8">
        <section className="flex flex-col items-center justify-center py-20 bg-hero-pattern bg-cover bg-center text-black">
            <h1 className="text-5xl font-bold mb-4">My Applets</h1>
            <p className="text-xl mb-8">Connect your favorite apps to our service.</p>
        </section>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {applets.map((applet, index) => (
                <AppletCard
                    key={index}
                    title={applet.title}
                    isConnected={applet.isConnected}
                    provider={applet.provider}
                    icon={applet.icon}
                />
            ))}
        </div>
    </div>
);
export default MyAppletsPage;
