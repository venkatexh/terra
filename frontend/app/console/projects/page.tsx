"use client";

import useAxios from "@/hooks/useAxios";
import { useEffect, useState } from "react";
import { useModal } from "@/contexts/modal-context";

import { Project } from "@/types/console/projects/Project";

import { DefaultButton } from "@/components/common/Buttons";
import ProjectCard from "@/components/app/console/projects/ProjectCard";
import NewProjectModal from "@/components/app/console/projects/NewProjectModal";

export default function ProjectsPage() {
  const axios = useAxios();
  const { openModal, closeModal } = useModal();

  const [projects, setProjects] = useState<Project[]>([]);

  useEffect(() => {
    const fetchProjects = async () => {
      try {
        const res = await axios("/me/projects");
        setProjects(res.data);
      } catch (err) {
        console.error(err);
      }
    };

    fetchProjects();
  }, [axios]);

  async function createProject(formData: {
    projectName: string;
    projectDescription: string;
  }) {
    try {
      const res = await axios.post("/projects", {
        name: formData.projectName,
        description: formData.projectDescription,
      });
      closeModal();
      setProjects([...projects, res.data]);
    } catch (err) {
      console.error(err);
    }
  }

  const handleNewProjectClick = () => {
    openModal(<NewProjectModal handleCreateProjectClick={createProject} />);
  };

  return (
    <div className='flex flex-col gap-4'>
      <div className='flex justify-between items-center'>
        <div className='text-2xl'>All projects</div>
        <DefaultButton handleButtonClick={() => handleNewProjectClick()}>
          + New
        </DefaultButton>
      </div>
      <div className='py-6 flex flex-col gap-4'>
        {projects.map(
          (project: { name: string; description: string; id: string }) => (
            <ProjectCard
              key={project.id}
              id={project.id}
              name={project.name}
              description={project.description}
            />
          ),
        )}
      </div>
    </div>
  );
}
