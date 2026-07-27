// semantics.ts

import { colors } from "./colors";

export const semantics = {
  brand: {
    primary: colors.primitive.burgundy[500],
    primaryHover: colors.primitive.burgundy[600],
    primaryActive: colors.primitive.burgundy[700],
    primaryForeground: colors.primitive.white,

    secondary: colors.primitive.gold[500],
    secondaryHover: colors.primitive.gold[600],
    secondaryActive: colors.primitive.gold[700],
    secondaryForeground: colors.primitive.gray[900],

    accent: colors.primitive.purple[500],
  },

  background: {
    page: colors.primitive.gray[50],
    surface: colors.primitive.white,
    subtle: colors.primitive.gray[100],
    elevated: colors.primitive.white,
    overlay: "rgba(0,0,0,0.45)",
    inverse: colors.primitive.gray[900],
  },

  surface: {
    sunken: colors.primitive.gray[100],
    default: colors.primitive.white,
    raised: colors.primitive.white,
    floating: colors.primitive.white,
  },

  text: {
    primary: colors.primitive.gray[900],
    secondary: colors.primitive.gray[500],
    muted: colors.primitive.gray[400],
    inverse: colors.primitive.white,
    disabled: colors.primitive.gray[300],
    link: colors.primitive.blue[600],
  },

  border: {
    subtle: colors.primitive.gray[100],
    default: colors.primitive.gray[200],
    strong: colors.primitive.gray[300],
    focus: colors.primitive.burgundy[500],
  },

  status: {
    success: colors.primitive.green[600],
    warning: colors.primitive.yellow[500],
    danger: colors.primitive.red[600],
    info: colors.primitive.blue[600],
  },

  session: {
    live: colors.primitive.green[600],
    scheduled: colors.primitive.blue[600],
    paused: colors.primitive.yellow[500],
    completed: colors.primitive.gray[500],
    cancelled: colors.primitive.red[600],
  },

  prayer: {
    active: colors.primitive.burgundy[500],
    answered: colors.primitive.green[600],
    waiting: colors.primitive.yellow[500],
    archived: colors.primitive.gray[400],
  },

  avatar: {
    background: colors.primitive.burgundy[100],
    foreground: colors.primitive.burgundy[700],
  },

  overlay: {
    light: "rgba(255,255,255,0.75)",
    dark: "rgba(0,0,0,0.45)",
  },

  focus: {
    ring: colors.primitive.burgundy[500],
    ringOffset: colors.primitive.white,
  },
} as const;

export default semantics;