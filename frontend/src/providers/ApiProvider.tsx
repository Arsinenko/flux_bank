import React, { createContext, useContext, useMemo } from 'react';
import { createPromiseClient, type PromiseClient } from '@connectrpc/connect';
import { transport } from '../api/transport';
import { TransactionAnalyticService } from '../gen/transaction_analitic_connect';
import { CustomerAnalyticService } from '../gen/customer_analytic_connect';

interface ApiContextType {
    transactionClient: PromiseClient<typeof TransactionAnalyticService>;
    customerClient: PromiseClient<typeof CustomerAnalyticService>;
}

const ApiContext = createContext<ApiContextType | null>(null);

export const ApiProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
    const clients = useMemo(() => ({
        transactionClient: createPromiseClient(TransactionAnalyticService, transport),
        customerClient: createPromiseClient(CustomerAnalyticService, transport),
    }), []);

    return (
        <ApiContext.Provider value={clients}>
            {children}
        </ApiContext.Provider>
    );
};

export const useApi = () => {
    const context = useContext(ApiContext);
    if (!context) {
        throw new Error('useApi must be used within an ApiProvider');
    }
    return context;
};
