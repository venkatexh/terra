"use client";

const Modal = ({
  isOpen,
  children,
}: {
  isOpen: boolean;
  onClose: () => void;
  children: React.ReactNode;
}) => {
  return (
    isOpen && (
      <div className='fixed w-screen h-screen top-0 left-0 bg-black/60 bg-opacity-50 z-40'>
        <div
          className='absolute w-[50vw] max-h-[50vh] top-0 left-0 right-0 bottom-0 m-auto
          flex flex-col bg-zinc-900 border-[0.5px] border-gray-500 rounded-2xl z-50'>
          {children}
        </div>
      </div>
    )
  );
};

export default Modal;

export type ModalProps = {
  isOpen: boolean;
  onClose: () => void;
  children: React.ReactNode;
};
