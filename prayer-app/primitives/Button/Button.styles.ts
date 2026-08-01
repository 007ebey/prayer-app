import { cva } from "class-variance-authority";

export const buttonVariants = cva(
  [
    "inline-flex items-center justify-center gap-2",
    "font-medium",
    "transition-colors",
    "focus-visible:outline-none",
    "focus-visible:ring-2",
    "focus-visible:ring-primary",
    "focus-visible:ring-offset-2",
    "disabled:pointer-events-none",
    "disabled:opacity-50",
  ],
  {
    variants: {
      variant: {
        primary: [
          "bg-primary",
          "text-primary-foreground",
          "hover:opacity-90",
        ],

        secondary: [
          "bg-secondary",
          "text-secondary-foreground",
          "hover:opacity-90",
        ],

        outline: [
          "border",
          "border-border",
          "bg-transparent",
          "text-foreground",
          "hover:bg-muted",
        ],

        ghost: [
          "bg-transparent",
          "text-foreground",
          "hover:bg-muted",
        ],

        danger: [
          "bg-destructive",
          "text-destructive-foreground",
          "hover:opacity-90",
        ],

        success: [
          "bg-success",
          "text-success-foreground",
          "hover:opacity-90",
        ],
      },

      size: {
        sm: "h-8 px-3 text-sm",
        md: "h-10 px-4 text-sm",
        lg: "h-12 px-6 text-base",

        iconSm: "h-8 w-8 p-0",
        icon: "h-10 w-10 p-0",
        iconLg: "h-12 w-12 p-0",
      },

      rounded: {
        none: "rounded-none",
        sm: "rounded-sm",
        md: "rounded-md",
        lg: "rounded-lg",
        full: "rounded-full",
      },

      fullWidth: {
        true: "w-full",
        false: "",
      },
    },

    defaultVariants: {
      variant: "primary",
      size: "md",
      rounded: "md",
      fullWidth: false,
    },
  }
);