import { useState } from "react";
import { useModal } from "@/contexts/modal-context";

import { AiFillDelete } from "react-icons/ai";
import { FormInput } from "@/components/common/Inputs";
import { DefaultButton } from "@/components/common/Buttons";
import { DefaultText, LargeText } from "@/components/common/Texts";

import { ClientForm } from "@/types/console/projects/[id]/ClientForm";

export default function NewClientModal({
  handleNewClientClick,
}: {
  handleNewClientClick: (formData: ClientForm) => void;
}) {
  const { closeModal } = useModal();

  const [form, setForm] = useState<ClientForm>({ name: "", redirectUris: [] });
  const [redirectUrisInput, setRedirectUrisInput] = useState<number[]>([0]);

  return (
    <form
      className='h-full p-4 flex flex-col justify-between'
      onSubmit={(e) => {
        e.preventDefault();
        handleNewClientClick(form);
      }}>
      <div className='flex flex-col gap-2 overflow-scroll'>
        <LargeText className='sticky top-0 bg-zinc-900 pb-2'>
          Create a new client
        </LargeText>
        <FormInput
          type='text'
          name='name'
          label='Client name'
          value={form.name}
          placeholder='Client name'
          onChange={(e) => setForm({ ...form, name: e.target.value })}
        />
        {redirectUrisInput.map((uri: number, idx: number) => (
          <div key={uri} className='w-full flex justify-end gap-4 items-end'>
            <FormInput
              type='url'
              name='uri'
              className='w-full'
              label={`Client URI ${idx + 1}`}
              placeholder={`Client URI ${idx + 1}`}
              onChange={(e) =>
                setForm((prev) => ({
                  ...prev,
                  redirectUris: [
                    ...prev.redirectUris.slice(0, idx),
                    e.target.value,
                    ...prev.redirectUris.slice(idx + 1),
                  ],
                }))
              }
            />
            {idx != 0 && (
              <AiFillDelete
                className=' w-6 h-6 text-red-500 cursor-pointer'
                onClick={() => {
                  setForm((prev) => ({
                    ...prev,
                    redirectUris: [
                      ...prev.redirectUris.slice(0, idx),
                      ...prev.redirectUris.slice(idx + 1),
                    ],
                  }));
                  setRedirectUrisInput((prev) => [
                    ...prev.slice(0, idx),
                    ...prev.slice(idx + 1),
                  ]);
                }}
              />
            )}
          </div>
        ))}
        <DefaultButton
          type='button'
          transparent
          className='my-1 flex items-center justify-center'
          handleButtonClick={() =>
            setRedirectUrisInput([
              ...redirectUrisInput,
              redirectUrisInput.length,
            ])
          }>
          <span className='text-md mr-2'>+</span>Add URI
        </DefaultButton>
      </div>
      <div className='pt-4 flex justify-end items-center'>
        <DefaultButton type='submit'>Create</DefaultButton>
        <DefaultText
          className='text-gray-500 ml-4 cursor-pointer'
          onClick={() => closeModal()}>
          Close
        </DefaultText>
      </div>
    </form>
  );
}
