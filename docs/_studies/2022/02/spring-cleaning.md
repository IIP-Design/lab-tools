---
title: Commons Dependency 2022 Spring Cleaning
date: 2022-02-23
excerpt: 'In preparation of adding Commons support for reports, we review the existing file data structures to inform the creation of a new Document type.'
---

The devs have each upgraded one of the environments up to the minor versions of each dependency. Below is an assessment of outstanding major version dependencies updates and their potential impact. The following key indicates what each icon in the "Update" column signifies:

|----|--------------------------------------------------------------------------------------|
| ✅ | We can safely update now. |
| ❌ | Will break the app or require significant refactoring. Hold off on updating for now. |
| ❓ | Needs further investigation before we can decide. |

## API - Marek

| Package       | Current | Latest | Update          | Breaking Changes                                |
| ------------- | ------- | ------ | --------------- | ----------------------------------------------- |
| elasticsearch | ^15.5.0 | 16.7.2 | ❓ <sup>1</sup> | Deprecated in favor of new package <sup>2</sup> |

<ul style="list-style:none;padding-left:1rem">
  <li><sup>1</sup> <a href="https://www.npmjs.com/package/@elastic/elasticsearch">@elastic/elasticsearch</a> - Elasticsearch client <a href="https://www.elastic.co/guide/en/elasticsearch/client/javascript-api/current/breaking-changes.html">migration guide</a></li>
  <li><sup>2</sup> Elasticsearch client announcement <a href="https://www.elastic.co/blog/new-elasticsearch-javascript-client-released">blog post</a></li>
</ul>

## Client - Edwin

| Package            | Current | Latest | Update          | Breaking Changes                                 |
| ------------------ | ------- | ------ | --------------- | ------------------------------------------------ |
| babel-jest         | ^26.6.3 | 27.2.0 | ❓ <sup>1</sup> | Migrate to ESM                                   |
| next               | ^10.1.3 | 11.1.2 | ❓ <sup>2</sup> | Uses Webpack 5 to compile <sup>3, 4</sup>        |
| next-redux-wrapper | ^6.0.2  | 7.0.5  | ❓ <sup>5</sup> | Changed function signatures <sup>6</sup>         |
| node-sass          | ^4.14.1 | 6.0.1  | ❓              | Dropped Node 10 support, remove process.sass API |
| react-markdown     | ^5.0.3  | 7.0.1  | ❓              | Migrate to ESM                                   |

<ul style="list-style:none;padding-left:1rem">
  <li><sup>1</sup> Jest monorepo <a href="https://github.com/facebook/jest/releases/tag/v27.0.0">changelog</a></li>
  <li><sup>2</sup> Next.js <a href="https://nextjs.org/docs/upgrading">upgrade guide</a></li>
  <li><sup>3</sup> Next.js Webpack 5 <a href="https://nextjs.org/docs/messages/webpack5">adoption guide</a></li>
  <li><sup>4</sup> Next.js <a href="https://github.com/vercel/next.js/releases/tag/v11.0.0">changelog</a></li>
  <li><sup>5</sup> Next Redux wrapper <a href="https://github.com/kirill-konshin/next-redux-wrapper#upgrade-from-6x-to-7x">migration guide</a></li>
  <li><sup>6</sup> Changed signature of <code>createWrapper</code>, <code>GetStaticPropsContext</code> and <code>GetServerSidePropsContext</code> no longer exported, <code>getInitialProps</code> must be manually wrapped</li>
</ul>

## Server - Terri

| Package               | Current | Latest | Update          | Breaking Changes                          |
| --------------------- | ------- | ------ | --------------- | ----------------------------------------- |
| apollo-server-express | ^2.21.2 | 3.3.0  | ❌ <sup>1</sup> | Many breaking changes listed <sup>2</sup> |
| graphql               | ^14.7.0 | 15.5.3 | ❌ <sup>3</sup> | Many breaking changes listed <sup>4</sup> |

<ul style="list-style:none;padding-left:1rem">
  <li><sup>1</sup> Apollo Server v3 <a href="https://www.apollographql.com/docs/apollo-server/migration/">migration guide</a></li>
  <li><sup>2</sup> Apollo Server <a href="https://github.com/apollographql/apollo-server/blob/main/CHANGELOG.md#v300">changelog</a></li>
  <li><sup>3</sup> Peer dependency conflict with <code>prisma-client-lib</code>. Install with <code>npm install --legacy-peer-deps</code></li>
  <li><sup>4</sup> graphql <a href="https://github.com/graphql/graphql-js/releases/tag/v15.0.0">changelog</a></li>
</ul>
