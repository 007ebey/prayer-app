type HeaderProps = {
  title: string;
  subtitle?: string;
};

export default function Header({ title, subtitle }: HeaderProps) {
  return (
    <header className="space-y-1">
      <h1 className="text-2xl font-semibold text-slate-900">{title}</h1>
      {subtitle ? <p className="text-sm text-slate-600">{subtitle}</p> : null}
    </header>
  );
}
