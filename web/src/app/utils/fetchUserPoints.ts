import { User } from "@/states/storage";

const fetchUserPoints = async (url: string) => {
  const res = await fetch(url, {
    method: "GET",
  });

  const userData = (await res.json()) as User;

  return userData.points;
};

export default fetchUserPoints;
