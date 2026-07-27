type InputProps = {
  label?: string;
  placeholder?: string;
  type?: string;
  className?: string;
};

export default function Input({ label, placeholder, type = "text", className = "" }: InputProps) {
  return (
    <label className="block space-y-2 text-sm text-slate-200">
      {label ? <span>{label}</span> : null}
      <input
        type={type}
        placeholder={placeholder}
        className={`w-full rounded-lg border border-slate-700 bg-slate-950 px-3 py-2 text-white outline-none focus:border-emerald-500 ${className}`}
      />
    </label>
  );
}
