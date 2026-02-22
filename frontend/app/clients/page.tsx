"use client";

import useAxios from "@/hooks/useAxios";
import Link from "next/link";
import { useEffect, useState } from "react";

export default function ClientsPage() {
  const axios = useAxios();
  const [clients, setClients] = useState([]);

  useEffect(() => {
    const getData = async () => {
      try {
        const res = await axios.get("/oauth/clients");
        setClients(res.data);
      } catch (err) {
        console.error(err);
      }
    };

    getData();
  }, []);

  return (
    <div className='flex flex-col'>
      <div className='flex justify-between items-center'>
        <div className='text-xl'>OAuth clients</div>
        <Link href='clients/new'>
          <button className='bg-orange-600 px-4 py-1 rounded-sm hover:bg-orange-700 cursor-pointer'>
            + New client
          </button>
        </Link>
      </div>
      <div>
        {clients.map((client: { name: string; id: string }) => (
          <div key={client.id}>{client.name}</div>
        ))}
      </div>
    </div>
  );
}
