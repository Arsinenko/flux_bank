import React from 'react';

const MainLayout: React.FC<{ children: React.ReactNode }> = ({ children }) => {
    return (
        <div className="min-h-screen bg-background-deep text-text-secondary flex">
            {/* Sidebar */}
            <aside className="w-64 bg-surface-card border-r border-surface-border flex flex-col">
                <div className="p-8">
                    <h1 className="text-2xl font-bold text-accent-neon tracking-tight flex items-center gap-2">
                        <div className="w-8 h-8 bg-accent-neon rounded-lg shadow-neon flex items-center justify-center text-background-deep font-black">F</div>
                        FluxBank
                    </h1>
                </div>

                <nav className="flex-1 px-4 space-y-2">
                    <div className="px-4 mb-4 text-text-muted font-bold uppercase text-[10px] tracking-widest">Analytics Console</div>
                    <NavItem icon={<DashboardIcon />} label="Dashboards" active />
                    <NavItem icon={<AccountsIcon />} label="Accounts" />
                    <NavItem icon={<TransactionsIcon />} label="Transactions" />
                    <NavItem icon={<CustomersIcon />} label="Customers" />
                    <NavItem icon={<InfrastructureIcon />} label="Infrastructure" />
                </nav>

                <div className="p-4 border-t border-surface-border">
                    <div className="flex items-center gap-3 p-3 rounded-xl bg-background-dark/50">
                        <div className="w-10 h-10 rounded-full bg-accent-neon/20 border border-accent-neon/30 overflow-hidden">
                            <img src="https://ui-avatars.com/api/?name=Admin&background=00FFC2&color=0B0C10" alt="Avatar" />
                        </div>
                        <div className="flex-1 min-w-0">
                            <p className="text-sm font-bold text-text-primary truncate">Admin User</p>
                            <p className="text-xs text-text-muted truncate">System Architect</p>
                        </div>
                    </div>
                </div>
            </aside>

            {/* Main Content */}
            <main className="flex-1 flex flex-col h-screen overflow-hidden">
                <header className="h-20 bg-background-deep/80 backdrop-blur-md border-b border-surface-border flex items-center justify-between px-8 flex-shrink-0 z-10">
                    <div className="flex items-center gap-4">
                        <h2 className="text-xl font-bold text-text-primary tracking-tight">Overview Dashboard</h2>
                        <div className="bg-accent-neon/10 text-accent-neon px-3 py-1 rounded-full text-[10px] font-bold uppercase tracking-widest border border-accent-neon/20">
                            Live Status
                        </div>
                    </div>

                    <div className="flex items-center gap-6">
                        <div className="hidden md:flex items-center gap-2 bg-background-dark/50 border border-surface-border rounded-xl px-4 py-2">
                            <SearchIcon />
                            <input
                                type="text"
                                placeholder="Search analytics..."
                                className="bg-transparent border-none outline-none text-sm text-text-secondary placeholder:text-text-muted w-48"
                            />
                        </div>

                        <div className="flex items-center gap-4">
                            <button className="text-text-muted hover:text-accent-neon transition-colors relative p-2 rounded-xl hover:bg-surface-border/30">
                                <NotificationIcon />
                                <span className="absolute top-2 right-2 w-2 h-2 bg-accent-neon rounded-full shadow-neon"></span>
                            </button>
                            <div className="h-8 w-px bg-surface-border"></div>
                            <div className="flex items-center gap-3 text-text-primary font-bold bg-surface-card border border-surface-border rounded-xl px-4 py-2">
                                <CalendarIcon />
                                <span className="text-sm">Mar 21, 2024</span>
                            </div>
                        </div>
                    </div>
                </header>

                <div className="flex-1 overflow-y-auto p-8 lg:p-12 bg-background-deep scrollbar-hide">
                    <div className="w-full">
                        {children}
                    </div>
                </div>
            </main>
        </div>
    );
};

/* Internal Helper Components */

const NavItem = ({ icon, label, active = false }: { icon: React.ReactNode, label: string, active?: boolean }) => (
    <a href="#" className={`flex items-center gap-4 p-3 rounded-xl transition-all duration-200 group ${active ? 'bg-accent-neon/10 text-accent-neon border border-accent-neon/20 shadow-neon opacity-100' : 'hover:bg-surface-border/50 text-text-muted hover:text-text-primary'}`}>
        <span className={`${active ? 'text-accent-neon' : 'text-text-muted group-hover:text-text-primary'} transition-colors`}>{icon}</span>
        <span className="font-semibold">{label}</span>
    </a>
);

/* Icons */
const DashboardIcon = () => <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><rect x="3" y="3" width="7" height="7" /><rect x="14" y="3" width="7" height="7" /><rect x="14" y="14" width="7" height="7" /><rect x="3" y="14" width="7" height="7" /></svg>;
const AccountsIcon = () => <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2" /><circle cx="9" cy="7" r="4" /><path d="M22 21v-2a4 4 0 0 0-3-3.87" /><path d="M16 3.13a4 4 0 0 1 0 7.75" /></svg>;
const TransactionsIcon = () => <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><polyline points="22 7 13.5 15.5 8.5 10.5 2 17" /><polyline points="16 7 22 7 22 13" /></svg>;
const CustomersIcon = () => <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M12 2H2v10c0 5.52 4.48 10 10 10s10-4.48 10-10V2H12z" /><path d="M12 22V2" /><path d="M2 12h20" /><path d="m12 12 4 4" /><path d="m12 12-4 4" /><path d="m12 12 4-4" /><path d="m12 12-4-4" /></svg>;
const InfrastructureIcon = () => <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><rect x="2" y="2" width="20" height="8" rx="2" ry="2" /><rect x="2" y="14" width="20" height="8" rx="2" ry="2" /><line x1="6" y1="6" x2="6.01" y2="6" /><line x1="6" y1="18" x2="6.01" y2="18" /></svg>;
const SearchIcon = () => <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><circle cx="11" cy="11" r="8" /><line x1="21" y1="21" x2="16.65" y2="16.65" /></svg>;
const NotificationIcon = () => <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9" /><path d="M13.73 21a2 2 0 0 1-3.46 0" /></svg>;
const CalendarIcon = () => <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><rect x="3" y="4" width="18" height="18" rx="2" ry="2" /><line x1="16" y1="2" x2="16" y2="6" /><line x1="8" y1="2" x2="8" y2="6" /><line x1="3" y1="10" x2="21" y2="10" /></svg>;

export default MainLayout;
