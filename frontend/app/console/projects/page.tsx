"use client";

import ProjectCard from "@/components/console/projects/ProjectCard";
import useAxios from "@/hooks/useAxios";
import { useEffect, useState } from "react";

export default function ProjectsPage() {
  const axios = useAxios();
  const [projects, setProjects] = useState([]);

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
  return (
    <div className='flex flex-col gap-4'>
      {projects.map((project: { name: string; id: string }) => (
        <ProjectCard key={project.id} id={project.id} name={project.name} />
      ))}
    </div>
  );
}
