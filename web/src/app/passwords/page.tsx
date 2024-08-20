import PasswordForm from "@/components/PasswordForm"
import getUserAction from "../actions/getUserAction"
import PasswordList from "@/components/PasswordList"

type Params = {
  searchParams: {
    id: string
  }
}

const PasswordsPage = async ({ searchParams }: Params) => {
  const userId = searchParams.id

  const userData = await getUserAction(userId)


  return (
    <main className="z-10 gap-10 flex flex-col ">
      <h1 className="md:text-6xl text-3xl text-center font-bold bg-clip-text text-transparent bg-gradient-to-b from-neutral-200 to to-neutral-600 font-sans">¡Escribe una Contraseña!</h1>


      <PasswordForm {...userData!} />


      <PasswordList url={`http://localhost:5000/passwords/all/${userId}`} />
    </main>
  )
}

export default PasswordsPage 
