"use client";

import useAxios from "@/hooks/useAxios";
import { useEffect, useState } from "react";

import { baseURL } from "@/lib/constants/baseURL";

import GroupCard from "@/components/app/app-groups/GroupCard";

export default function ProfilesPage() {
  const axios = useAxios();
  const [groups, setGroups] = useState([]);

  useEffect(() => {
    const getData = async () => {
      try {
        const res = await axios.get(`${baseURL}/me/groups`);
        setGroups(res.data);
      } catch (err) {
        console.error(err);
      }
    };

    getData();
  }, [axios]);

  return (
    <div className='flex flex-col'>
      <div className='flex justify-between items-center'>
        <div className='text-2xl'>Your app groups</div>
        <button className='px-4 py-1 bg-orange-600 hover:bg-orange-700 rounded-sm'>
          + New
        </button>
      </div>
      <div className='py-12'>
        {groups.map((group: { name: string; id: string }) => (
          <GroupCard key={group.id} id={group.id} name={group.name} />
        ))}
      </div>
    </div>
  );
}
