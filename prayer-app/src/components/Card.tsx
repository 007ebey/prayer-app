type CardProps = {
  title?: string;
  description?: string;
  children?: React.ReactNode;
  className?: string;
};

export default function Card({ title, description, children, className = "" }: CardProps) {
  return (
    <section className={`rounded-2xl border border-slate-200 bg-white p-5 shadow-sm ${className}`}>
      {title ? <h2 className="text-lg font-semibold text-slate-900">{title}</h2> : null}
      {description ? <p className="mt-2 text-sm text-slate-600">{description}</p> : null}
      {children}
    </section>
  );
}
