---
# https://vitepress.dev/reference/default-theme-home-page
layout: home

hero:
  name: "Moonlogs"
  text: 'Some crazy shit'
  tagline: 'Business-event logging with ease'
  image:
    src: /logo.svg
    alt: Moonlogs
  actions:
    - theme: brand
      text: Get Started
      link: /tutorial/install

features:
  - icon: 🗄️
    title: Events schemas based on business domains
    details: Create separate schemas to categorize events by domain areas. Events within each schema are recorded independently, facilitating efficient event retrieval.
  - icon: 🗃️
    title: Query-based log subgrouping
    details: Group events within a schema based on specified queries to enhance information integrity. Be sure unrelated events remain separate even if in the same schema.
  - icon: 🔍
    title: Convenient schema-based filters
    details: Generate convenient filters on the web interface for each schema, simplifying event search by allowing users to simply input values. No more DSL for trivial things.
  - icon: 🏷️
    title: Granular access control with tags
    details: Create and assign tags to schemas and users, enabling granular access control. Define access privileges based on tags, ensuring that users can only access the schemas and events relevant to their responsibilities.
---

