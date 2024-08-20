"use client"
import { Password } from "@/states/passwords";
import { BackgroundGradient } from "./ui/background-gradient";
import { MouseEvent, useState } from "react";
import { ShieldX, ShieldBan, ShieldCheck, ShieldPlus, Shield, } from "lucide-react";
import SecurityButton from "./SecurityButton";

const PasswordItem = ({ id, content, rot, sha, user_id, security_level }: Password) => {
  const [isShowed, setIsShowed] = useState(false)
  const handleIsShowed = (e: MouseEvent<HTMLDivElement, MouseEvent>) => {
    // e.preventDefault()
    setIsShowed(prev => !prev)
  }

  return (
    <div onClick={handleIsShowed}>
      <BackgroundGradient className="text-clip cursor-pointer flex-col gap-2 rounded-[22px] p-5 flex bg-zinc-900 text-white">
        <span className="font-bold">{content}</span>

        {
          isShowed && <span className="text-ellipsis">ROT: {rot}</span>
        }
        {
          isShowed && <span className="">SHA: {sha}</span>
        }
        {
          isShowed &&
          <div className="flex items-center gap-2 w-full justify-around transition">
            <SecurityButton securityLevel="terrible" passwordId={id}>
              <ShieldBan />
            </SecurityButton>
            <SecurityButton securityLevel="bad" passwordId={id}>
              <ShieldX />
            </SecurityButton>
            <SecurityButton securityLevel="mid" passwordId={id}>
              <Shield />
            </SecurityButton>
            <SecurityButton securityLevel="good" passwordId={id}>
              <ShieldCheck />
            </SecurityButton>
            <SecurityButton securityLevel="excellent" passwordId={id}>
              <ShieldPlus />
            </SecurityButton>
          </div>
        }



      </BackgroundGradient>
    </div>
  );
}

export default PasswordItem
