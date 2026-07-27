// New file generated from FileUpload.tsx
import {
    InputHTMLAttributes,
    ReactNode,
} from "react";

export interface FileUploadProps
    extends Omit<
        InputHTMLAttributes<HTMLInputElement>,
        "type" | "value" | "onChange"
    > {

    label?: ReactNode;

    helperText?: ReactNode;

    error?: ReactNode;

    files?: File[];

    onFilesChange?(
        files: File[],
    ): void;

    accept?: string;

    multiple?: boolean;

    maxSize?: number;

    dragAndDrop?: boolean;

    disabled?: boolean;

    fullWidth?: boolean;

}