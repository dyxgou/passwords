import { Password } from "@/states/passwords";

const fetchUserPasswords = async (url: string) => {
  const res = await fetch(url, {
    method: "GET",
  });

  const passwordsData = (await res.json()) as Password[];

  return passwordsData;
};

export default fetchUserPasswords;
