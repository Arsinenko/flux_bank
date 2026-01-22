import MainLayout from './components/MainLayout';
import KPIWidget from './components/analytics/KPIWidget';
import { MOCK_DATA } from './utils/mockData';
import Card from './components/ui/Card';
import ActivityTrendChart from './components/analytics/ActivityTrendChart';
import CategoryDonutChart from './components/analytics/CategoryDonutChart';
import TransactionHeatmap from './components/analytics/TransactionHeatmap';

function App() {
  const { general, loansAndDeposits, alerts, transactions } = MOCK_DATA;

  return (
    <MainLayout>
      <div className="space-y-8">
        {/* KPI Row */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <KPIWidget
            title="Total Assets"
            value={general.totalAssets.toLocaleString('en-US', { maximumFractionDigits: 0 })}
            prefix="$"
            trend={general.assetTrend}
          />
          <KPIWidget
            title="Active Users"
            value={general.activeUsers.toLocaleString()}
            trend={general.userTrend}
          />
          <KPIWidget
            title="Daily Volume"
            value={general.dailyTransactionVolume.toLocaleString('en-US', { maximumFractionDigits: 0 })}
            prefix="$"
            trend={general.volumeTrend}
          />
        </div>

        {/* Charts and Analytics Grid */}
        <div className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-4 gap-6">
          {/* Main Charts Row */}
          <Card title="Acitvity Trend" className="xl:col-span-2 h-[480px]">
            <div className="h-[380px] w-full mt-4">
              <ActivityTrendChart data={loansAndDeposits.trend} />
            </div>
          </Card>

          <Card title="Transaction Density" className="h-[480px]">
            <div className="h-[380px] w-full mt-4">
              <TransactionHeatmap data={transactions.heatmap} />
            </div>
          </Card>

          <Card title="Spending Categories" className="h-[480px]">
            <div className="h-[380px] w-full mt-4">
              <CategoryDonutChart data={transactions.categories} />
            </div>
          </Card>

          {/* Infrastructure & Customers Row */}
          <Card title="Infrastructure Alerts" className="h-[450px]">
            <div className="space-y-3 mt-4 overflow-y-auto max-h-[350px] pr-2 scrollbar-hide">
              {alerts.map(alert => (
                <div key={alert.id} className="flex flex-col gap-1 p-4 rounded-xl bg-background-dark/50 border border-surface-border hover:border-accent-neon/30 transition-colors group">
                  <div className="flex items-center justify-between">
                    <span className={`text-[10px] font-bold px-2 py-0.5 rounded ${alert.level === 'HIGH' ? 'bg-status-error/20 text-status-error' : 'bg-status-warning/20 text-status-warning'}`}>
                      {alert.level}
                    </span>
                    <span className="text-text-muted text-[10px]">{alert.time}</span>
                  </div>
                  <span className="text-text-primary font-bold text-sm mt-1">{alert.title}</span>
                </div>
              ))}
            </div>
          </Card>

          <Card title="Latest Transactions" className="xl:col-span-3 h-[450px]">
            <div className="h-full flex flex-col items-center justify-center text-center p-8 text-text-muted">
              <div className="w-16 h-16 rounded-full border-2 border-dashed border-accent-neon/20 mb-4 animate-pulse"></div>
              <p className="font-bold text-text-primary">Transactions Ledger</p>
              <p className="text-sm">Real-time ledger stream will appear here in the next update.</p>
            </div>
          </Card>
        </div>
      </div>
    </MainLayout>
  );
}

export default App;
