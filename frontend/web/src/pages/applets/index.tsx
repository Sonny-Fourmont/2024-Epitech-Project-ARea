"use client";

import AppletCard from '../../components/AppletCard';
import Navbar from '@/components/Navbar';
import Footer from '@/components/Footer';
import useAppletModel from '@/viewmodels/AppletModel';

const MyAppletsPage = () => {
    const { applets } = useAppletModel();

    return (
        <div>
            <Navbar />
            <AppletsSection applets={applets} />
            <Footer />
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
            {applets.map((applet: any, index: any) => (
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
