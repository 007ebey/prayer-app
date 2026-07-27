import { colors } from "./colors";

export const borders = {
  width: {
    none: 0,
    hairline: 1,
    thin: 1,
    medium: 2,
    thick: 3,
    heavy: 4,
  },

  style: {
    solid: "solid",
    dashed: "dashed",
    dotted: "dotted",
    double: "double",
    none: "none",
  },

  color: {
    transparent: colors.primitive.transparent,

    subtle: colors.primitive.gray[100],
    default: colors.primitive.gray[200],
    muted: colors.primitive.gray[300],
    strong: colors.primitive.gray[500],

    primary: colors.primitive.burgundy[500],
    secondary: colors.primitive.gold[500],

    success: colors.primitive.green[600],
    warning: colors.primitive.yellow[500],
    danger: colors.primitive.red[600],
    info: colors.primitive.blue[600],
  },

  focus: {
    width: 2,
    offset: 2,
    color: colors.primitive.burgundy[500],
  },
} as const;

export default borders;