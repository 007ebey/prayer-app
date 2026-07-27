import { Monitor, Moon, Sun } from "lucide-react";

import Button  from "../../primitives/Button";

import { useTheme, type ThemeMode } from "./ThemeProvider";

const themes: {
    value: ThemeMode;
    label: string;
    icon: typeof Sun;
}[] = [
    {
        value: "light",
        label: "Light",
        icon: Sun,
    },
    {
        value: "dark",
        label: "Dark",
        icon: Moon,
    },
    {
        value: "system",
        label: "System",
        icon: Monitor,
    },
];

export default function ThemeSwitcher() {
    const {
        theme,
        setTheme,
    } = useTheme();

    return (

        <div
            className="
                inline-flex
                items-center
                rounded-xl
                border
                bg-surface
                p-1
                shadow-sm
            "
        >

            {themes.map((item) => {

                const Icon = item.icon;

                const active =
                    theme === item.value;

                return (

                    <Button
                        key={item.value}
                        variant={
                            active
                                ? "primary"
                                : "ghost"
                        }
                        size="sm"
                        onClick={() =>
                            setTheme(
                                item.value
                            )
                        }
                        leftIcon={
                            <Icon
                                size={16}
                            />
                        }
                    >

                        {item.label}

                    </Button>

                );

            })}

        </div>

    );

}