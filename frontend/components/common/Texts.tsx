type TextProps = {
  name?: boolean;
  className?: string;
  children: React.ReactNode;
};

export function TinyText({ children, className, name }: TextProps) {
  return (
    <div className={`${className} ${name && "font-semibold"} text-xs`}>
      {children}
    </div>
  );
}

export function SmallText({ children, className, name }: TextProps) {
  return (
    <div className={`${className} ${name && "font-semibold"} text-sm`}>
      {children}
    </div>
  );
}

export function MediumText({ children, className, name }: TextProps) {
  return (
    <div className={`${className} ${name && "font-semibold"} text-lg`}>
      {children}
    </div>
  );
}

export function LargeText({ children, className, name }: TextProps) {
  return (
    <div className={`${className} ${name && "font-semibold"} text-xl`}>
      {children}
    </div>
  );
}

export function HeadingText({ children, className, name }: TextProps) {
  return (
    <div className={`${className} ${name && "font-semibold"} text-3xl`}>
      {children}
    </div>
  );
}

export function SubheadingText({ children, className, name }: TextProps) {
  return (
    <div className={`${className} ${name && "font-semibold"} text-2xl`}>
      {children}
    </div>
  );
}

export function DefaultText({ children, className, name }: TextProps) {
  return (
    <div className={`${className} ${name && "font-semibold"} text-md`}>
      {children}
    </div>
  );
}
