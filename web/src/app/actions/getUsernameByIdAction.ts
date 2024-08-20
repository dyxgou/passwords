const getUsernameById = async (userId: string) => {
  const res = await fetch(`http://localhost:5000/users/nameid/${userId}`, {
    method: "GET",
  });

  const PASSWORDS_FOUND = 200;
  if (res.status !== PASSWORDS_FOUND) {
    return null;
  }

  type response = {
    user_name: string;
  };

  const userData = (await res.json()) as response;

  return userData.user_name;
};

export default getUsernameById;
