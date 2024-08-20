"use client"

import fetchUsersSortedByPoints from "@/app/utils/fetchUsersSortedByPoints";
import PasswordLink from "@/components/PasswordLink"
import UserCard from "@/components/UserCard";
import useSWR from "swr";

const UsersTop = () => {
  const { data: users } = useSWR(`http://localhost:5000/users/all/points`, fetchUsersSortedByPoints, { refreshInterval: 2000 });

  return (
    <main className="z-10 gap-10 flex flex-col ">
      <h1 className="md:text-6xl text-3xl text-center font-bold bg-clip-text text-transparent bg-gradient-to-b from-neutral-200 to to-neutral-600 font-sans">¡Usuarios en el Top!</h1>
      <PasswordLink href="/top">Regresar al top de Contrseñas</PasswordLink>

      {
        users?.map((user, index) => {
          return <UserCard key={index} {...user} index={index} />
        })
      }

    </main>
  )
}

export default UsersTop
