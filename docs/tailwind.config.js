module.exports = {
  corePlugins: {
    preflight: false, // disable Tailwind's reset
  },
  content: ["./src/**/*.{js,jsx,ts,tsx}"],
  darkMode: ["class", '[data-theme="dark"]'],
  blocklist: ["container"],
  theme: {
    screens: {
      md: "768px",
      lg: "1280px",
      xl: "1920px",
    },
    extend: {
      width: {
        mobile: "375px",
        md: "768px",
        lg: "1280px",
        xl: "1920px",
      },
      padding: {
        xMobile: "20px",
        xMd: "40px",
        xLg: "90px",
        xXl: "100px",
        yMobile: "26px",
        yMd: "60px",
        yLg: "80px",
        yXl: "80px",
        "navbar-mobile": "60px",
        "navbar-md": "60px",
      },
      margin: {
        xMobile: "20px",
        xMd: "40px",
        xLg: "90px",
        xXl: "100px",
        yMobile: "26px",
        yMd: "60px",
        yLg: "80px",
        yXl: "80px",
        "navbar-mobile": "60px",
        "navbar-md": "60px",
      },
    },
  },
  plugins: [],
};
