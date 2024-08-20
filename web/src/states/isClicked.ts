import { create } from "zustand";

type IsClicked = {
  clicked: number[];
  setIsClicked: (passwordId: number) => void;
};

export const useIsClicked = create<IsClicked>()((set) => ({
  clicked: [],
  setIsClicked: (passwordId) =>
    set((state) => ({ clicked: [...state.clicked, passwordId] })),
}));
