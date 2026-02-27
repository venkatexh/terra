"use client";

import { useModal } from "@/contexts/modal-context";
import Link from "next/link";
import SelectProjectModal from "./SelectProjectModal";

export default function PrimaryNavbar() {
  const { openModal } = useModal();
  const handleProjectSelectorClick = () => {
    openModal(<SelectProjectModal />);
  };

  return (
    <div className='w-screen h-12 fixed top-0 left-0 border-b border-gray-500'>
      <div className='w-1/2 h-full flex justify-between items-center mx-auto'>
        <Link href='/'>
          <span className='font-semibold text-xl'>Terra</span>
        </Link>
        <div className='flex gap-2 items-center justify-center'>
          <div
            className='h-8 px-4 border border-gray-500 rounded flex justify-center 
            items-center cursor-pointer font-semibold'
            onClick={() => handleProjectSelectorClick()}>
            default
          </div>
          <div>settings</div>
          <div>profile</div>
        </div>
      </div>
    </div>
  );
}
