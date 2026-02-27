import { useState } from "react";
import { useModal } from "@/contexts/modal-context";

import { DefaultButton } from "@/components/common/Buttons";
import { DefaultText, LargeText } from "@/components/common/Texts";
import { FormInput, FormTextArea } from "@/components/common/Inputs";

export default function NewProjectModal({
  handleCreateProjectClick,
}: {
  handleCreateProjectClick: (formData: {
    projectName: string;
    projectDescription: string;
  }) => void;
}) {
  const { closeModal } = useModal();
  const [form, setForm] = useState({
    projectName: "",
    projectDescription: "",
  });

  function handleFormChange(
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
  ) {
    setForm({
      ...form,
      [e.target.name]: e.target.value,
    });
  }

  return (
    <form
      className='h-full p-4 flex flex-col gap-4 justify-between'
      onSubmit={(e) => {
        e.preventDefault();
        handleCreateProjectClick(form);
      }}>
      <div className='flex flex-col gap-2'>
        <LargeText>Create a new project</LargeText>
        <FormInput
          type='text'
          name='projectName'
          value={form.projectName}
          placeholder='Project name'
          className=''
          onChange={(e) => handleFormChange(e)}
          label='Project name'
        />
        <FormTextArea
          name='projectDescription'
          value={form.projectDescription}
          placeholder='Project description'
          onChange={(e) => handleFormChange(e)}
          label='Description'
          resize={false}
        />
      </div>
      <div className='flex justify-end items-center'>
        <DefaultButton
          type='submit'
          handleButtonClick={() => handleCreateProjectClick(form)}>
          Create
        </DefaultButton>
        <DefaultText
          className='ml-4 text-gray-500 cursor-pointer'
          onClick={() => closeModal()}>
          Close
        </DefaultText>
      </div>
    </form>
  );
}
