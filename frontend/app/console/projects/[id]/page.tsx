"use client";

import useAxios from "@/hooks/useAxios";
import { useParams } from "next/navigation";
import { useEffect, useState } from "react";
import { useModal } from "@/contexts/modal-context";

import {
  LargeText,
  DefaultText,
  SubheadingText,
} from "@/components/common/Texts";
import { DefaultButton } from "@/components/common/Buttons";
import ClientCard from "@/components/app/console/projects/[id]/ClientCard";
import NewClientModal from "@/components/app/console/projects/[id]/NewClientModal";

import { Project } from "@/types/console/projects/Project";
import { Client } from "@/types/console/projects/[id]/Client";
import { ClientForm } from "@/types/console/projects/[id]/ClientForm";

export default function ProjectPage() {
  const axios = useAxios();
  const params = useParams();
  const { openModal, closeModal } = useModal();

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

  const createNewClient = async (formData: ClientForm) => {
    try {
      const res = await axios.post(`/projects/${params.id}/clients`, {
        name: formData.name,
        redirectUris: formData.redirectUris,
      });
      closeModal();
      setClients([...clients, res.data]);
    } catch (err) {
      console.error(err);
    }
  };

  const handleNewClientClick = () => {
    openModal(
      <NewClientModal
        handleNewClientClick={(formData) => createNewClient(formData)}
      />,
    );
  };

  return (
    <div>
      <div>
        <div className='flex justify-between items-center'>
          <SubheadingText name>{project?.name}</SubheadingText>
          <DefaultButton handleButtonClick={() => handleNewClientClick()}>
            + New
          </DefaultButton>
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
