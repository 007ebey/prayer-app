// New file generated from FileUpload.tsx
import {
    forwardRef,
    useRef,
    useState,
} from "react";

import {
    Upload,
    X,
} from "lucide-react";

import { cn } from "../../utils/cn";

import {
    fileUploadVariants,
} from "./FileUpload.styles";

import type {
    FileUploadProps,
} from "./FileUpload.types";

const FileUpload = forwardRef<
    HTMLInputElement,
    FileUploadProps
>(({
    label,
    helperText,
    error,
    files = [],
    onFilesChange,
    accept,
    multiple = false,
    maxSize,
    dragAndDrop = true,
    disabled = false,
    fullWidth = true,
    className,
    ...props
}, ref) => {

    const inputRef =
        useRef<HTMLInputElement>(null);

    const [dragging, setDragging] =
        useState(false);

    const mergedRef = (
        node: HTMLInputElement | null,
    ) => {

        inputRef.current = node;

        if (
            typeof ref === "function"
        ) {

            ref(node);

        } else if (ref) {

            ref.current = node;

        }

    };

    const updateFiles = (
        list: FileList | null,
    ) => {

        if (!list) {
            return;
        }

        let next =
            Array.from(list);

        if (
            maxSize !== undefined
        ) {

            next = next.filter(
                file =>
                    file.size <= maxSize
            );

        }

        onFilesChange?.(
            multiple
                ? [
                      ...files,
                      ...next,
                  ]
                : next.slice(
                      0,
                      1
                  )
        );

    };

    return (

        <div
            className={cn(
                fullWidth &&
                    "w-full"
            )}
        >

            {label && (

                <label className="mb-2 block text-sm font-medium">

                    {label}

                </label>

            )}

            <div
                className={cn(
                    fileUploadVariants({
                        state:
                            disabled
                                ? "disabled"
                                : dragging
                                ? "dragging"
                                : error
                                ? "error"
                                : "default",
                    }),
                    className
                )}
                onClick={() => {

                    if (
                        !disabled
                    ) {

                        inputRef.current?.click();

                    }

                }}
                onDragOver={e => {

                    if (
                        !dragAndDrop
                    ) {

                        return;

                    }

                    e.preventDefault();

                    setDragging(
                        true
                    );

                }}
                onDragLeave={() =>
                    setDragging(
                        false
                    )
                }
                onDrop={e => {

                    if (
                        !dragAndDrop
                    ) {

                        return;

                    }

                    e.preventDefault();

                    setDragging(
                        false
                    );

                    updateFiles(
                        e.dataTransfer.files
                    );

                }}
            >

                <Upload
                    className="mx-auto mb-3"
                    size={32}
                />

                <p className="font-medium">

                    Drag & drop files here

                </p>

                <p className="text-sm text-muted-foreground">

                    or click to browse

                </p>

                <input
                    {...props}
                    ref={mergedRef}
                    hidden
                    type="file"
                    accept={accept}
                    multiple={
                        multiple
                    }
                    disabled={
                        disabled
                    }
                    onChange={e =>
                        updateFiles(
                            e.target
                                .files
                        )
                    }
                />

            </div>

            {files.length >
                0 && (

                <div className="mt-4 space-y-2">

                    {files.map(
                        file => (

                            <div
                                key={
                                    file.name
                                }
                                className="flex items-center justify-between rounded-lg border p-2"
                            >

                                <span className="truncate">

                                    {
                                        file.name
                                    }

                                </span>

                                <button
                                    type="button"
                                    onClick={() =>
                                        onFilesChange?.(
                                            files.filter(
                                                f =>
                                                    f !==
                                                    file
                                            )
                                        )
                                    }
                                >

                                    <X
                                        size={
                                            16
                                        }
                                    />

                                </button>

                            </div>

                        )
                    )}

                </div>

            )}

            {helperText &&
                !error && (

                    <p className="mt-2 text-sm text-muted-foreground">

                        {
                            helperText
                        }

                    </p>

                )}

            {error && (

                <p className="mt-2 text-sm text-destructive">

                    {error}

                </p>

            )}

        </div>

    );

});

FileUpload.displayName =
    "FileUpload";

export default FileUpload;