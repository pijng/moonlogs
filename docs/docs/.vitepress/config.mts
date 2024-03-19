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
          { text: 'Choosing DB adapter', link: '/tutorial/choosing-db-adapter' }
        ]
      }
    ],

    search: {
      provider: 'local',
      options: {}
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
