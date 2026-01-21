import { useQuery } from '@tanstack/react-query';
import { useApi } from '../providers/ApiProvider';

export const useTransactionAnalytics = () => {
    const { transactionClient } = useApi();

    return useQuery({
        queryKey: ['transaction-analytics', 'count'],
        queryFn: async () => {
            const response = await transactionClient.processGetCount({});
            return response;
        },
    });
};
