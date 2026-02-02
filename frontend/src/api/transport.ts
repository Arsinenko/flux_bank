import { createConnectTransport } from "@connectrpc/connect-web";

// Normally this would come from an environment variable
const BASE_URL = "http://localhost:localhost:8080";

export const transport = createConnectTransport({
    baseUrl: BASE_URL,
});
