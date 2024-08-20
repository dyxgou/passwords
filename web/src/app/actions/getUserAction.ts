import { User } from "@/states/storage";

const getUserAction = async (id: string) => {
  const res = await fetch(`http://localhost:5000/users/${id}`, {
    method: "GET",
  });
  const USER_FOUND = 200;

  if (res.status !== USER_FOUND) {
    return null;
  }

  const userData = (await res.json()) as User;

  return userData;
};

export default getUserAction;
