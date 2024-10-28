import React from "react";

type NavbarButtonPros = {
    text: string;
    onClick: () => void;
}

const NavbarButton: React.FC<NavbarButtonPros> = ({ text, onClick }) => {
    return (
        <button
            className="p-4 text-lg bg-buttonColor text-white hover:bg-buttonHoverColor"
            onClick={onClick}
        >
            {text}
        </button>
    );
}

export default NavbarButton;
