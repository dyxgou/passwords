import { redirect } from "next/navigation";

const createPasswordAction = async (formData: FormData, user_id: number) => {
  const password = formData.get("password");

  const res = await fetch("http://localhost:5000/passwords/create", {
    method: "POST",
    body: JSON.stringify({ password, user_id }),
  });

  console.log({ res });

  const PASSWORD_CREATED = 200;

  if (res.status !== PASSWORD_CREATED) {
    redirect(`/passwords?id=${user_id}&invalid=true`);
  }
};

export default createPasswordAction;
