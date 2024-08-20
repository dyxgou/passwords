"use client"

import { BackgroundGradient } from "./ui/background-gradient"
import { LogOut } from "lucide-react"
import Link from "next/link"


type PasswordLink = {
  href: string
  children: React.ReactNode
}

const PasswordLink = ({ href, children }: PasswordLink) => {
  return (
    <BackgroundGradient className="rounded grid place-items-center cursor-pointer">
      <Link
        href={href} className="text-white font-bold gap-3 flex items-center justify-center text-md w-full h-full p-2"
      >
        <LogOut />
        {children}
      </Link>
    </BackgroundGradient>
  )
}

export default PasswordLink
