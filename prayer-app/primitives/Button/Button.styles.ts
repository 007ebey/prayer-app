import { cva } from "class-variance-authority";

export const buttonVariants = cva(
  [
    "inline-flex",
    "items-center",
    "justify-center",
    "font-medium",
    "transition-all",
    "duration-200",
    "select-none",
    "focus:outline-none",
    "focus:ring-2",
    "focus:ring-offset-2",
    "disabled:pointer-events-none",
    "disabled:opacity-50"
  ],
  {
    variants: {

      variant: {

        primary:
          "bg-primary-500 text-white hover:bg-primary-600 focus:ring-primary-300",

        secondary:
          "bg-secondary-500 text-white hover:bg-secondary-600 focus:ring-secondary-300",

        outline:
          "border border-slate-300 bg-white text-slate-800 hover:bg-slate-50",

        ghost:
          "text-slate-700 hover:bg-slate-100",

        danger:
          "bg-red-600 text-white hover:bg-red-700",

        success:
          "bg-green-600 text-white hover:bg-green-700"
      },

      size: {

        sm:
          "h-9 px-3 text-sm gap-2",

        md:
          "h-11 px-5 text-base gap-2",

        lg:
          "h-14 px-8 text-lg gap-3",

        icon: 
          "h-10 w-10 p-0",

        iconSm: 
          "h-8 w-8 p-0",

        iconLg: 
          "h-12 w-12 p-0"
      },

      rounded: {

        true: "rounded-full",

        false: "rounded-xl"
      },

      fullWidth: {

        true: "w-full",

        false: ""
      }

    },

    defaultVariants: {

      variant: "primary",

      size: "md",

      rounded: false,

      fullWidth: false

    }

  }
);