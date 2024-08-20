"use client"

import UsernameForm from "@/components/UsernameForm"
import { useUser } from "@/states/user"
import { redirect } from "next/navigation"
import { useEffect } from "react"

const Page = () => {
  const userId = useUser((state) => state.id)

  console.log(userId)

  useEffect(() => {
    if (userId !== -1) {
      redirect(`/passwords?id=${userId}`)
    }
  }, [userId])

  return (
    <div className="h-full w-full grid place-items-center">
      <main className="z-10 gap-10 flex flex-col justify-center items-center w-3/5">
        <h1 className="md:text-7xl text-4xl text-center font-bold bg-clip-text text-transparent bg-gradient-to-b from-neutral-200 to to-neutral-600 font-sans">Descubre que tan segura es tu contraseña</h1>

        <UsernameForm />
      </main>
    </div>
  )
}

export default Page
