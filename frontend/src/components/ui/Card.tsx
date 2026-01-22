import React from 'react';

interface CardProps {
    title?: string;
    children: React.ReactNode;
    className?: string;
    showMenu?: boolean;
}

const Card: React.FC<CardProps> = ({ title, children, className = '', showMenu = true }) => {
    return (
        <div className={`bg-surface-card rounded-2xl border border-surface-border p-6 shadow-xl transition-all duration-300 hover:border-accent-neon/30 hover:shadow-neon/20 hover:-translate-y-1 ${className}`}>
            {title && (
                <div className="flex justify-between items-center mb-6">
                    <h3 className="text-xl font-bold text-text-primary">{title}</h3>
                    {showMenu && (
                        <button className="text-text-muted hover:text-text-primary transition-colors p-1">
                            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                                <circle cx="12" cy="12" r="1" /><circle cx="12" cy="5" r="1" /><circle cx="12" cy="19" r="1" />
                            </svg>
                        </button>
                    )}
                </div>
            )}
            {children}
        </div>
    );
};

export default Card;
