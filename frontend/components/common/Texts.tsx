type TextProps = {
  name?: boolean;
  className?: string;
  children: React.ReactNode;
  onClick?: () => void;
};

export function TinyText({ children, className, name, onClick }: TextProps) {
  return (
    <div
      className={`${className} ${name && "font-semibold"} text-xs`}
      onClick={onClick}>
      {children}
    </div>
  );
}

export function SmallText({ children, className, name, onClick }: TextProps) {
  return (
    <div
      className={`${className} ${name && "font-semibold"} text-sm`}
      onClick={onClick}>
      {children}
    </div>
  );
}

export function MediumText({ children, className, name, onClick }: TextProps) {
  return (
    <div
      className={`${className} ${name && "font-semibold"} text-lg`}
      onClick={onClick}>
      {children}
    </div>
  );
}

export function LargeText({ children, className, name, onClick }: TextProps) {
  return (
    <div
      className={`${className} ${name && "font-semibold"} text-xl`}
      onClick={onClick}>
      {children}
    </div>
  );
}

export function HeadingText({ children, className, name, onClick }: TextProps) {
  return (
    <div
      className={`${className} ${name && "font-semibold"} text-3xl`}
      onClick={onClick}>
      {children}
    </div>
  );
}

export function SubheadingText({
  children,
  className,
  name,
  onClick,
}: TextProps) {
  return (
    <div
      className={`${className} ${name && "font-semibold"} text-2xl`}
      onClick={onClick}>
      {children}
    </div>
  );
}

export function DefaultText({ children, className, name, onClick }: TextProps) {
  return (
    <div
      className={`${className} ${name && "font-semibold"} text-md`}
      onClick={onClick}>
      {children}
    </div>
  );
}
