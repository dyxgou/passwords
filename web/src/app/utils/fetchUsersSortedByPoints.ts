import { User } from "@/states/storage";

const fetchUsersSortedByPoints = async (url: string) => {
  const res = await fetch(url, {
    method: "GET",
  });

  const passwordsData = (await res.json()) as User[];

  return passwordsData;
};

export default fetchUsersSortedByPoints;
