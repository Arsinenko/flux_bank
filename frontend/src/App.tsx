import MainLayout from './components/MainLayout';
import { useTransactionAnalytics } from './hooks/useTransactionAnalytics';

function App() {
  const { data: transactionCount, isLoading: isTransactionsLoading } = useTransactionAnalytics();

  return (
    <MainLayout>
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div className="bg-slate-800 p-6 rounded-xl border border-slate-700 shadow-xl">
          <h3 className="text-slate-400 text-sm font-medium mb-2">Total Transactions</h3>
          <p className="text-3xl font-bold">
            {isTransactionsLoading ? '...' : transactionCount?.count?.toString() || '0'}
          </p>
          <div className="mt-4 text-green-400 text-sm">↑ Real-time from gRPC</div>
        </div>
        <div className="bg-slate-800 p-6 rounded-xl border border-slate-700 shadow-xl">
          <h3 className="text-slate-400 text-sm font-medium mb-2">Average Fee</h3>
          <p className="text-3xl font-bold">$4.20</p>
          <div className="mt-4 text-red-400 text-sm">↓ 5% from last month</div>
        </div>
        <div className="bg-slate-800 p-6 rounded-xl border border-slate-700 shadow-xl">
          <h3 className="text-slate-400 text-sm font-medium mb-2">Active Users</h3>
          <p className="text-3xl font-bold">1,240</p>
          <div className="mt-4 text-green-400 text-sm">↑ 3% from last week</div>
        </div>
      </div>

      <div className="mt-8 bg-slate-800 p-8 rounded-xl border border-slate-700 shadow-xl h-96 flex items-center justify-center">
        <p className="text-slate-500 italic">Chart visualization will be here...</p>
      </div>
    </MainLayout>
  );
}

export default App;
