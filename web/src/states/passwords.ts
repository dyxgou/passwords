import { create } from "zustand";
import { createJSONStorage, devtools, persist } from "zustand/middleware";

export type Password = {
  id: number;
  content: string;
  rot: string;
  sha: string;
  security_level: string;
  user_id: number;
};

interface UserWithMetods extends Password {
  insertPassword: (password: Password) => void;
  reset: () => void;
}

const defaultPassword: Password = {
  id: -1,
  content: "",
  rot: "",
  sha: "",
  security_level: "pending",
  user_id: -1,
};

export const useUser = create<UserWithMetods>()(
  devtools(
    persist(
      (set) => ({
        ...defaultPassword,
        insertPassword: (password: Password) => set(() => ({ ...password })),
        reset: () => ({ ...defaultPassword }),
      }),
      {
        name: "passwords",
        storage: createJSONStorage(() => localStorage),
      },
    ),
  ),
);
