'use client';

import React from 'react';
import Navbar from '@/components/Navbar';
import Footer from '@/components/Footer';

const HomePage: React.FC = () => {
    return (
        <div>
            <Navbar />
            <HeroSection />
            <FeaturesSection />
            <Footer />
        </div>
    );
};

const HeroSection: React.FC = () => (
    <section className="flex flex-col items-center justify-center py-20 bg-hero-pattern bg-cover bg-center text-black">
        <h1 className="text-5xl font-bold mb-4">Welcome to Our Service</h1>
        <p className="text-xl mb-8">Your journey to better productivity starts here.</p>
        <button className="px-6 py-3 bg-buttonColor hover:bg-buttonHoverColor text-white font-bold rounded">Get Started</button>
    </section>
);

const FeaturesSection: React.FC = () => (
    <section className="py-16 bg-gray-100">
        <div className="container mx-auto px-4">
            <h2 className="text-3xl font-bold text-center mb-8">Features</h2>
            <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
                <FeatureCard title="Feature 1" description="Description of feature 1." />
                <FeatureCard title="Feature 2" description="Description of feature 2." />
                <FeatureCard title="Feature 3" description="Description of feature 3." />
            </div>
        </div>
    </section>
);

const FeatureCard: React.FC<{ title: string, description: string }> = ({ title, description }) => (
    <div className="p-6 bg-white rounded shadow">
        <h3 className="text-xl font-bold mb-2">{title}</h3>
        <p>{description}</p>
    </div>
);

const TestimonialsSection: React.FC = () => (
    <section className="py-16 bg-white">
        <div className="container mx-auto px-4">
            <h2 className="text-3xl font-bold text-center mb-8">Testimonials</h2>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
                <TestimonialCard name="User 1" feedback="This service is amazing!" />
                <TestimonialCard name="User 2" feedback="I love using this app every day." />
            </div>
        </div>
    </section>
);

const TestimonialCard: React.FC<{ name: string, feedback: string }> = ({ name, feedback }) => (
    <div className="p-6 bg-gray-100 rounded shadow">
        <p className="italic mb-4">{feedback}</p>
        <p className="font-bold">- {name}</p>
    </div>
);

export default HomePage;