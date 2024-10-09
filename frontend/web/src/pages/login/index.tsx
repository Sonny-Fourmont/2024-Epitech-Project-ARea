"use client";

import useLoginViewModel from '../../viewmodels/LoginViewModel';
import Button from '../../views/components/Button';
import OAuthButtons from '../../views/components/OAuthButtons';
import Input from '../../views/components/TextInput';
import Link from 'next/link';

export default function LoginPage() {
    const { email, setEmail, password, setPassword, handleLogin } = useLoginViewModel();

    return (
        <div className="flex flex-col items-center justify-center min-h-screen py-2">
            <h1 className="text-2xl font-bold mb-4">Connexion à votre compte</h1>
            <div className="flex flex-col items-center space-y-4 w-4/5 md:w-1/2 lg:w-1/3">
                <form onSubmit={handleLogin} className="flex flex-col space-y-4 w-full">
                    <div className="relative w-full">
                        <label className="absolute -top-3 left-2 text-xs bg-white px-1" style={{ color: '#6C6A67' }}>
                            Email
                        </label>
                        <Input
                            label="Email"
                            type="email"
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                            placeholder="exemple@gmail.com"
                        />
                        <Input
                            label="Mot de passe"
                            type="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            placeholder="••••••••••••"
                            showPasswordToggle={true}
                        />
                    </div>
                    <Button type="submit" className="w-full">
                        Connexion
                    </Button>
                </form>
                <div className="flex justify-center mt-4 w-full">
                    <OAuthButtons />
                </div>
            </div>
            <div>
                <p className="text-sm text-gray-500 p-4">
                    Pas de compte ?{' '}
                    <Link href="/signup" className="text-blue-500 hover:underline">
                        Inscrivez-vous
                    </Link>
                </p>
            </div>
        </div>
    );
}