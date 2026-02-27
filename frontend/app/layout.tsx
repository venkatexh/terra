import "./globals.css";
import { Geist, Geist_Mono } from "next/font/google";
import { ModalProvider } from "@/contexts/modal-context";
import PrimaryNavbar from "@/components/nav/PrimaryNavbar";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang='en'>
      <body
        className={`w-1/2 h-auto mx-auto ${geistSans.variable} ${geistMono.variable} antialiased py-12`}>
        <ModalProvider>
          <PrimaryNavbar />
          <div className='mt-12'>{children}</div>
        </ModalProvider>
      </body>
    </html>
  );
}
