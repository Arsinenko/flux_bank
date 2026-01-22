import React from 'react';
import {
    Chart as ChartJS,
    ArcElement,
    Tooltip,
    Legend,
    type ChartOptions
} from 'chart.js';
import { Doughnut } from 'react-chartjs-2';

ChartJS.register(ArcElement, Tooltip, Legend);

interface CategoryDonutChartProps {
    data: {
        name: string;
        value: number;
        color: string;
    }[];
}

const CategoryDonutChart: React.FC<CategoryDonutChartProps> = ({ data }) => {
    const options: ChartOptions<'doughnut'> = {
        responsive: true,
        maintainAspectRatio: false,
        cutout: '70%',
        plugins: {
            legend: {
                position: 'right' as const,
                labels: {
                    color: '#A1A1AA',
                    usePointStyle: true,
                    pointStyle: 'circle',
                    padding: 20,
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
            }
        },
    };

    const chartData = {
        labels: data.map(d => d.name),
        datasets: [
            {
                data: data.map(d => d.value),
                backgroundColor: data.map(d => d.color),
                borderColor: '#1C1C1E',
                borderWidth: 2,
                hoverOffset: 10,
            }
        ],
    };

    return <Doughnut options={options} data={chartData} />;
};

export default CategoryDonutChart;
