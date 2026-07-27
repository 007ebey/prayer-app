type ButtonProps = {
  children: React.ReactNode;
  className?: string;
};

export default function Button({ children, className = "" }: ButtonProps) {
  return (
    <button className={`rounded-lg bg-emerald-500 px-4 py-2 font-medium text-white hover:bg-emerald-600 ${className}`}>
      {children}
    </button>
  );
}
