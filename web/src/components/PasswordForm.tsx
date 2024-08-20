"use client"

import { useUser } from "@/states/user";
import { User } from "@/states/storage";
import SendButton from "./SendButtton";
import createPasswordAction from "@/app/actions/createPasswordAction";
import { FormEvent, useEffect, useState } from "react";
import { redirect } from "next/navigation";
import PasswordLink from "./PasswordLink";
import useSWR from "swr";
import fetchUserPoints from "@/app/utils/fetchUserPoints";

const PasswordForm = (userData: User) => {
  const [input, setInput] = useState("")
  const setUser = useUser((state) => state.setUser)
  setUser(userData)

  const name = useUser(state => state.name)
  const userId = useUser(state => state.id)
  const { data: points } = useSWR(`http://localhost:5000/users/points/${userId}`, fetchUserPoints, { refreshInterval: 1000 })

  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    setInput("")
  }


  useEffect(() => {
    if (userId === -1 || !userId) {
      redirect("/")
    }
  }, [userId])

  return (
    <form
      action={(data) => createPasswordAction(data, userId)}
      onSubmit={handleSubmit}
      className="flex flex-col justify-center items-center p-1.5 gap-5 w-full"
    >
      <h2 className="text-white font-sans text-3xl font-bold">Bienvenido <span className="text-blue-400">{name}</span>
        <p className="text-gray-800">Tienes {points} puntos</p>
      </h2>
      <input
        type="text"
        value={input}
        className="input md:w-1/2"
        placeholder="Escribe tu contraseña"
        onChange={e => setInput(e.target.value)}
        name="password"
        minLength={5}
        required
      />

      <div className="flex gap-4">
        <SendButton text="Enviar contraseña" />
        <PasswordLink href="/top">Ver Top!</PasswordLink>
      </div>

    </form>
  );
};

export default PasswordForm;
