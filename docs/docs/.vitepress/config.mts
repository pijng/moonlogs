import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "Moonlogs",
  description: "Business-event logging tool with a built-in user-friendly web interface for easy access to events",
  themeConfig: {
    logo: '/logo.svg',
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'Get Started', link: '/tutorial/install' },
    ],

    sidebar: [
      {
        text: 'Get Started',
        collapsed: false,
        items: [
          { text: 'What is Moonlogs', link: '/tutorial/what-is-moonlogs' },
          { text: 'Installation', link: '/tutorial/install' },
          { text: 'Configuration', link: '/tutorial/configuration' },
          { text: 'Choosing DB Adapter', link: '/tutorial/choosing-db-adapter' }
        ]
      },
      {
        text: 'Web UI',
        collapsed: false,
        items: [
          { text: 'Introduction', link: '/web-ui/introduction' },
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
