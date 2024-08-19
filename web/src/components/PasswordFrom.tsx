import SendButton from "./SendButtton";
import submitPasswordForm from "@/app/actions/submitPasswordForm";

const PasswordForm = () => {
  return (
    <form
      action={submitPasswordForm}
      className="flex flex-col justify-center p-1.5 gap-5 md:w-1/2"
    >
      <input
        type="text"
        placeholder="Escribe tu nombre"
        className="input"
        name="name"
        minLength={5}
        required
      />
      <SendButton text="Registrame" />
    </form>
  );
};

export default PasswordForm;
