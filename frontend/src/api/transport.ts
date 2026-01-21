import { createConnectTransport } from "@connectrpc/connect-web";

// Normally this would come from an environment variable
const BASE_URL = "http://localhost:50051";

export const transport = createConnectTransport({
    baseUrl: BASE_URL,
});
