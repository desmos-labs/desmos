function findMenuEntryById(id, generatedMenu) {
  for (let item of generatedMenu) {
    if (item.type === "category") {
      // It's a category, search in the category sub items.
      const entry = findMenuEntryById(id, item.items);
      if (entry !== undefined) {
        // Entry found return it.
        return entry
      }
    } else if (item.type === "doc" && item.id === id) {
      // Entry found
      return item;
    }
  }

  // No entry found with the provided id, return undefined
  return undefined;
}

function convertPageWithHrefToExternLinks(docs, generatedMenu) {
  // Find the md pages that have the href field in the metadata.
  const toPatch = docs.filter(doc => doc.frontMatter.href !== undefined);
  for (let doc of toPatch) {
    const entry = findMenuEntryById(doc.id, generatedMenu);
    if (entry !== undefined) {
      // Convert the menu entry in an external reference
      entry.type = "link";
      entry.href = doc.frontMatter.href;
      // Remove the id field, is not allowed in the entries with type link.
      delete entry["id"];
    }
  }

  return generatedMenu;
}

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
  webpack: {
    jsLoader: (isServer) => ({
      loader: require.resolve('swc-loader'),
      options: {
        jsc: {
          parser: {
            syntax: 'typescript',
            tsx: true,
          },
          target: 'es2017',
        },
        module: {
          type: isServer ? 'commonjs' : 'es6',
        },
      },
    }),
  },
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
          async sidebarItemsGenerator({defaultSidebarItemsGenerator, ...args}) {
            const defaultItems = await defaultSidebarItemsGenerator(args);
            return convertPageWithHrefToExternLinks(args.docs, defaultItems);
          },
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
