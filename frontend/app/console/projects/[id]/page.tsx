"use client";

import useAxios from "@/hooks/useAxios";
import { useParams } from "next/navigation";
import { useEffect, useState } from "react";

import {
  LargeText,
  DefaultText,
  SubheadingText,
} from "@/components/common/Texts";
import { DefaultButton } from "@/components/common/Buttons";
import ClientCard from "@/components/app/console/projects/[id]/ClientCard";

type Project = {
  id: string;
  name: string;
  description: string;
};

type Client = {
  id: string;
  name: string;
  clientId: string;
  clientSecret: string;
};

export default function ProjectPage() {
  const axios = useAxios();
  const params = useParams();

  const [project, setProject] = useState<Project>();
  const [clients, setClients] = useState<Client[]>([]);

  useEffect(() => {
    const fetchProject = async () => {
      try {
        const res = await axios(`/projects/${params.id}`);
        setProject(res.data);
      } catch (err) {
        console.error(err);
      }
    };

    const fetchProjectClients = async () => {
      try {
        const res = await axios(`/projects/${params.id}/clients`);
        setClients(res.data);
      } catch (err) {
        console.error(err);
      }
    };
    fetchProject();
    fetchProjectClients();
  }, [params.id, axios]);

  return (
    <div>
      <div>
        <div className='flex justify-between items-center'>
          <SubheadingText name>{project?.name}</SubheadingText>
          <DefaultButton handleButtonClick={() => {}}>+ New</DefaultButton>
        </div>
        <DefaultText>{project?.description}</DefaultText>
      </div>
      <LargeText className='py-4'>OAuth 2.0 clients</LargeText>
      <div>
        <div className='flex flex-col gap-4'>
          {clients.map(
            (client: {
              name: string;
              id: string;
              clientId: string;
              clientSecret: string;
            }) => (
              <ClientCard
                key={client.id}
                id={client.id}
                name={client.name}
                clientId={client.clientId}
                clientSecret={client.clientSecret}
              />
            ),
          )}
        </div>
      </div>
    </div>
  );
}
