import AppletCard from '../../views/components/AppletCard';

const MyAppletsPage = () => {
    const applets = [
        {
            title: 'Send email when new tweet',
            author: 'Kyroz',
            isConnected: true,
            provider: 'microsoft',
            icon: '/images/logos/microsoft.svg',
        },
        {
            title: 'Reddit Saved => RSS Email',
            author: 'Kyroz',
            isConnected: true,
            provider: 'google',
            icon: '/images/logos/google.svg',
        },
        {
            title: 'Reddit Saved => RSS Email',
            author: 'Kyroz',
            isConnected: true,
            provider: 'youtube',
            icon: '/images/logos/youtube.svg',
        },
        {
            title: 'Reddit Saved => RSS Email',
            author: 'Kyroz',
            isConnected: false,
            provider: 'gmail',
            icon: '/images/logos/gmail.svg',
        },
    ];

    return (
        <div className="p-8">
            <h1 className="text-3xl font-bold mb-8">My Applets</h1>
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                {applets.map((applet, index) => (
                    <AppletCard
                        key={index}
                        title={applet.title}
                        author={applet.author}
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
