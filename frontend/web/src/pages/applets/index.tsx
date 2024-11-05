"use client";

import AppletCard from '../../components/AppletCard';
import { useRouter } from 'next/router';
import Button from '@/components/Button';
import Navbar from '@/components/Navbar';
import Footer from '@/components/Footer';
import { Applet } from '@/viewmodels/AppletModel';
import useAppletModel from '@/viewmodels/AppletModel';

interface AppletsSectionProps {
    applets: Applet[];
    createApplet: () => void;
}

const mockApplets: Applet[] = [
    {
        isConnected: true,
        if: 'Temperature drops below 0°C',
        ifType: 'google',
        that: 'Send a notification',
        thatType: 'spotify',
    },
    {
        isConnected: false,
        if: 'Stock price goes above $1000',
        ifType: 'github',
        that: 'Send an email',
        thatType: 'github',
    },
    {
        isConnected: true,
        if: 'Steps count exceeds 10,000',
        ifType: 'microsoft',
        that: 'Log the activity',
        thatType: 'gmail',
    },
];

const MyAppletsPage = () => {
    const {
        applets,
    } = useAppletModel();
    // const applets = mockApplets;


    return (
        <div>
            <Navbar />
            <AppletsSection applets={applets} createApplet={() => {}} />
            <Footer />
        </div>
    );
};

const PlusButton: React.FC<{ onClick: () => void }> = ({ onClick }) => (
    <Button className="flex items-center justify-center w-10 h-10 bg-buttonColor hover:bg-buttonHoverColor rounded-lg cursor-pointer" onClick={onClick}>
        <span className="text-xl font-bold">+</span>
    </Button>
);

const AppletsSection: React.FC<AppletsSectionProps> = ({ applets }) => {
    const router = useRouter();

    const handleAddElement = () => {
        router.push("/applets/new");
    };
    return (
        <div className="px-8">
            <section className="flex flex-col items-center justify-center py-20 bg-hero-pattern bg-cover bg-center text-black">
                <h1 className="text-5xl font-bold mb-4">My Applets</h1>
                <p className="text-xl mb-8">Connect your favorite apps to our service.</p>
            </section>
            <div className="flex justify-left mb-8">
                <PlusButton onClick={handleAddElement} />
            </div>
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
};

export default MyAppletsPage;
