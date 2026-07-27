export const typography = {
    fontFamily: {
        heading: "Belanosima",
        body: "Lexend Deca",
        mono: "JetBrains Mono",
    },

    fontWeight: {
        light: 300,
        regular: 400,
        medium: 500,
        semibold: 600,
        bold: 700,
        extrabold: 800,
    },

    variants: {

        display: {
            fontFamily: "heading",
            fontSize: "4rem",
            fontWeight: 700,
            lineHeight: 1.1,
            letterSpacing: "-0.03em",
            textTransform: "none",
        },

        h1: {
            fontFamily: "heading",
            fontSize: "3rem",
            fontWeight: 700,
            lineHeight: 1.15,
            letterSpacing: "-0.02em",
            textTransform: "none",
        },

        h2: {
            fontFamily: "heading",
            fontSize: "2.25rem",
            fontWeight: 700,
            lineHeight: 1.2,
            letterSpacing: "-0.02em",
            textTransform: "none",
        },

        h3: {
            fontFamily: "heading",
            fontSize: "1.875rem",
            fontWeight: 600,
            lineHeight: 1.3,
            letterSpacing: "-0.01em",
            textTransform: "none",
        },

        h4: {
            fontFamily: "heading",
            fontSize: "1.5rem",
            fontWeight: 600,
            lineHeight: 1.35,
            letterSpacing: "0",
            textTransform: "none",
        },

        title: {
            fontFamily: "body",
            fontSize: "1.25rem",
            fontWeight: 600,
            lineHeight: 1.4,
            letterSpacing: "0",
            textTransform: "none",
        },

        subtitle: {
            fontFamily: "body",
            fontSize: "1.125rem",
            fontWeight: 500,
            lineHeight: 1.5,
            letterSpacing: "0",
            textTransform: "none",
        },

        body: {
            fontFamily: "body",
            fontSize: "1rem",
            fontWeight: 400,
            lineHeight: 1.7,
            letterSpacing: "0",
            textTransform: "none",
        },

        "body-sm": {
            fontFamily: "body",
            fontSize: ".875rem",
            fontWeight: 400,
            lineHeight: 1.6,
            letterSpacing: "0",
            textTransform: "none",
        },

        caption: {
            fontFamily: "body",
            fontSize: ".75rem",
            fontWeight: 400,
            lineHeight: 1.5,
            letterSpacing: ".01em",
            textTransform: "none",
        },

        overline: {
            fontFamily: "body",
            fontSize: ".6875rem",
            fontWeight: 600,
            lineHeight: 1.5,
            letterSpacing: ".08em",
            textTransform: "uppercase",
        },

        code: {
            fontFamily: "mono",
            fontSize: ".875rem",
            fontWeight: 400,
            lineHeight: 1.6,
            letterSpacing: "0",
            textTransform: "none",
        },

    },

} as const;

export default typography;