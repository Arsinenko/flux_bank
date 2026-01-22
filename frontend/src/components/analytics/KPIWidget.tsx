import React from 'react';
import Card from '../ui/Card';

interface KPIWidgetProps {
    title: string;
    value: string | number;
    trend?: number;
    suffix?: string;
    prefix?: string;
}

const KPIWidget: React.FC<KPIWidgetProps> = ({ title, value, trend, suffix = '', prefix = '' }) => {
    const isPositive = trend && trend > 0;

    return (
        <Card className="flex flex-col">
            <h4 className="text-text-muted text-xs font-bold uppercase tracking-widest mb-2">{title}</h4>
            <div className="flex items-baseline gap-1">
                <span className="text-text-muted text-xl font-medium">{prefix}</span>
                <span className="text-3xl font-black text-text-primary tabular-nums tracking-tight">
                    {value}
                </span>
                <span className="text-text-muted text-xl font-medium">{suffix}</span>
            </div>

            {trend !== undefined && (
                <div className={`mt-4 flex items-center gap-1.5 text-sm font-bold ${isPositive ? 'text-accent-neon' : 'text-status-error'}`}>
                    <span className="flex items-center justify-center w-5 h-5 rounded-full bg-current/10">
                        {isPositive ? (
                            <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="3" strokeLinecap="round" strokeLinejoin="round"><polyline points="18 15 12 9 6 15" /></svg>
                        ) : (
                            <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="3" strokeLinecap="round" strokeLinejoin="round"><polyline points="6 9 12 15 18 9" /></svg>
                        )}
                    </span>
                    <span>{Math.abs(trend * 100).toFixed(1)}%</span>
                    <span className="text-text-muted font-medium ml-1">from last month</span>
                </div>
            )}
        </Card>
    );
};

export default KPIWidget;
