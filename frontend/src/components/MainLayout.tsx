import React from 'react';

const MainLayout: React.FC<{ children: React.ReactNode }> = ({ children }) => {
    return (
        <div className="min-h-screen bg-slate-900 text-slate-100 flex">
            {/* Sidebar Placeholder */}
            <aside className="w-64 bg-slate-800 p-4 border-r border-slate-700">
                <h1 className="text-xl font-bold text-blue-400 mb-8">FluxBank Analytics</h1>
                <nav className="space-y-4">
                    <div className="text-slate-400 font-semibold uppercase text-xs">Menu</div>
                    <a href="#" className="block p-2 hover:bg-slate-700 rounded transition">Dashboard</a>
                    <a href="#" className="block p-2 hover:bg-slate-700 rounded transition text-blue-400 bg-slate-700">Transactions</a>
                    <a href="#" className="block p-2 hover:bg-slate-700 rounded transition">Customers</a>
                </nav>
            </aside>

            {/* Main Content */}
            <main className="flex-1 flex flex-col">
                <header className="h-16 bg-slate-800 border-b border-slate-700 flex items-center px-8">
                    <h2 className="text-lg font-medium">Analytics Overview</h2>
                </header>
                <section className="p-8">
                    {children}
                </section>
            </main>
        </div>
    );
};

export default MainLayout;
