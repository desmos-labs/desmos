module.exports = {
    title: "Desmos",
    description: "Documentation for the Desmos blockchain.",
    // ga: "UA-108489905-8",
    head: [
        ['link', {rel: 'icon', href: '/favicon.png'}]
    ],
    markdown: {
        lineNumbers: true,
    },
    plugins: [
        'vuepress-plugin-element-tabs',
        '@vuepress/google-analytics',
        {
            'ga': 'UA-108489905-8' // UA-00000000-0
        }
    ],
    themeConfig: {
        repo: "desmos-labs/desmos",
        editLinks: true,
        docsDir: "docs",
        docsBranch: "master",
        editLinkText: 'Edit this page on Github',
        lastUpdated: true,
        nav: [
            {text: "Forbole", link: "https://forbole.com"},
        ],
        sidebarDepth: 2,
        sidebar: [
            {
                title: "Developers",
                collapsable: false,
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
                collapsable: false,
                children: [
                    ["fullnode/overview", "Overview"],
                    ["fullnode/installation", "Install and Running a Fullnode"],
                ]
            },
            {
                title: "Validators",
                collapsable: false,
                children: [
                    ["validators/overview", "Overview"],
                    ["validators/security", "Security"],
                    ["validators/validator-setup", "Setup"],
                ]
            },
        ],
    }
};
