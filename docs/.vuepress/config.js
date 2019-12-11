module.exports = {
    title: "Desmos",
    description: "Documentation for the Desmos blockchain.",
    // ga: "UA-51029217-2",
    head: [
        ['link', {rel: 'icon', href: '/icon.png'}]
    ],
    markdown: {
        lineNumbers: true,
    },
    plugins: [
        'latex'
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
                title: "Validators",
                collapsable: false,
                children: [
                    ["validators/overview", "Overview"],
                    ["validators/security", "Security"],
                    ["validators/validator-setup", "Setup"],
                    ["validators/validator-faq", "F.A.Q."],
                ]
            },
        ],
    }
};
