import { useFormStatus } from "react-dom"
import { BackgroundGradient } from "./ui/background-gradient"


type ButtonProps = Readonly<{
  text : string
}>

const SendButton = ({ text } : ButtonProps) => {
  const { pending } = useFormStatus()

  return (
    <BackgroundGradient className="rounded p-2 grid place-items-center cursor-pointer">
      <button className="text-white font-bold text-md" type="submit">
        {
          pending ? "Cargando" : text
        }
      </button>
    </BackgroundGradient>
  )
}

export default SendButton
