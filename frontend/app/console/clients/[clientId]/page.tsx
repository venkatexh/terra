"use client";

import useAxios from "@/hooks/useAxios";
import { useParams } from "next/navigation";
import { useEffect, useState } from "react";

import { Client, UpdateRedirectURI } from "./types";

import {
  DefaultText,
  SmallText,
  SubheadingText,
} from "@/components/common/Texts";
import { DefaultButton } from "@/components/common/Buttons";

export default function ClientPage() {
  const axios = useAxios();
  const params = useParams();

  const [newURIInput, setNewURIInput] = useState<number[]>([]);
  const [client, setClient] = useState<Client | null>(null);
  const [redirectURIs, setRedirectURIs] = useState<UpdateRedirectURI[]>([]);

  useEffect(() => {
    const fetchClient = async () => {
      try {
        const res = await axios(`/oauth/clients/${params.clientId}`);
        setClient(res.data);
        setRedirectURIs(
          res.data.redirectUris.map((uri: UpdateRedirectURI) => ({
            id: uri.id,
            uri: uri.uri,
          })),
        );
      } catch (err) {
        console.error(err);
      }
    };

    fetchClient();
  }, [axios, params.clientId]);

  const handleRedirectURIs = (id: string, value: string) => {
    setRedirectURIs((prev) => [
      ...prev.filter((uri) => uri.id !== id),
      { id, uri: value },
    ]);
  };

  const handleNewUriClick = () => {
    setNewURIInput((prev) => [...prev, prev.length]);
    console.log(newURIInput);
  };

  return (
    <div>
      <SubheadingText name>{client?.name}</SubheadingText>
      <div className='flex border border-gray-500 border-l-0 border-r-0 py-4 my-4'>
        <div className='w-1/2'>
          <SmallText>Client ID</SmallText>
          <DefaultText>{client?.clientId}</DefaultText>
        </div>
        <div className='w-1/2'>
          <SmallText>Created at</SmallText>
          <DefaultText>{client?.createdAt}</DefaultText>
        </div>
      </div>
      <DefaultText className='py-2' name>
        Redirect URIS
      </DefaultText>
      <div className='flex flex-col'>
        {redirectURIs.map((uri: UpdateRedirectURI, idx) => (
          <div className='py-2' key={uri.id}>
            <SmallText>URI {idx + 1}</SmallText>
            <input
              type='text'
              value={uri.uri}
              onChange={(e) => handleRedirectURIs(uri.id, e.target.value)}
              className='w-full px-2 py-1 border border-gray-500 focus:outline-none rounded'
            />
          </div>
        ))}
        {newURIInput.map((_, idx) => (
          <div className='py-2' key={idx}>
            <SmallText>New URI {idx + 1}</SmallText>
            <input
              type='text'
              placeholder='https://www.example.com/callback'
              className='w-full px-2 py-1 border border-gray-500 focus:outline-none rounded'
            />
          </div>
        ))}
        <DefaultButton
          transparent
          handleButtonClick={() => handleNewUriClick()}
          className='my-4 flex items-center justify-center'>
          <span className='text-xl pr-2'>+</span> Add new URI
        </DefaultButton>
      </div>
      <div className='w-full flex justify-end'>
        <DefaultButton className='my-4' handleButtonClick={() => {}}>
          Save changes
        </DefaultButton>
      </div>
    </div>
  );
}
