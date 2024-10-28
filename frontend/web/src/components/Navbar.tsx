import React from "react";
import { useRouter } from "next/router";
import NavbarButton from "./NavbarButton";

const Navbar: React.FC = () => {
    const router = useRouter();

    const handleNavigation = (path: string) => {
        router.push(path);
    };

    const navItems = [
        { label: "Home", path: "/home" },
        { label: "Applets", path: "/applets" },
        { label: "Services", path: "/services" },
        { label: "Contact", path: "/contact" },
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
        <div className="flex justify-center items-start min-h-screen">
            <div className="bg-buttonColor w-full flex justify-between items-center">
                <div className="flex-1 flex justify-center">
                    {renderNavButtons()}
                </div>
                <NavbarButton
                    text="Profile"
                    onClick={() => handleNavigation("/profile")}
                />
            </div>
        </div>
    );
}

export default Navbar;