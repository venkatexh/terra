import { useMemo } from "react";
import axiosInstance from "@/lib/axios";

export default function useAxios() {
  return useMemo(() => axiosInstance, []);
}
