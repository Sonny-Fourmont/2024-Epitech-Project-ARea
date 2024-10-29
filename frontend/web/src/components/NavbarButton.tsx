import React from "react";

type NavbarButtonPros = {
    text: string;
    onClick: () => void;
}

const NavbarButton: React.FC<NavbarButtonPros> = ({ text, onClick }) => {
    return (
        <button
            className="p-2 bg-blue-500 text-white rounded hover:bg-blue-600"
            onClick={onClick}
        >
            {text}
        </button>
    );
}

export default NavbarButton;