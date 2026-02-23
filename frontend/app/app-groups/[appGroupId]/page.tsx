"use client";

import AuthorizedAppCard from "@/components/app/app-groups/appGroupId/AuthorizedAppCard";
import useAxios from "@/hooks/useAxios";
import { baseURL } from "@/lib/constants/baseURL";
import { useParams } from "next/navigation";
import { useEffect, useState } from "react";

export default function AppGroupPage() {
  const axios = useAxios();
  const params = useParams();
  const { appGroupId } = params;

  const [group, setGroup] = useState(null);
  const [authorizations, setAuthorizations] = useState([]);

  useEffect(() => {
    const getData = async () => {
      try {
        const res = await axios.get(`${baseURL}/me/groups/${appGroupId}`);
        setGroup(res.data);
        const authorizations = await axios.get(
          `${baseURL}/authorizations?groupId=${appGroupId}`,
        );
        setAuthorizations(authorizations.data);
      } catch (err) {
        console.error(err);
      }
    };

    getData();
  }, [appGroupId, axios]);

  return (
    <div>
      <div className='pb-4'>
        <div className='text-xs'>App group</div>
        <div className='text-xl'>{group?.name}</div>
      </div>

      <div>
        <div>Apps with access</div>
        <div className='py-4 flex flex-col gap-4'>
          {authorizations.map((a: { clientName: string; id: string }) => (
            <AuthorizedAppCard key={a.id} id={a.id} name={a.clientName} />
          ))}
        </div>
      </div>
    </div>
  );
}
