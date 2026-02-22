"use client";

import { useState } from "react";
import useAxios from "@/hooks/useAxios";

export default function NewClientPage() {
  const axios = useAxios();
  const [name, setName] = useState("");
  const [uris, setUris] = useState("");
  const [result, setResult] = useState(null);

  async function submit() {
    const redirectUris = uris.split("\n").map((u) => u.trim());

    const res = await axios.post("/oauth/clients", {
      body: JSON.stringify({
        name,
        redirect_uris: redirectUris,
      }),
    });

    setResult(res.data);
  }

  return (
    <div className='flex flex-col'>
      <div className='text-2xl'>Create OAuth app</div>
      <form className='flex flex-col py-12 gap-4'>
        <input
          placeholder='App Name'
          value={name}
          className='p-2 focus:outline-none border-gray-300 border rounded-3xl'
          onChange={(e) => setName(e.target.value)}
        />
        <input
          placeholder='Redirect URI'
          value={uris}
          className='p-2 focus:outline-none border-gray-300 border rounded-3xl'
          onChange={(e) => setUris(e.target.value)}
        />
        <button
          className='py-2 bg-orange-600 hover:bg-orange-700 rounded-3xl cursor-pointer'
          onClick={submit}>
          Create
        </button>
      </form>

      {result && (
        <div>
          <h2>Credentials</h2>
          <p>Client ID:</p>
          <p>Client Secret:</p>
        </div>
      )}
    </div>
  );
}
