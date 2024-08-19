"use client"

import PasswordForm from "@/components/PasswordFrom"

const Page = () => {
  return (
    <main className="z-10 gap-10 flex flex-col w-3/5 items-center">
      <h1 className="md:text-6xl text-4xl text-center font-bold bg-clip-text text-transparent bg-gradient-to-b from-neutral-200 to to-neutral-600 font-sans">Descubre que tan segura es tu contraseña</h1>

      <PasswordForm />
    </main>
  )
}

export default Page
