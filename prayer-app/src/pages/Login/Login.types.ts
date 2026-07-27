// New file generated from Login.tsx
import {
    ReactNode,
} from "react";

export interface LoginProps {

    logo?: ReactNode;

    title?: ReactNode;

    subtitle?: ReactNode;

    hero?: ReactNode;

    footer?: ReactNode;

    onGoogleLogin?(): void;

    onInstagramLogin?(): void;

    loading?: boolean;

}