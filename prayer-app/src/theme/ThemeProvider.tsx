import {
  createContext,
  useCallback,
  useContext,
  useEffect,
  useMemo,
  useState,
  type PropsWithChildren,
} from "react";

export type ThemeMode =
  | "light"
  | "dark"
  | "system";

interface ThemeContextValue {
  theme: ThemeMode;

  resolvedTheme: "light" | "dark";

  setTheme: (theme: ThemeMode) => void;

  toggleTheme: () => void;
}

const ThemeContext =
  createContext<ThemeContextValue | null>(
    null
  );

const STORAGE_KEY = "theme";

function getSystemTheme() {
  if (
    typeof window === "undefined"
  ) {
    return "light";
  }

  return window.matchMedia(
    "(prefers-color-scheme: dark)"
  ).matches
    ? "dark"
    : "light";
}

export function ThemeProvider({
  children,
  defaultTheme = "system",
}: PropsWithChildren<{
  defaultTheme?: ThemeMode;
}>) {
  const [theme, setThemeState] =
    useState<ThemeMode>(defaultTheme);

  const [
    resolvedTheme,
    setResolvedTheme,
  ] = useState<"light" | "dark">(
    "light"
  );

  const applyTheme = useCallback(
    (mode: ThemeMode) => {
      const resolved =
        mode === "system"
          ? getSystemTheme()
          : mode;

      document.documentElement.dataset.theme =
        resolved;

      document.documentElement.classList.remove(
        "light",
        "dark"
      );

      document.documentElement.classList.add(
        resolved
      );

      setResolvedTheme(resolved);
    },
    []
  );

  useEffect(() => {
    const stored =
      localStorage.getItem(
        STORAGE_KEY
      ) as ThemeMode | null;

    if (stored) {
      setThemeState(stored);
    }
  }, []);

  useEffect(() => {
    applyTheme(theme);

    localStorage.setItem(
      STORAGE_KEY,
      theme
    );

    if (
      theme !== "system"
    ) {
      return;
    }

    const media =
      window.matchMedia(
        "(prefers-color-scheme: dark)"
      );

    const listener = () =>
      applyTheme("system");

    media.addEventListener(
      "change",
      listener
    );

    return () =>
      media.removeEventListener(
        "change",
        listener
      );
  }, [theme, applyTheme]);

  const setTheme = useCallback(
    (value: ThemeMode) => {
      setThemeState(value);
    },
    []
  );

  const toggleTheme =
    useCallback(() => {
      setThemeState((previous) =>
        previous === "dark"
          ? "light"
          : "dark"
      );
    }, []);

  const value = useMemo(
    () => ({
      theme,
      resolvedTheme,
      setTheme,
      toggleTheme,
    }),
    [
      theme,
      resolvedTheme,
      setTheme,
      toggleTheme,
    ]
  );

  return (
    <ThemeContext.Provider
      value={value}
    >
      {children}
    </ThemeContext.Provider>
  );
}

export function useTheme() {
  const context =
    useContext(ThemeContext);

  if (!context) {
    throw new Error(
      "useTheme must be used within ThemeProvider."
    );
  }

  return context;
}