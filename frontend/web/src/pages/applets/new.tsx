'use client';

import React, { useState } from 'react';
import { useRouter } from 'next/router';
import Navbar from '@/components/Navbar';
import Footer from '@/components/Footer';
import useAppletModel from '@/viewmodels/AppletModel';
import { Applet } from '@/viewmodels/AppletModel';

const NewAppletPage = () => {
    const { createApplet, applets } = useAppletModel();
    const router = useRouter();

    const [ifCondition, setIfCondition] = useState('');
    const [ifType, setIfType] = useState('google');
    const [isActive, setIsActive] = useState(false);
    const [thatCondition, setThatCondition] = useState('');
    const [thatType, setThatType] = useState('google');

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault();
        const newApplet: Applet = {
            isConnected: isActive,
            if: ifCondition,
            ifType: ifType,
            that: thatCondition,
            thatType: thatType,
        };
        await createApplet(newApplet, router);
    };

    return (
        <div>
            <Navbar />
            <div className="flex items-center justify-center h-screen">
                <div className="w-1/3 p-4 bg-white rounded-lg shadow-lg">
                    <h1 className="text-2xl font-bold text-center">Create a new applet</h1>
                    <form className="mt-4" onSubmit={handleSubmit}>
                        <div className="mb-4">
                            <label className="block text-sm font-bold mb-2" htmlFor="if">If</label>
                            <input
                                className="w-full p-2 border border-gray-300 rounded-lg"
                                type="text"
                                id="if"
                                name="if"
                                value={ifCondition}
                                onChange={(e) => setIfCondition(e.target.value)}
                            />
                            <label className="block text-sm font-bold mb-2" htmlFor="ifType">If Type</label>
                            <select
                                className="w-full p-2 border border-gray-300 rounded-lg"
                                id="ifType"
                                name="ifType"
                                value={ifType}
                                onChange={(e) => setIfType(e.target.value)}
                            >
                                <option value="google">Google</option>
                                <option value="microsoft">Microsoft</option>
                                <option value="youtube">YouTube</option>
                                <option value="github">GitHub</option>
                                <option value="spotify">Spotify</option>
                            </select>
                        </div>
                        <div className="mb-4 flex items-center">
                            <input
                                className="mr-2 p-2 border border-gray-300 rounded-lg"
                                type="checkbox"
                                id="isActive"
                                name="isActive"
                                checked={isActive}
                                onChange={(e) => setIsActive(e.target.checked)}
                            />
                            <label className="text-sm font-bold" htmlFor="isActive">Active</label>
                        </div>
                        <div className="mb-4">
                            <label className="block text-sm font-bold mb-2" htmlFor="that">That</label>
                            <input
                                className="w-full p-2 border border-gray-300 rounded-lg"
                                type="text"
                                id="that"
                                name="that"
                                value={thatCondition}
                                onChange={(e) => setThatCondition(e.target.value)}
                            />
                            <label className="block text-sm font-bold mb-2" htmlFor="thatType">That Type</label>
                            <select
                                className="w-full p-2 border border-gray-300 rounded-lg"
                                id="thatType"
                                name="thatType"
                                value={thatType}
                                onChange={(e) => setThatType(e.target.value)}
                            >
                                <option value="google">Google</option>
                                <option value="microsoft">Microsoft</option>
                                <option value="youtube">YouTube</option>
                                <option value="github">GitHub</option>
                                <option value="spotify">Spotify</option>
                            </select>
                        </div>
                        <button className="w-full p-2 bg-buttonColor hover:bg-buttonHoverColor text-white font-bold rounded-lg" type="submit">Create</button>
                    </form>
                </div>
            </div>
            <Footer />
        </div>
    );
};

export default NewAppletPage;