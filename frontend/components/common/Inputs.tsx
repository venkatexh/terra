import { SmallText } from "./Texts";

export const FormInput = ({
  type,
  name,
  value,
  label,
  onChange,
  className,
  placeholder,
}: Input) => {
  return (
    <div className='w-full flex flex-col gap-1'>
      {label && <SmallText>{label}</SmallText>}
      <input
        type={type}
        name={name}
        value={value}
        onChange={onChange}
        className={`${className} 
        px-2 py-1 border border-gray-500 rounded focus:outline-none`}
        placeholder={placeholder}
      />
    </div>
  );
};

export const FormTextArea = ({
  name,
  value,
  label,
  resize,
  onChange,
  className,
  placeholder,
}: Input) => {
  return (
    <div className='w-full flex flex-col gap-1'>
      {label && <SmallText>{label}</SmallText>}
      <textarea
        name={name}
        value={value}
        onChange={onChange}
        className={`${className} ${!resize && "resize-none"} 
        px-2 py-1 border border-gray-500 rounded focus:outline-none`}
        placeholder={placeholder}
      />
    </div>
  );
};

export const SearchInput = ({ placeholder, className, onChange }: Input) => {
  return (
    <input
      type='search'
      onChange={onChange}
      className={className}
      placeholder={placeholder}
    />
  );
};

type Input = {
  type?: string;
  name: string;
  value?: string;
  label: string;
  resize?: boolean;
  className?: string;
  placeholder: string;
  onChange: (
    event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
  ) => void;
};
