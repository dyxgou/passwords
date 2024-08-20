import SuperJSON from "superjson";
import { PersistStorage } from "zustand/middleware";

export type User = {
  id: number;
  name: string;
  passwords: number;
  points: number;
};

export const storage: PersistStorage<User> = {
  getItem: (name) => {
    const str = localStorage.getItem(name);
    if (!str) return null;
    return SuperJSON.parse(str);
  },
  setItem: (name, value) => {
    localStorage.setItem(name, SuperJSON.stringify(value));
  },
  removeItem: (name) => localStorage.removeItem(name),
};
