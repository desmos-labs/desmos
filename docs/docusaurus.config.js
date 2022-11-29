module.exports = {
  title: 'Desmos documentation',
  staticDirectories: ['static'],
  tagline: 'Desmos network official documentation for developers and validators',
  url: 'https://test-docs.desmos.network',
  baseUrl: '/',
  onBrokenLinks: 'warn',
  onBrokenMarkdownLinks: 'warn',
  onDuplicateRoutes: 'warn',
  favicon: 'assets/favicon.ico',
  organizationName: 'desmos-labs', // Usually your GitHub org/user name.
  projectName: 'desmos', // Usually your repo name.
  themeConfig: {
    colorMode: {
      defaultMode: 'dark',
      respectPrefersColorScheme: true,
    },
    algolia: {
      apiKey: '492b6729d095b18f5599d6584e00ae11',
      appId: '1IAGPKAXGP',
      indexName: 'desmos',
      contextualSearch: false,
    },
    docs: {
      sidebar: {
        hideable: true,
      }
    },
    navbar: {
      logo: {
        alt: 'Desmos logo',
        src: 'assets/logo.svg',
        srcDark: 'assets/logo.svg',
        href: 'https://docs.desmos.network'
      },
      items: [
        {
          type: 'doc',
          docId: 'intro', // open page of section
          position: 'left',
          label: 'Documentation',
        },
        // {to: '/blog', label: 'Blog', position: 'left'}, to add extra sections
        {
          type: 'docsVersionDropdown',
          position: 'right',
          dropdownActiveClassDisabled: true,
        },
        /*{
          // Re-add this if we want to use localization
          type: 'localeDropdown',
          position: 'right',
        },*/
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Related docs',
          items: [
            {
              label: 'Cosmos SDK',
              href: 'https://docs.cosmos.network',
            },
            {
              label: 'CosmWasm',
              href: 'https://docs.cosmwasm.com/en/docs/1.0/'
            }
          ],
        },
        {
          title: 'Community',
          items: [
            {
              label: 'Twitter',
              href: 'https://twitter.com/DesmosNetwork',
            },
            {
              label: 'Discord',
              href: 'https://discord.desmos.network/',
            },
            {
              label: 'Medium',
              href: 'https://medium.com/desmosnetwork'
            },
            {
              label: 'Telegram',
              href: 'https://t.me/desmosnetwork',
            },
            {
              label: 'Reddit (not-official)',
              href: 'https://www.reddit.com/r/DesmosNetwork/'
            }
          ],
        },
        {
          title: 'More',
          items: [
            {
              label: 'Website',
              to: 'https://www.desmos.network',
            },
            {
              label: 'GitHub',
              href: 'https://github.com/desmos-labs/desmos',
            },
          ],
        },
      ],
      logo: {
        alt: 'Desmos Logo',
        src: 'assets/logo.png',
        href: 'https://www.desmos.network',
      },
      copyright: `Copyright © ${new Date().getFullYear()} Desmos Network`,
    },
  },
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        docs: {
          routeBasePath: '/',
          sidebarPath: require.resolve('./sidebars.js'),
          sidebarCollapsible: true,
          editUrl: 'https://github.com/desmos-labs/desmos/tree/master/docs',
          showLastUpdateTime: true,
          lastVersion: "current",
          exclude: [
            './architecture/adr-template.md'
          ],
          versions: {
            current: {
              label: "master"
            },
          }
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      },
    ],
  ],
  themes: [
    '@you54f/theme-github-codeblock'
  ],
  plugins: [
    [
      "@edno/docusaurus2-graphql-doc-generator",
      {
        schema: "docs/07-graphql/schema.graphql",
        root: "docs/",
        baseURL: "07-graphql",
        homepage: "docs/07-graphql/01-overview.md",
        pretty: true,
      }
    ],
  ]
  /*i18n: { // add for localization
    defaultLocale: 'en',
    locales: ['en', 'chinese'],
  },*/
};
