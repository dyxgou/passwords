import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "Top",
  description: "Top de las contraseñas",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <div className="flex justify-center w-full h-full p-14">
      {children}
    </div>
  );
}
