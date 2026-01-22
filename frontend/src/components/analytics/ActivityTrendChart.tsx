import React from 'react';
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Filler,
    Legend,
    type ChartOptions
} from 'chart.js';
import { Line } from 'react-chartjs-2';

ChartJS.register(
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Filler,
    Legend
);

interface ActivityTrendChartProps {
    data: {
        date: string;
        loans: number;
        deposits: number;
    }[];
}

const ActivityTrendChart: React.FC<ActivityTrendChartProps> = ({ data }) => {
    const options: ChartOptions<'line'> = {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
            legend: {
                position: 'top' as const,
                align: 'end',
                labels: {
                    color: '#A1A1AA',
                    usePointStyle: true,
                    pointStyle: 'circle',
                    font: {
                        size: 12,
                        weight: 'bold'
                    }
                }
            },
            tooltip: {
                backgroundColor: '#1C1C1E',
                titleColor: '#FFFFFF',
                bodyColor: '#A1A1AA',
                borderColor: '#2E2E30',
                borderWidth: 1,
                padding: 12,
                displayColors: true,
                usePointStyle: true,
            }
        },
        scales: {
            x: {
                grid: {
                    display: false
                },
                ticks: {
                    color: '#52525B',
                    font: {
                        size: 10
                    }
                }
            },
            y: {
                grid: {
                    color: 'rgba(46, 46, 48, 0.5)',
                },
                ticks: {
                    color: '#52525B',
                    font: {
                        size: 10
                    }
                }
            }
        },
        interaction: {
            mode: 'index',
            intersect: false,
        },
    };

    const chartData = {
        labels: data.map(d => d.date.split('-')[2]), // Just the day
        datasets: [
            {
                fill: true,
                label: 'Loans',
                data: data.map(d => d.loans),
                borderColor: '#00FFC2',
                backgroundColor: (context: any) => {
                    const ctx = context.chart.ctx;
                    const gradient = ctx.createLinearGradient(0, 0, 0, 400);
                    gradient.addColorStop(0, 'rgba(0, 255, 194, 0.2)');
                    gradient.addColorStop(1, 'rgba(0, 255, 194, 0)');
                    return gradient;
                },
                tension: 0.4,
                borderWidth: 2,
                pointRadius: 0,
                pointHoverRadius: 6,
                pointHoverBackgroundColor: '#00FFC2',
                pointHoverBorderColor: '#0B0C10',
                pointHoverBorderWidth: 3,
            },
            {
                fill: true,
                label: 'Deposits',
                data: data.map(d => d.deposits),
                borderColor: '#2196F3',
                backgroundColor: (context: any) => {
                    const ctx = context.chart.ctx;
                    const gradient = ctx.createLinearGradient(0, 0, 0, 400);
                    gradient.addColorStop(0, 'rgba(33, 150, 243, 0.2)');
                    gradient.addColorStop(1, 'rgba(33, 150, 243, 0)');
                    return gradient;
                },
                tension: 0.4,
                borderWidth: 2,
                pointRadius: 0,
                pointHoverRadius: 6,
                pointHoverBackgroundColor: '#2196F3',
                pointHoverBorderColor: '#0B0C10',
                pointHoverBorderWidth: 3,
            }
        ],
    };

    return <Line options={options} data={chartData} />;
};

export default ActivityTrendChart;
