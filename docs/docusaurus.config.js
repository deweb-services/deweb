// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion

const lightCodeTheme = require("prism-react-renderer/themes/github");
const darkCodeTheme = require("prism-react-renderer/themes/dracula");

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: "Decentralized Web Services (DWS)",
  tagline: "Web3.0 starts here",
  url: "https://deweb-services.github.com ",
  baseUrl: "/deweb/",
  onBrokenLinks: "throw",
  onBrokenMarkdownLinks: "warn",
  favicon: "img/favicon.ico",
  organizationName: "deweb-services", // Usually your GitHub org/user name.
  projectName: "deweb", // Usually your repo name.

  presets: [
    [
      "classic",
      /** @type {import('@docusaurus/preset-classic').Options} */
      ({
        docs: {
          routeBasePath: "/",
        },
        blog: false,
        theme: {
          customCss: require.resolve("./src/css/custom.css"),
        },
      }),
    ],
  ],

  themeConfig:
    /** @type {import('@docusaurus/preset-classic').ThemeConfig} */
    ({
      navbar: {
        title: "Decentralized Web Services (DWS)",
        logo: {
          alt: "Decentralized Web Services (DWS)",
          src: "img/dws.png",
        },
        items: [
          {
            href: "https://github.com/deweb-services/deweb",
            label: "GitHub",
            position: "right",
          },
        ],
      },
      footer: {
        style: "dark",
        links: [
          {
            title: "Docs",
            items: [
              {
                label: "Introduction",
                to: "/introduction",
              },
              {
                label: "Validator Setup Guide",
                to: "/guides/validator-setup-guide",
              },
            ],
          },
          {
            title: "Community",
            items: [
              {
                label: "Discord",
                href: "https://discord.gg/NENJeq58Tc",
              },
              {
                label: "Twitter",
                href: "https://twitter.com/dewebservices",
              },
            ],
          },
          {
            title: "More",
            items: [
              {
                label: "Medium",
                href: "https://blog.deweb.services/",
              },
              {
                label: "Mastodon",
                href: "https://mastodon.social/@dewebservices",
              },
              {
                label: "GitHub",
                href: "https://github.com/deweb-services",
              },
              {
                label: "Explorer",
                href: "https://explore.deweb.services/",
              },
            ],
          },
        ],
        copyright: `Copyright © ${new Date().getFullYear()} Decentralized Web Services.`,
      },
      prism: {
        theme: lightCodeTheme,
        darkTheme: darkCodeTheme,
      },
    }),
};

module.exports = config;
