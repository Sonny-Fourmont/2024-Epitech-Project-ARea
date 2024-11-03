"use client";

import AppletCard from '../../components/AppletCard';
import Navbar from '@/components/Navbar';
import Footer from '@/components/Footer';
import { Applet } from '@/viewmodels/AppletModel';
import useAppletModel from '@/viewmodels/AppletModel';

interface AppletsSectionProps {
    applets: Applet[];
}

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

const AppletsSection: React.FC<AppletsSectionProps> = ({ applets }) => (
    <div className="px-8">
        <section className="flex flex-col items-center justify-center py-20 bg-hero-pattern bg-cover bg-center text-black">
            <h1 className="text-5xl font-bold mb-4">My Applets</h1>
            <p className="text-xl mb-8">Connect your favorite apps to our service.</p>
        </section>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {Array.isArray(applets) && applets.map((applet, index) => (
                <AppletCard
                    key={index}
                    title={applet.if + ' -> ' + applet.that}
                    isConnected={applet.isConnected}
                    provider={applet.ifType}
                    if={applet.if}
                    ifType={applet.ifType}
                    that={applet.that}
                    thatType={applet.thatType}
                />
            ))}
        </div>
    </div>
);

export default MyAppletsPage;