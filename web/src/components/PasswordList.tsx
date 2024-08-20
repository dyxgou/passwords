"use client"
import PasswordItem from "./PasswordItem"
import { Suspense } from "react"
import { LoaderIcon } from "lucide-react"
import useSWR from "swr"
import fetchUserPasswords from "@/app/utils/fetchUserPasswords"


type PasswordProps = {
  url: string
}

const PasswordList = ({ url }: PasswordProps) => {
  // const { data: passwords } = useSWR(`http://localhost:5000/passwords/all/${userId}`, fetchUserPasswords, { refreshInterval: 2000 });
  const { data: passwords } = useSWR(url, fetchUserPasswords, { refreshInterval: 2000 });
  console.log(passwords)



  return (
    <Suspense fallback={<LoaderIcon />}>
      <div className="flex flex-col gap-4 sm:p-5">
        {
          passwords?.toReversed().map((password, index) => {
            return (
              <PasswordItem {...password} key={index} />
            )
          })
        }
      </div>
    </Suspense>
  )
}

export default PasswordList
