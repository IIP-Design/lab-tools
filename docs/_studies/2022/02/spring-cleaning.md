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

| Package            | Current | Latest | Update          | Breaking Changes                                                                                    |
| ------------------ | ------- | ------ | --------------- | --------------------------------------------------------------------------------------------------- |
| csv-parse          | ^4.16.3 | 5.0.4  | ✅              | Changed import path <sup>1</sup>, Renamed functions that we don't use <sup>2</sup>                  |
| dotenv             | ^10.0.0 | 16.0.0 | ✅              | Dropped Node 10 support, new features including inline comments and multi-line entries <sup>3</sup> |
| elasticsearch      | ^15.5.0 | 16.7.2 | ❓ <sup>4</sup> | Deprecated in favor of new package <sup>5</sup>                                                     |
| eslint             | ^7.32.0 | 8.12.0 | ✅              | Safe to update in conjunction with eslint config v1.6.0                                             |
| eslint-plugin-jest | ^24.4.2 | 26.1.3 | ✅              | Safe to update in conjunction with eslint config v1.6.0                                             |
| helmet             | ^4.4.1  | 5.0.2  | ✅              | Dropped Node 10 & 11 support, specify default values for a number of header <sup>6</sup>            |

<ul style="list-style:none;padding-left:1rem">
  <li><sup>1</sup> Import statement changed from <code>import parse from 'csv-parse/lib/sync';</code> to <code>import { parse } from 'csv-parse/sync'</code>. This change necessitated adding the <code>"ignore": "csv-parse/sync"</code> property to the <code>import/no-unresolved</code> ESLint rule.</li>
  <li><sup>2</sup> All changes are listed in the package's <a href="https://github.com/adaltas/node-csv/blob/master/packages/csv-parse/CHANGELOG.md">changelog</a></li>
  <li><sup>3</sup> All "breaking" changes are listed in the package's <a href="https://github.com/motdotla/dotenv/blob/master/CHANGELOG.md">changelog</a></li>
  <li><sup>4</sup> <a href="https://www.npmjs.com/package/@elastic/elasticsearch">@elastic/elasticsearch</a> - Elasticsearch client <a href="https://www.elastic.co/guide/en/elasticsearch/client/javascript-api/current/breaking-changes.html">migration guide</a></li>
  <li><sup>5</sup> Elasticsearch client announcement <a href="https://www.elastic.co/blog/new-elasticsearch-javascript-client-released">blog post</a></li>
 <li><sup>6</sup> The new header defaults are noted in the repo's <a href="https://github.com/helmetjs/helmet/blob/main/CHANGELOG.md#501---2022-01-03">changelog</a></li>
</ul>

## Client - Edwin

| Package                   | Current | Latest | Update          | Breaking Changes                                        |
| ------------------------- | ------- | ------ | --------------- | ------------------------------------------------------- |
| @ckeditor/ckeditor5-react | ^3.0.2  | 4.0.0  |                 |                                                         |
| @next/bundle-analyzer     | ^11.1.2 | 12.1.2 |                 |                                                         |
| babel-jest                | ^26.6.3 | 27.2.0 | ❓ <sup>1</sup> | Migrate to ESM                                          |
| cypress                   | ^8.4.1  | 9.5.3  |                 |                                                         |
| dotenv                    | ^10.0.0 | 16.0.0 |                 |                                                         |
| eslint                    | ^7.32.0 | 8.12.0 | ✅              | Safe to update in conjunction with eslint config v1.6.0 |
| eslint-plugin-jest        | ^24.4.2 | 26.1.3 | ✅              | Safe to update in conjunction with eslint config v1.6.0 |
| graphql                   | ^15.6.0 | 16.3.0 |                 |                                                         |
| lint-staged               | ^11.1.2 | 12.3.7 |                 |                                                         |
| next                      | ^10.1.3 | 12.1.2 | ❓ <sup>2</sup> | Uses Webpack 5 to compile <sup>3, 4</sup>               |
| next-redux-wrapper        | ^6.0.2  | 7.0.5  | ❓ <sup>5</sup> | Changed function signatures <sup>6</sup>                |
| node-sass                 | ^4.14.1 | 7.0.1  | ❓              | Dropped Node 10 support, remove process.sass API        |
| react-markdown            | ^5.0.3  | 8.0.1  | ❓              | Migrate to ESM                                          |

<ul style="list-style:none;padding-left:1rem">
  <li><sup>1</sup> Jest monorepo <a href="https://github.com/facebook/jest/releases/tag/v27.0.0">changelog</a></li>
  <li><sup>2</sup> Next.js <a href="https://nextjs.org/docs/upgrading">upgrade guide</a></li>
  <li><sup>3</sup> Next.js Webpack 5 <a href="https://nextjs.org/docs/messages/webpack5">adoption guide</a></li>
  <li><sup>4</sup> Next.js <a href="https://github.com/vercel/next.js/releases/tag/v11.0.0">changelog</a></li>
  <li><sup>5</sup> Next Redux wrapper <a href="https://github.com/kirill-konshin/next-redux-wrapper#upgrade-from-6x-to-7x">migration guide</a></li>
  <li><sup>6</sup> Changed signature of <code>createWrapper</code>, <code>GetStaticPropsContext</code> and <code>GetServerSidePropsContext</code> no longer exported, <code>getInitialProps</code> must be manually wrapped</li>
</ul>

## Server - Terri

| Package               | Current | Latest | Update          | Breaking Changes                                        |
| --------------------- | ------- | ------ | --------------- | ------------------------------------------------------- |
| apollo-server-express | ^2.25.2 | 3.6.6  | ❌ <sup>1</sup> | Many breaking changes listed <sup>2</sup>               |
| dotenv                | ^10.0.0 | 16.0.0 |                 |                                                         |
| eslint                | ^7.32.0 | 8.12.0 | ✅              | Safe to update in conjunction with eslint config v1.6.0 |
| eslint-plugin-jest    | ^24.4.2 | 26.1.3 | ✅              | Safe to update in conjunction with eslint config v1.6.0 |
| graphql               | ^14.7.0 | 16.3.0 | ❌ <sup>3</sup> | Many breaking changes listed <sup>4</sup>               |
| helmet                | ^4.6.0  | 5.0.2  |                 |                                                         |
| lint-staged           | ^11.1.2 | 12.3.7 |                 |                                                         |

<ul style="list-style:none;padding-left:1rem">
  <li><sup>1</sup> Apollo Server v3 <a href="https://www.apollographql.com/docs/apollo-server/migration/">migration guide</a></li>
  <li><sup>2</sup> Apollo Server <a href="https://github.com/apollographql/apollo-server/blob/main/CHANGELOG.md#v300">changelog</a></li>
  <li><sup>3</sup> Peer dependency conflict with <code>prisma-client-lib</code>. Install with <code>npm install --legacy-peer-deps</code></li>
  <li><sup>4</sup> graphql <a href="https://github.com/graphql/graphql-js/releases/tag/v15.0.0">changelog</a></li>
</ul>
