import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "Moonlogs",
  description: "Business-event logging with ease",
  themeConfig: {
    logo: '/logo.png',
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'Get Started', link: '/tutorial/install' },
    ],

    sidebar: [
      {
        text: 'Get Started',
        collapsed: false,
        items: [
          { text: 'Installation', link: '/tutorial/install' },
          { text: 'Configuration', link: '/tutorial/configuration' },
          { text: 'Choosing DB Adapter', link: '/tutorial/choosing-db-adapter' }
        ]
      }
    ],

    search: {
      provider: 'algolia',
      options: {
        appId: 'YWZIJ9RIQS',
        apiKey: '9a92c576f2edfbb051ac7d2f933b1e48',
        indexName: 'moonlogs',
        insights: true
      }
    },

    footer: {
      message: 'Released under the Apache 2.0 License.',
      copyright: 'Copyright © 2024-present Mark Cholak',
    },

    socialLinks: [
      { icon: 'github', link: 'https://github.com/pijng/moonlogs' }
    ]

  },
  ignoreDeadLinks: 'localhostLinks'
})
