import updateUserPoints, { UpdateUserBody } from "@/app/utils/updateUserPoints"
import { useIsClicked } from "@/states/isClicked"
import useSWRMutation from "swr/mutation"

type SecurityButton = Readonly<{
  children: React.ReactNode
  securityLevel: string
  passwordId: number,
}>


const SecurityButton = ({ children, securityLevel, passwordId }: SecurityButton) => {
  const isClicked = useIsClicked(state => state.clicked.some(id => id === passwordId))
  const setIsClicked = useIsClicked(state => state.setIsClicked)

  const { trigger, isMutating, error } = useSWRMutation("http://localhost:5000/users/points", updateUserPoints)

  const handlerUpdate = async () => {
    const body: UpdateUserBody = {
      password_id: passwordId,
      security_level: securityLevel
    }
    console.log({ body })

    const data = await trigger(body)
    console.log({ buttonData: data })
    setIsClicked(passwordId)
  }


  const levels = {
    terrible: "text-red-800",
    bad: "text-red-500",
    mid: "text-blue-500",
    good: "text-green-300",
    excellent: "text-green-500"
  } as const
  return (
    <button
      onClick={handlerUpdate}
      disabled={isClicked}
      className={`${levels[securityLevel]} bg-[#0d0d0d] p-2 rounded cursor-pointer hover:bg-violet-900 transition z-50`}
    >
      {children}
    </button>
  )

}

export default SecurityButton
