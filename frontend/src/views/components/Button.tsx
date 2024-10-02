import React from 'react';

type ButtonProps = {
  children: React.ReactNode;
  onClick?: () => void;
  type?: 'button' | 'submit' | 'reset';
  className?: string;
};

const Button: React.FC<ButtonProps> = ({ children, onClick, type = 'button', className = '' }) => {
  return (
    <button
      type={type}
      onClick={onClick}
      className={`p-2 bg-blue-500 text-white rounded hover:bg-blue-600 ${className}`}
    >
      {children}
    </button>
  );
};

export default Button;