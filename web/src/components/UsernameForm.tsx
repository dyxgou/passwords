import createUserAction from "@/app/actions/createUserAction";
import SendButton from "./SendButtton";
import { useLayoutEffect, useState } from "react";
import { useSearchParams } from "next/navigation";

const UsernameForm = () => {
  const [isInvalid, setIsInvalid] = useState(false)
  const searchParams = useSearchParams()
  const invalidInput = searchParams.get("invalid")

  useLayoutEffect(() => {
    if (invalidInput) {
      setIsInvalid(true)
    }
  }, [invalidInput])

  return (
    <form
      action={createUserAction}
      className="flex flex-col justify-center p-1.5 gap-5 md:w-1/2"
    >
      <input
        type="text"
        placeholder="Escribe tu nombre"
        className={`input ${isInvalid ? "invalid" : ""}`}
        name="name"
        minLength={5}
        required
      />
      {isInvalid && <span className="font-sans text-red-500">El nombre de usuario es invalido</span>}
      <SendButton text="Registrame" />
    </form>
  );
};

export default UsernameForm;
