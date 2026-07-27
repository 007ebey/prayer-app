import type { ReactNode } from "react";

type ModalProps = {
  title?: string;
  children: ReactNode;
  open?: boolean;
};

export default function Modal({ title, children, open = true }: ModalProps) {
  if (!open) return null;

  return (
    <div className="fixed inset-0 flex items-center justify-center bg-slate-950/60 p-4">
      <div className="w-full max-w-md rounded-2xl bg-white p-6 shadow-xl">
        {title ? <h3 className="text-lg font-semibold text-slate-900">{title}</h3> : null}
        <div className="mt-4">{children}</div>
      </div>
    </div>
  );
}
