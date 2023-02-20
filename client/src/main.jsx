import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter } from "react-router-dom";
import { QueryClient, QueryClientProvider } from "react-query";
import { AppProvider } from "@/context/AppContext";
import App from "./App";
import "./main.css";

const client = new QueryClient();

ReactDOM.createRoot(document.getElementById("root")).render(
	<React.StrictMode>
		<AppProvider>
			<QueryClientProvider client={client}>
				<BrowserRouter>
					<App />
				</BrowserRouter>
			</QueryClientProvider>
		</AppProvider>
	</React.StrictMode>
);
