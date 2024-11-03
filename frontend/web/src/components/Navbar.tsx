import React, { useEffect, useState } from "react";
import { useRouter } from "next/router";
import NavbarButton from "./NavbarButton";

const Navbar: React.FC = () => {
    const router = useRouter();
    const [isLoggedIn, setIsLoggedIn] = useState(false);

    useEffect(() => {
        const authToken = localStorage.getItem("Authorization");
        if (authToken) {
            setIsLoggedIn(true);
        }
    }, []);

    const handleNavigation = (path: string) => {
        router.push(path);
    };

    const navItems = [
        { label: "Home", path: "/home" },
        { label: "Applets", path: "/applets" },
    ];

    const renderNavButtons = () => {
        return navItems.map((item, index) => (
            <NavbarButton
                key={index}
                text={item.label}
                onClick={() => handleNavigation(item.path)}
            />
        ));
    };

    return (
        <div className="flex justify-center items-start">
            <div className="bg-buttonColor w-full flex justify-between items-center">
                <div className="flex justify-center">
                    {renderNavButtons()}
                </div>
                {isLoggedIn ? (
                    <NavbarButton
                        text="Profile"
                        onClick={() => handleNavigation("/profile")}
                    />
                ) : (
                    <NavbarButton
                        text="Login"
                        onClick={() => handleNavigation("/login")}
                    />
                )}
            </div>
        </div>
    );
};

export default Navbar;
