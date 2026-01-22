import React from 'react';

type BadgeType = 'SUCCESS' | 'WARNING' | 'ERROR' | 'INFO' | 'NEUTRAL';

interface BadgeProps {
    children: React.ReactNode;
    type?: BadgeType;
}

const Badge: React.FC<BadgeProps> = ({ children, type = 'NEUTRAL' }) => {
    const styles = {
        SUCCESS: 'bg-status-success/10 text-status-success',
        WARNING: 'bg-status-warning/10 text-status-warning',
        ERROR: 'bg-status-error/10 text-status-error',
        INFO: 'bg-status-info/10 text-status-info',
        NEUTRAL: 'bg-text-muted/10 text-text-muted',
    };

    return (
        <span className={`px-2 py-0.5 rounded-full text-xs font-semibold uppercase tracking-wider ${styles[type]}`}>
            {children}
        </span>
    );
};

export default Badge;
