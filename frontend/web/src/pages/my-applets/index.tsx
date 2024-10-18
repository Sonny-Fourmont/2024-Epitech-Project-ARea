"use client";

import AppletCard from '../../views/components/AppletCard';

const MyAppletsPage = () => {
    const applets = [
    ];

    return (
        <div className="p-8">
            <h1 className="text-3xl font-bold mb-8">My Applets</h1>
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
};

export default MyAppletsPage;
