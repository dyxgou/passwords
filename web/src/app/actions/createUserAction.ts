import { redirect } from "next/navigation";

type UserId = {
  user_id: number;
};

async function createUserAction(data: FormData) {
  const name = data.get("name");

  const res = await fetch("http://localhost:5000/users/create", {
    method: "POST",
    body: JSON.stringify({ name }),
  });

  const USER_FOUND = 200;
  if (res.status !== USER_FOUND) {
    const revalidateRes = await fetch(
      `http://localhost:5000/users/id/${name}`,
      {
        method: "GET",
      },
    );

    if (revalidateRes.status !== USER_FOUND) {
      redirect("/?invalid=true");
    }

    const userData = (await revalidateRes.json()) as UserId;

    redirect(`/passwords?id=${userData.user_id}`);
  }

  const userData = (await res.json()) as UserId;

  redirect(`/passwords?id=${userData.user_id}`);
}

export default createUserAction;
