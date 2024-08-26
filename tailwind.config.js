/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./app/**/*.templ",
    "./app/**/*.html",
    "./app/**/*.go",
    "./app/**/*.ts",
  ],
  theme: {
    extend: {
      fontFamily: {
        sans: ["Roboto"],
      },
      fontWeight: {
        hairline: "100",
        thin: "200",
        light: "300",
        normal: "400",
        medium: "500",
        semibold: "600",
        bold: "700",
        extrabold: "800",
        black: "900",
      },
      screens: {
        "3xl": "2048px",
      },
      typography: (theme) => ({
        DEFAULT: {
          css: {
            h2: {
              fontWeight: theme("fontWeight.light"),
              letterSpacing: theme("letterSpacing.tight"),
              color: theme("colors.stone.600"),
            },
            a: {
              color: theme("colors.red.600"),
              "&:hover": {
                color: theme("colors.red.500"),
              },
            },
          },
        },
      }),
    },
  },
  plugins: [require("@tailwindcss/typography")],
  variants: {
    extend: {
      translate: ["responsive", "hover", "focus", "active", "group-hover"],
      opacity: ["responsive", "hover", "focus", "active", "group-hover"],
    },
  },
  safelist: [
    "translate-x-0",
    "-translate-x-full",
    "opacity-0",
    "opacity-100",
    "translate-y-0",
    "translate-y-1",
    "rotate-180",
    "pointer-events-none",
    "font-semibold",
    "shadow-sm",
    "font-hairline",
    "font-thin",
    "font-light",
    "font-normal",
    "font-medium",
    "font-semibold",
    "font-bold",
    "font-extrabold",
    "font-black",
    {
      pattern:
        /^(bg|text|border)-(.+)-(50|100|200|300|400|500|600|700|800|900)$/,
      variants: ["hover", "focus", "active"],
    },
    {
      pattern: /^(rounded|sharp|soft)-(none|sm|md|lg|xl|2xl|3xl|full)$/,
    },
    {
      pattern: /^(m|p|px|py|mt|mb|ml|mr)-[0-9]+$/,
    },
    {
      pattern: /^text-(xs|sm|base|lg|xl|2xl|3xl|4xl|5xl|6xl)$/,
    },
    {
      pattern: /^(outline|ring|border)(-offset)?(-[0-9]+)?$/,
      variants: ["focus-visible"],
    },
  ],
};
