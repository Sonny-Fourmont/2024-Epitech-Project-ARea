"use client";

import useLoginViewModel from '../../viewmodels/LoginViewModel';
import Button from '../../views/components/Button';
import OAuthButtons from '../../views/components/OAuthButtons';

export default function LoginPage() {
  const { email, setEmail, password, setPassword, handleLogin } = useLoginViewModel();

  return (
    <div className="flex flex-col items-center justify-center min-h-screen py-2">
      <h1 className="text-2xl font-bold mb-4">Login</h1>
      <div className="flex flex-col items-center space-y-4 w-4/5 md:w-1/2 lg:w-1/3">
        <form onSubmit={handleLogin} className="flex flex-col space-y-4 w-full">
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="Email"
            className="p-2 border border-gray-300 rounded w-full"
          />
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="Password"
            className="p-2 border border-gray-300 rounded w-full"
          />
          <Button type="submit" className="w-full">
            Login
          </Button>
        </form>
        <div className="flex justify-center mt-4 w-full">
          <OAuthButtons />
        </div>
      </div>
    </div>
  );
}