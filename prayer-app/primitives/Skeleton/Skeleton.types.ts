export interface SkeletonProps
    extends React.HTMLAttributes<HTMLDivElement> {
    variant?: "text" | "rect" | "rounded" | "circle";
    animation?: "pulse" | "wave" | "none";
    width?: number | string;
    height?: number | string;
    size?: number | string;
}