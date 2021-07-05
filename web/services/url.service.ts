import { toast } from "react-toastify";
import create from "zustand";
import { combine, devtools } from "zustand/middleware";
import { api } from "../utils/api";

export const useUrlService = create(
  devtools(
    combine({ shortId: null as string }, (set) => ({
      shortenUrl: async (url: string) => {
        try {
          const { data } = await api.post("/url", {
            url,
          });

          set(() => ({ shortId: data.shortId }));

          const initialShortIds: any =
            typeof window !== "undefined" &&
            JSON.parse(localStorage.getItem("shortIds"));

          typeof window !== "undefined" &&
            localStorage.setItem(
              "shortIds",
              JSON.stringify([...initialShortIds, data.shortId])
            );

          toast.success("Url has been shortened", {
            position: "top-center",
            autoClose: 3000,
          });
        } catch (error) {
          toast.error(error.response?.data.message, {
            position: "top-center",
            autoClose: 3000,
          });
        }
      },
    }))
  )
);
