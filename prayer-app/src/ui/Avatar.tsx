type AvatarProps = {
  name?: string;
  className?: string;
};

export default function Avatar({ name = "U", className = "" }: AvatarProps) {
  return (
    <div className={`flex h-10 w-10 items-center justify-center rounded-full bg-emerald-500 font-semibold text-white ${className}`}>
      {name.charAt(0).toUpperCase()}
    </div>
  );
}
