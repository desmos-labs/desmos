module.exports = {
    title: "Desmos Docs",
    description: "Documentation for the Desmos blockchain.",
    head: [
        ['link', {rel: 'icon', href: '/assets/logo.png'}],
        ['link', {rel: "apple-touch-icon", sizes: "57x57", href: "/assets/pwa/apple-icon-57x57.png"}],
        ['link', {rel: "apple-touch-icon", sizes: "60x60", href: "/assets/pwa/apple-icon-60x60.png"}],
        ['link', {rel: "apple-touch-icon", sizes: "72x72", href: "/assets/pwa/apple-icon-72x72.png"}],
        ['link', {rel: "apple-touch-icon", sizes: "76x76", href: "/assets/pwa/apple-icon-76x76.png"}],
        ['link', {rel: "apple-touch-icon", sizes: "114x114", href: "/assets/pwa/apple-icon-114x114.png"}],
        ['link', {rel: "apple-touch-icon", sizes: "120x120", href: "/assets/pwa/apple-icon-120x120.png"}],
        ['link', {rel: "apple-touch-icon", sizes: "144x144", href: "/assets/pwa/apple-icon-144x144.png"}],
        ['link', {rel: "apple-touch-icon", sizes: "152x152", href: "/assets/pwa/apple-icon-152x152.png"}],
        ['link', {rel: "apple-touch-icon", sizes: "180x180", href: "/assets/pwa/apple-icon-180x180.png"}],
        ['link', {rel: "icon", type: "image/png", sizes: "192x192", href: "/assets/pwa/android-icon-192x192.png"}],
        ['link', {rel: "icon", type: "image/png", sizes: "32x32", href: "/assets/pwa/favicon-32x32.png"}],
        ['link', {rel: "icon", type: "image/png", sizes: "96x96", href: "/assets/pwa/favicon-96x96.png"}],
        ['link', {rel: "icon", type: "image/png", sizes: "16x16", href: "/assets/pwa/favicon-16x16.png"}],
        ['link', {rel: "manifest", href: "/assets/pwa/manifest.json",}],
        ['meta', {name: "msapplication-TileColor", content: "#ffffff"}],
        ['meta', {name: "msapplication-TileImage", content: "/ms-icon-144x144.png"}],
        ['meta', {name: "theme-color", content: "#ffffff"}],
        ['meta', {property: "og:title", content: "Desmos Documentation"}],
        ['meta', {property: "og:url", content: "https://docs.desmos.network/"}],
        ['meta', {property: "og:description", content: "Learn how to integrate with the Desmos blockchain"}],
        ['meta', {property: "og:image", content: "https://docs.desmos.network/assets/logo.png"}],
        ['meta', {roperty: "og:type", content: "website"}],
        ['meta', {property: "og:locale", content: "en_US"}],
    ],
    markdown: {
        lineNumbers: false,
    },
    plugins: [
        'tabs',
        '@vuepress/google-analytics',
        {
            'ga': 'UA-108489905-8'
        }
    ],
    themeConfig: {
        repo: "desmos-labs/desmos",
        editLinks: true,
        docsDir: "docs",
        docsBranch: "master",
        editLinkText: 'Edit this page on Github',
        lastUpdated: true,
        logo: "/assets/logo.png",
        nav: [
            {text: "Website", link: "https://desmos.network", target: "_blank"},
        ],
        sidebarDepth: 1,
        sidebar: [
            {
                title: "Types",
                collapsable: true,
                children: [
                    {
                        title: "Posts",
                        collapsable: true,
                        children: [
                            ["types/posts/post", "Post"],
                            ["types/posts/attachment", "Attachment"],
                            ["types/posts/post-poll-data", "PollData"],
                            ["types/posts/reaction", "Reactions"],
                        ]
                    },
                    {
                        title: "Profiles",
                        collapsable: true,
                        children: [
                            ["types/profiles/profile", "Profile"],
                            ["types/profiles/profile-pictures", "Profile pictures"],
                            ["types/profiles/dtag-transfer-request", "DTag transfer request"],
                        ]
                    },
                    {
                        title: "Relationships",
                        collapsable: true,
                        children: [
                            ["types/relationships/relationship", "Relationship"],
                            ["types/relationships/user-block", "User block"],
                        ]
                    },
                    {
                        title: "Reports",
                        collapsable: true,
                        children: [
                            ["types/reports/report", "Report"],
                        ]
                    },
                ]
            },
            {
                title: "Developers",
                collapsable: true,
                children: [
                    ["developers/overview", "Overview"],
                    ["developers/perform-transactions", "Performing transactions"],
                    ["developers/query-data", "Querying data"],
                    ["developers/observe-data", "Observing data"],
                    ["developers/developer-faq", "F.A.Q"],
                ]
            },
            {
                title: "Running a Fullnode",
                collapsable: true,
                children: [
                    ["fullnode/overview", "Overview"],
                    {
                        title: "Setup",
                        collapsable: false,
                        children: [
                            ["fullnode/setup/overview", "Overview"],
                            ["fullnode/setup/automatic", "Automatic setup"],
                            ["fullnode/setup/manual", "Manual setup"],
                        ]
                    },
                    ["fullnode/rocksdb-installation", "Using RocksDB"],
                    ["fullnode/update", "Update"],
                ]
            },
            {
                title: "Validators",
                collapsable: true,
                children: [
                    ["validators/overview", "Overview"],
                    ["validators/security", "Security"],
                    ["validators/setup", "Setup"],
                    {
                        title: "Upgrades",
                        collapsable: true,
                        children: [
                            ["validators/upgrades/manual", "Manual procedure"],
                            {
                                title: "Automatic procedure",
                                collapsable: true,
                                children: [
                                    ["validators/upgrades/automatic/overview", "Overview"],
                                    ["validators/upgrades/automatic/setup", "Setup"],
                                    ["validators/upgrades/automatic/workflow", "Workflow"],
                                ]
                            },
                        ]
                    },
                    ["validators/migrating", "Migrating"],
                    ["validators/halting", "Halting"],
                    ["validators/common-problems", "Common problems"],
                ]
            },
            {
                title: "Testnets",
                collapsable: true,
                children: [
                    ["testnets/overview", "Overview"],
                    ["testnets/create-local", "Create a local testnet"],
                    ["testnets/join-public", "Join the public testnet"],
                ],
            },
        ],
    }
};
