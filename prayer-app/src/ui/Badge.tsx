type BadgeProps = {
  children: React.ReactNode;
  className?: string;
};

export default function Badge({ children, className = "" }: BadgeProps) {
  return <span className={`rounded-full bg-slate-100 px-3 py-1 text-sm text-slate-700 ${className}`}>{children}</span>;
}
