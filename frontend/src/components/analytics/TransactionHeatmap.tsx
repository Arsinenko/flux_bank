import React from 'react';

interface HeatmapData {
    day: number;
    hour: number;
    value: number;
}

interface TransactionHeatmapProps {
    data: HeatmapData[];
}

const DAYS = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'];
const HOURS = ['00', '04', '08', '12', '16', '20'];

const TransactionHeatmap: React.FC<TransactionHeatmapProps> = ({ data }) => {
    // Helper to get color based on value (0-100)
    const getColor = (value: number) => {
        if (value === 0) return 'bg-surface-border/20';
        if (value < 20) return 'bg-accent-neon/10';
        if (value < 40) return 'bg-accent-neon/30';
        if (value < 60) return 'bg-accent-neon/50';
        if (value < 80) return 'bg-accent-neon/70';
        return 'bg-accent-neon shadow-neon';
    };

    return (
        <div className="flex flex-col gap-4 w-full h-full">
            <div className="flex-1 grid grid-cols-[auto_1fr] gap-2">
                {/* Day Labels */}
                <div className="grid grid-rows-7 gap-1 pt-6 text-[10px] font-bold text-text-muted uppercase">
                    {DAYS.map(day => <div key={day} className="h-4 flex items-center">{day}</div>)}
                </div>

                {/* The Grid */}
                <div className="flex flex-col gap-2 overflow-hidden">
                    <div className="grid grid-cols-24 gap-1">
                        {Array.from({ length: 7 * 24 }).map((_, i) => {
                            const entry = data.find(d => d.day === Math.floor(i / 24) && d.hour === (i % 24));
                            const val = entry ? entry.value : 0;
                            return (
                                <div
                                    key={i}
                                    className={`h-4 rounded-sm transition-all hover:scale-125 hover:z-10 cursor-pointer ${getColor(val)}`}
                                    title={`Value: ${val}`}
                                />
                            );
                        })}
                    </div>

                    {/* Hour Labels */}
                    <div className="grid grid-cols-24 gap-1 text-[8px] font-bold text-text-muted">
                        {Array.from({ length: 24 }).map((_, i) => (
                            <div key={i} className="text-center">
                                {i % 4 === 0 ? HOURS[Math.floor(i / 4)] : ''}
                            </div>
                        ))}
                    </div>
                </div>
            </div>

            {/* Legend */}
            <div className="flex items-center gap-4 mt-2">
                <span className="text-[10px] text-text-muted font-bold uppercase">Activity:</span>
                <div className="flex gap-1 items-center">
                    <span className="text-[8px] text-text-muted">Less</span>
                    <div className="w-3 h-3 rounded-sm bg-surface-border/20"></div>
                    <div className="w-3 h-3 rounded-sm bg-accent-neon/20"></div>
                    <div className="w-3 h-3 rounded-sm bg-accent-neon/50"></div>
                    <div className="w-3 h-3 rounded-sm bg-accent-neon/80"></div>
                    <div className="w-3 h-3 rounded-sm bg-accent-neon shadow-neon"></div>
                    <span className="text-[8px] text-text-muted">More</span>
                </div>
            </div>
        </div>
    );
};

export default TransactionHeatmap;
