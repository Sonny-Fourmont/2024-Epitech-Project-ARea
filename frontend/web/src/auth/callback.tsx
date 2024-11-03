import { useEffect } from "react";
import { useRouter } from "next/router";

const AuthCallback = () => {
    const router = useRouter();
    const { isReady, query } = router;

    useEffect(() => {
        if (isReady && query.token) {
            const token = query.token as string;
            localStorage.setItem("Authorization", `Bearer ${token}`);
            router.replace("/home");
        }
    }, [isReady, query.token, router]);

    return <p>Logging in...</p>;
};

export default AuthCallback;
