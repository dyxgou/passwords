export type UpdateUserBody = {
  password_id: number;
  security_level: string;
};

type Body = {
  arg: UpdateUserBody;
};

const updateUserPoints = async (url: string, { arg }: Body) => {
  const res = await fetch(url, {
    method: "PUT",
    body: JSON.stringify(arg),
  });

  if (!res.ok) {
    throw new Error("Invalid sercurity level");
  }

  return res.json();
};

export default updateUserPoints;
