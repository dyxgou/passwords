import { useFormStatus } from "react-dom"
import { BackgroundGradient } from "./ui/background-gradient"


type ButtonProps = Readonly<{
  text: string
}>

const SendButton = ({ text }: ButtonProps) => {
  const { pending } = useFormStatus()

  return (
    <BackgroundGradient className="rounded grid place-items-center cursor-pointer">
      <button className="text-white font-bold text-md w-full h-full p-2" type="submit">
        {
          pending ? "Cargando" : text
        }
      </button>
    </BackgroundGradient>
  )
}

export default SendButton
