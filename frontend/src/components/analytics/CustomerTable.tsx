import React from 'react';
import Badge from '../ui/Badge';

interface Customer {
    id: string;
    name: string;
    email: string;
    avatar: string;
    status: string;
    type: string;
    risk: string;
    joinDate: string;
    balance: number;
}

interface CustomerTableProps {
    customers: Customer[];
}

const CustomerTable: React.FC<CustomerTableProps> = ({ customers }) => {
    return (
        <div className="w-full overflow-x-auto scrollbar-hide">
            <table className="w-full text-left border-collapse">
                <thead>
                    <tr className="border-b border-surface-border text-text-muted text-[10px] font-bold uppercase tracking-widest">
                        <th className="py-4 px-4 font-bold">Customer</th>
                        <th className="py-4 px-4 font-bold">Status</th>
                        <th className="py-4 px-4 font-bold">Verification</th>
                        <th className="py-4 px-4 font-bold">Joined</th>
                        <th className="py-4 px-4 font-bold text-right">Balance</th>
                        <th className="py-4 px-4"></th>
                    </tr>
                </thead>
                <tbody className="divide-y divide-surface-border/30">
                    {customers.map((customer) => (
                        <tr key={customer.id} className="group hover:bg-background-dark/30 transition-colors">
                            <td className="py-4 px-4">
                                <div className="flex items-center gap-3">
                                    <div className="w-10 h-10 rounded-full bg-accent-neon/10 border border-accent-neon/20 flex items-center justify-center text-accent-neon font-bold text-sm overflow-hidden shadow-sm group-hover:shadow-neon/20 transition-all">
                                        {customer.avatar ? (
                                            <img src={customer.avatar} alt={customer.name} />
                                        ) : (
                                            <span>{customer.name.split(' ').map(n => n[0]).join('')}</span>
                                        )}
                                    </div>
                                    <div className="flex flex-col">
                                        <span className="text-text-primary font-bold text-sm">{customer.name}</span>
                                        <span className="text-text-muted text-[10px]">{customer.email}</span>
                                    </div>
                                </div>
                            </td>
                            <td className="py-4 px-4">
                                <span className={`text-[10px] font-bold px-2 py-0.5 rounded-full ${customer.type === 'VIP' ? 'bg-accent-neon/20 text-accent-neon border border-accent-neon/30' :
                                        customer.type === 'PREMIUM' ? 'bg-status-info/20 text-status-info border border-status-info/30' :
                                            'bg-text-muted/20 text-text-muted border border-text-muted/30'
                                    }`}>
                                    {customer.type}
                                </span>
                            </td>
                            <td className="py-4 px-4">
                                <Badge type={
                                    customer.status === 'VERIFIED' ? 'SUCCESS' :
                                        customer.status === 'PENDING' ? 'WARNING' :
                                            'ERROR'
                                }>
                                    {customer.status}
                                </Badge>
                            </td>
                            <td className="py-4 px-4 text-text-secondary text-xs tabular-nums">
                                {new Date(customer.joinDate).toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' })}
                            </td>
                            <td className="py-4 px-4 text-right">
                                <span className="text-text-primary font-bold tabular-nums">
                                    ${customer.balance.toLocaleString('en-US', { minimumFractionDigits: 2 })}
                                </span>
                            </td>
                            <td className="py-4 px-4 text-right">
                                <button className="text-text-muted hover:text-accent-neon p-2 rounded-lg hover:bg-accent-neon/10 transition-all opacity-0 group-hover:opacity-100">
                                    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="m9 18 6-6-6-6" /></svg>
                                </button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
};

export default CustomerTable;
