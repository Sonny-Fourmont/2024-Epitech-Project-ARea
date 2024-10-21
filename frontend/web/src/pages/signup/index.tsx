'use strict';

import useSignUpViewModel from '../../viewmodels/SignUpViewModel';
import Button from '../../views/components/Button';
import OAuthButtons from '../../views/components/OAuthButtons';
import Input from '../../views/components/TextInput';
import Link from 'next/link';

export default function SignupPage() {
    const {
        email,
        setEmail,
        password,
        setPassword,
        handleSignUp,
        confirmPassword,
        setConfirmPassword,
        firstName,
        setFirstName,
        lastName,
        setLastName,
        errors
    } = useSignUpViewModel();

    return (
        <div className="flex flex-col items-center justify-center min-h-screen py-2">
            <h1 className="text-2xl font-bold mb-4">Création de votre compte</h1>
            <div className="flex flex-col items-center space-y-4 w-4/5 md:w-1/2 lg:w-1/3">
                    <div className="relative w-full">
                        <Input
                            label="Nom"
                            type="text"
                            value={lastName}
                            onChange={(e) => setLastName(e.target.value)}
                            placeholder="ex: Dupont"
                            error={errors.lastName}
                        />
                        <Input
                            label="Prenom"
                            type="text"
                            value={firstName}
                            onChange={(e) => setFirstName(e.target.value)}
                            placeholder="ex: Jean"
                            error={errors.firstName}
                        />
                        <Input
                            label="Email"
                            type="email"
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                            placeholder="exemple@gmail.com"
                            error={errors.email}
                        />
                        <Input
                            label="Mot de passe"
                            type="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            placeholder="••••••••••••"
                            showPasswordToggle={true}
                            error={errors.password}
                        />
                        <Input
                            label="Confimer le mot de passe"
                            type="password"
                            value={confirmPassword}
                            onChange={(e) => setConfirmPassword(e.target.value)}
                            placeholder="••••••••••••"
                            showPasswordToggle={true}
                            error={errors.confirmPassword}
                        />
                    </div>
                    <Button type="submit" className="w-full" onClick={handleSignUp} >
                        Créer le compte
                    </Button>
                <div className="flex justify-center mt-4 w-full">
                    <OAuthButtons />
                </div>
            </div>
            <div>
                <p className="text-sm text-gray-500 p-4">
                    Déjà un compte ?{' '}
                    <Link href="/login" className="text-blue-500 hover:underline">
                        Se connecter
                    </Link>
                </p>
            </div>
        </div>
    );
}
