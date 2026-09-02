import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { ClerkProvider } from "@clerk/clerk-react";

import "./index.css";
import App from "./App";
import { ThemeProvider } from "./theme/ThemeProvider";

const publishableKey = import.meta.env.VITE_CLERK_PUBLISHABLE_KEY;

const app = (
  <ThemeProvider>
    <App />
  </ThemeProvider>
);

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    {publishableKey ? (
      <ClerkProvider publishableKey={publishableKey}>{app}</ClerkProvider>
    ) : (
      app
    )}
  </StrictMode>
);