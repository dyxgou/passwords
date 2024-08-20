"use client"
import { User } from "@/states/storage";
import { BackgroundGradient } from "./ui/background-gradient";

interface UserCard extends User {
  index: number
}

const UserCard = ({ id, name, passwords, points, index }: UserCard) => {
  return (
    <div>
      <BackgroundGradient className="justify-around items-center cursor-pointer gap-2 rounded-[22px] p-5 flex bg-zinc-900 text-white">
        <p className="font-bold text-3xl">{index + 1}</p>
        <p className="font-bold text-xl">{name}</p>
        <section className="flex flex-col items-center">
          <span className="text-blue-600">Contraseñas :</span>
          <span className="font-bold text-xl">{passwords}</span>
        </section>
        <section className="flex flex-col items-center">
          <span className="text-blue-600">Puntos:</span>
          <span className="font-bold text-xl">{points}</span>
        </section>
      </BackgroundGradient>
    </div>
  );
}

export default UserCard
