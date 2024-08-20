"use client"

import PasswordLink from "@/components/PasswordLink"
import PasswordList from "@/components/PasswordList"

const TopPage = () => {
  return (
    <main className="z-10 gap-10 flex flex-col ">
      <h1 className="md:text-6xl text-3xl text-center font-bold bg-clip-text text-transparent bg-gradient-to-b from-neutral-200 to to-neutral-600 font-sans">¡Todas las contraseñas!</h1>
      <PasswordLink href="/">Regresar al menu</PasswordLink>
      <PasswordLink href="/top/users">Ver los usuarios con mas puntos</PasswordLink>

      <PasswordList url="http://localhost:5000/passwords/all" />
    </main>
  )
}

export default TopPage
