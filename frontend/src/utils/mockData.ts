export const MOCK_DATA = {
    general: {
        totalAssets: 1245000000.50,
        activeUsers: 45672,
        dailyTransactionVolume: 890400.25,
        assetTrend: 0.054, // +5.4%
        userTrend: 0.012, // +1.2%
        volumeTrend: -0.021, // -2.1%
    },
    transactions: {
        heatmap: Array.from({ length: 7 * 24 }, (_, i) => ({
            day: Math.floor(i / 24),
            hour: i % 24,
            value: Math.floor(Math.random() * 100)
        })),
        categories: [
            { name: 'Food & Dining', value: 35, color: '#00FFC2' },
            { name: 'Travel', value: 20, color: '#2196F3' },
            { name: 'Shopping', value: 25, color: '#FFC107' },
            { name: 'Utilities', value: 15, color: '#FF4D4D' },
            { name: 'Other', value: 5, color: '#A1A1AA' },
        ]
    },
    loansAndDeposits: {
        trend: Array.from({ length: 30 }, (_, i) => ({
            date: `2024-03-${i + 1}`,
            loans: 5000 + Math.random() * 2000,
            deposits: 4500 + Math.random() * 3000,
            delays: Math.random() * 500
        }))
    },
    alerts: [
        { id: '1', title: 'ATM #2345 – Out of Service', level: 'HIGH', time: '10m ago' },
        { id: '2', title: 'Branch #5 – High Load', level: 'MEDIUM', time: '25m ago' },
        { id: '3', title: 'Suspicious Login Attempt', level: 'HIGH', time: '1h ago' },
        { id: '4', title: 'New VIP Customer Registered', level: 'INFO', time: '2h ago' },
    ],
    customers: [
        { id: '1', name: 'Alex Thompson', avatar: '', status: 'VERIFIED', type: 'PREMIUM', risk: 'LOW' },
        { id: '2', name: 'Maria Garcia', avatar: '', status: 'PENDING', type: 'STANDARD', risk: 'MEDIUM' },
        { id: '3', name: 'James Wilson', avatar: '', status: 'VERIFIED', type: 'VIP', risk: 'LOW' },
        { id: '4', name: 'Sarah Parker', avatar: '', status: 'BLOCKED', type: 'STANDARD', risk: 'HIGH' },
    ]
};
