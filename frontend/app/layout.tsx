import type { Metadata } from "next";
import "./globals.css";
import ServerUserProvider from "./providers/serverUserProvider";

export const metadata: Metadata = {
  title: "Workflow-App",
  description: "Workflow magentment items app",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        <ServerUserProvider>
          {children}
        </ServerUserProvider>
      </body>
    </html>
  );
}
