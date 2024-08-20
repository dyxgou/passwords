import { create } from "zustand";
import { createJSONStorage, devtools, persist } from "zustand/middleware";
import { User } from "./storage";

interface UserWithMetods extends User {
  setUser: (user: User) => void;
  logout: () => void;
}

const defaultUser: User = {
  id: -1,
  name: "nouser",
  passwords: -1,
  points: -1,
};

export const useUser = create<UserWithMetods>()(
  devtools(
    persist(
      (set) => ({
        ...defaultUser,
        setUser: (user) => set(() => ({ ...user })),
        logout: () => set(() => ({ id: -1 })),
      }),
      {
        name: "user",
        storage: createJSONStorage(() => localStorage),
      },
    ),
  ),
);
