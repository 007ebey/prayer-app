import {
    Globe,
    Camera,
    User,
    ShieldCheck,
} from "lucide-react";

import {
    Button,
    Card,
    Stack,
    Typography,
} from "../../../primitives";

import {
    cn,
} from "../../../utils/cn";

import {
    loginCardVariants,
    loginVariants,
} from "./Login.styles";

import type {
    LoginProps,
} from "./Login.types";

const Login = ({
    logo,
    title = "Welcome",
    subtitle = "Sign in to continue.",
    hero,
    footer,

    role = "participant",
    onRoleChange,

    onGoogleLogin,
    onInstagramLogin,

    loading = false,
}: LoginProps) => {

    return (

        <main
            className={cn(
                loginVariants()
            )}
        >

            <Card
                className={loginCardVariants()}
            >

                <Stack
                    align="center"
                    gap="lg"
                >

                    {logo}

                    <Typography
                        variant="h2"
                        align="center"
                    >
                        {title}
                    </Typography>

                    <Typography
                        variant="body"
                        color="muted"
                        align="center"
                    >
                        {subtitle}
                    </Typography>

                    {hero}

                </Stack>


                {/* Role */}

                <Stack gap="sm">

                    <Typography
                        variant="caption"
                        color="muted"
                    >
                        Continue as
                    </Typography>

                    <div className="grid grid-cols-2 gap-2">

                        <Button
                            variant={
                                role === "participant"
                                    ? "primary"
                                    : "outline"
                            }
                            leftIcon={
                                <User size={18} />
                            }
                            onClick={() =>
                                onRoleChange?.(
                                    "participant"
                                )
                            }
                        >
                            Participant
                        </Button>

                        <Button
                            variant={
                                role === "admin"
                                    ? "primary"
                                    : "outline"
                            }
                            leftIcon={
                                <ShieldCheck size={18} />
                            }
                            onClick={() =>
                                onRoleChange?.(
                                    "admin"
                                )
                            }
                        >
                            Admin
                        </Button>

                    </div>

                </Stack>


                {/* Provider */}

                <Stack gap="md">

                    <Button
                        fullWidth
                        loading={loading}
                        leftIcon={
                            <Globe size={18} />
                        }
                        onClick={() =>
                            onGoogleLogin?.(
                                role
                            )
                        }
                    >
                        Continue with Google
                    </Button>

                    <Button
                        fullWidth
                        variant="outline"
                        loading={loading}
                        leftIcon={
                            <Camera size={18} />
                        }
                        onClick={() =>
                            onInstagramLogin?.(
                                role
                            )
                        }
                    >
                        Continue with Instagram
                    </Button>

                </Stack>


                {footer ?? (

                    <Typography
                        variant="caption"
                        align="center"
                        color="muted"
                    >
                        By continuing you agree to the
                        Terms of Service and Privacy Policy.
                    </Typography>

                )}

            </Card>

        </main>

    );

};

export default Login;