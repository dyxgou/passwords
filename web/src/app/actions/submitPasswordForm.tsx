const submitPasswordForm = async (data: FormData) => {
  const name = data.get("name");

  const resBody = {
    name: name,
  } as const;

  console.log({ resBody });

  const res = await fetch("http://localhost:5000/users/create", {
    method: "POST",
    body: JSON.stringify(resBody),
  });

  console.log(res);
};

export default submitPasswordForm;
