import { forwardRef } from "react";

import { cn } from "../../utils/cn";

import type { ButtonProps } from "./Button.types";

import { buttonVariants } from "./Button.styles";

const Button = forwardRef<
  HTMLButtonElement,
  ButtonProps
>(
(
  {
    children,
    variant,
    size,
    rounded,
    fullWidth,
    loading = false,
    leftIcon,
    rightIcon,
    disabled,
    className,
    ...props
  },
  ref
) => {

  return (

    <button

      ref={ref}

      disabled={disabled || loading}

      className={cn(
        buttonVariants({
          variant,
          size,
          rounded,
          fullWidth
        }),
        className
      )}

      {...props}

    >

      {loading && (

        <svg
          className="h-4 w-4 animate-spin"
          viewBox="0 0 24 24"
          fill="none"
        >
          <circle
            className="opacity-20"
            cx="12"
            cy="12"
            r="10"
            stroke="currentColor"
            strokeWidth="4"
          />

          <path
            className="opacity-90"
            fill="currentColor"
            d="M12 2a10 10 0 0 1 10 10h-4a6 6 0 0 0-6-6V2z"
          />
        </svg>

      )}

      {!loading && leftIcon}

      <span>

        {loading
          ? "Loading..."
          : children}

      </span>

      {!loading && rightIcon}

    </button>

  );

});

Button.displayName = "Button";

export default Button;