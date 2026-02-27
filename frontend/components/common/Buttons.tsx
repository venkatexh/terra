type ButtonProps = {
  type?: "button" | "submit";
  className?: string;
  transparent?: boolean;
  children: React.ReactNode;
  handleButtonClick: () => void | null;
};

export function DefaultButton({
  type,
  children,
  className,
  transparent,
  handleButtonClick,
}: ButtonProps) {
  return (
    <button
      type={type}
      onClick={() => handleButtonClick()}
      className={`${className} ${
        !transparent
          ? "bg-orange-600 hover:bg-orange-700 text-white"
          : "border border-orange-600 text-orange-600"
      } 
      text-sm font-semibold py-1 px-4 rounded cursor-pointer`}>
      {children}
    </button>
  );
}

export function SecondaryButton({
  children,
  className,
  handleButtonClick,
}: ButtonProps) {
  return (
    <button
      onClick={() => handleButtonClick()}
      className={`${className} w-full bg-slate-200 hover:bg-slate-400 
      text-slate-900 font-bold py-2 px-4 rounded`}>
      {children}
    </button>
  );
}

export function SubmitButton({ children, className }: ButtonProps) {
  return (
    <button
      type='submit'
      className={`${className}w-full bg-blue-500 hover:bg-blue-700 
      text-white font-bold py-2 px-4 rounded`}>
      {children}
    </button>
  );
}
