const config = require('./contrib/config.js')
const fs = require('fs')
const admonitions = require('remark-admonitions')

const githubRepoName =
  config.projectSlug === 'ecosystem' ? 'docs' : config.projectSlug

const links = [
  {
    to: '/',
    activeBasePath: `${config.projectSlug}/docs`,
    label: `Docs`,
    position: 'left'
  },
  {
    href: 'https://www.ory.sh/docs',
    label: 'Ecosystem',
    position: 'left'
  },
  {
    href: 'https://www.ory.sh/blog',
    label: 'Blog',
    position: 'left'
  },
  {
    href: 'https://community.ory.sh',
    label: 'Forum',
    position: 'left'
  },
  {
    href: 'https://www.ory.sh/chat',
    label: 'Chat',
    position: 'left'
  },
  {
    href: `https://github.com/ory/${githubRepoName}`,
    label: 'GitHub',
    position: 'left'
  }
]

module.exports = {
  title: config.projectName,
  tagline: config.projectTagLine,
  url: `https://www.ory.sh/`,
  baseUrl: `/${config.projectSlug}/docs/`,
  favicon: 'img/favico.png',
  organizationName: 'ory', // Usually your GitHub org/user name.
  projectName: config.projectSlug, // Usually your repo name.
  themeConfig: {
    googleAnalytics: {
      trackingID: 'UA-71865250-1',
      anonymizeIP: true
    },
    algolia: {
      apiKey: '8463c6ece843b377565726bb4ed325b0',
      indexName: 'ory',
      contextualSearch: true,
      searchParameters: {
        // facetFilters: [
        //   `tags:${config.projectSlug}`,
        // ]
      }
    },
    navbar: {
      logo: {
        alt: config.projectName,
        src: `img/logo-${config.projectSlug}.svg`,
        href: `https://www.ory.sh/${
          config.projectSlug === 'ecosystem' ? '' : config.projectSlug
        }`
      },
      items: [
        ...links,
        {
          type: 'docsVersionDropdown',
          position: 'right',
          dropdownActiveClassDisabled: true,
          dropdownItemsAfter: [
            {
              to: '/versions',
              label: 'All versions'
            }
          ]
        }
      ]
    },
    footer: {
      style: 'dark',
      copyright: `Copyright © ${new Date().getFullYear()} ORY GmbH`,
      links: [
        {
          title: 'Company',
          items: [
            {
              label: 'Imprint',
              href: 'https://www.ory.sh/imprint'
            },
            {
              label: 'Privacy',
              href: 'https://www.ory.sh/privacy'
            },
            {
              label: 'Terms',
              href: 'https://www.ory.sh/tos'
            }
          ]
        }
      ]
    }
  },
  plugins: [
    [
      '@docusaurus/plugin-content-docs',
      {
        path:
          config.projectSlug === 'docusaurus-template'
            ? 'contrib/docs'
            : 'docs',
        sidebarPath: require.resolve('./contrib/sidebar.js'),
        editUrl: `https://github.com/ory/${githubRepoName}/edit/master/docs`,
        routeBasePath: '/',
        showLastUpdateAuthor: true,
        showLastUpdateTime: true,
        remarkPlugins: [admonitions],
        disableVersioning: false
      }
    ],
    '@docusaurus/plugin-content-pages',
    '@docusaurus/plugin-google-analytics',
    '@docusaurus/plugin-sitemap'
  ],
  themes: [
    [
      '@docusaurus/theme-classic',
      {
        customCss:
          config.projectSlug === 'docusaurus-template'
            ? require.resolve('./contrib/theme.css')
            : require.resolve('./src/css/theme.css')
      }
    ],
    '@docusaurus/theme-search-algolia'
  ]
}
