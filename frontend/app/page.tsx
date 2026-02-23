"use client";

import { HeadingText } from "@/components/common/Texts";
import Console from "@/components/app/Console";
import Profiles from "@/components/app/Profiles";
import useAxios from "@/hooks/useAxios";
import { useEffect, useState } from "react";
import { Me } from "./types";

export default function Home() {
  const axios = useAxios();
  const [me, setMe] = useState<Me | null>(null);

  useEffect(() => {
    const getMe = async () => {
      const res = await axios("/me");
      setMe(res.data);
    };
    getMe();
  }, [axios]);

  return (
    <div>
      <div className='my-4'>
        <HeadingText name>Hi, {me?.name}.</HeadingText>
      </div>
      <div className='my-8 flex gap-4'>
        <Profiles />
        <Console />
      </div>
    </div>
  );
}
