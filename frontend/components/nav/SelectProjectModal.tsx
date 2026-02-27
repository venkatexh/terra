import useAxios from "@/hooks/useAxios";
import { useEffect, useState } from "react";

import { DefaultButton } from "../common/Buttons";
import { useModal } from "@/contexts/modal-context";
import { DefaultText, SmallText } from "../common/Texts";

export default function SelectProjectModal() {
  const axios = useAxios();
  const { closeModal } = useModal();

  const [projects, setProjects] = useState([]);
  const [searchInput, setSearchInput] = useState("");
  const [filteredProjects, setFilteredProjects] = useState([]);

  useEffect(() => {
    const fetchProjects = async () => {
      try {
        const res = await axios("/me/projects");
        setProjects(res.data);
        setFilteredProjects(res.data);
      } catch (err) {
        console.error(err);
      }
    };

    fetchProjects();
  }, [axios]);

  const handleSearch = (search: string) => {
    setSearchInput(search);
    const filtered = projects.filter(
      (project: Project) =>
        project?.name.toLowerCase().includes(search.toLowerCase()) ||
        project?.id.toLowerCase().includes(search.toLowerCase()) ||
        project?.description.toLowerCase().includes(search.toLowerCase()),
    );
    setFilteredProjects(filtered);
  };

  return (
    <div className='w-full h-full flex flex-col gap-4'>
      <div className='px-4 mt-4'>
        <DefaultText>Select current project</DefaultText>
      </div>
      <div className='px-4'>
        <input
          type='search'
          value={searchInput}
          placeholder='Search projects by name, ID or description'
          className='w-full px-2 py-1 border border-gray-500 focus:outline-none rounded-xl'
          onChange={(e) => handleSearch(e.target.value)}
        />
      </div>
      <div className='overflow-scroll'>
        <div className='sticky top-0 flex px-4 py-1 bg-zinc-900 border-b border-gray-500'>
          <SmallText className='w-1/5'>Name</SmallText>
          <SmallText className='w-2/5'>Project ID</SmallText>
          <SmallText className='w-2/5'>Description</SmallText>
        </div>
        <div>
          {filteredProjects.map((project: Project, idx) => (
            <div
              key={project.id}
              className='flex px-4 py-1 border-b border-gray-500 hover:bg-gray-800  cursor-pointer'>
              <SmallText
                className={`${idx === 0 && "font-semibold text-orange-500"} w-1/5`}>
                {project.name}
              </SmallText>
              <SmallText className='w-2/5 line-clamp-1'>{project.id}</SmallText>
              <SmallText className='w-2/5 line-clamp-1'>
                {project.description}
              </SmallText>
            </div>
          ))}
        </div>
      </div>
      <div className='flex items-center gap-4 mt-auto ml-auto px-4 pb-4 text-white/50'>
        <DefaultButton handleButtonClick={() => {}} transparent>
          New Project
        </DefaultButton>
        <div className='cursor-pointer' onClick={closeModal}>
          Close
        </div>
      </div>
    </div>
  );
}

type Project = {
  id: string;
  name: string;
  description: string;
};
