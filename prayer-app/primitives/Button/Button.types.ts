import type {
  ButtonHTMLAttributes,
  ReactNode
} from "react";

export type ButtonVariant =
  | "primary"
  | "secondary"
  | "outline"
  | "ghost"
  | "danger"
  | "success";

export type ButtonSize =
  | "sm"
  | "md"
  | "lg"
  | "icon"
  | "iconSm"
  | "iconLg";

export interface ButtonProps
  extends ButtonHTMLAttributes<HTMLButtonElement> {

  variant?: ButtonVariant;

  size?: ButtonSize;

  fullWidth?: boolean;

  rounded?: "none" | "sm" | "md" | "lg" | "full";

  loading?: boolean;

  leftIcon?: ReactNode;

  rightIcon?: ReactNode;

  children: ReactNode;
}