---
title: GitHub Repository Review
date: 2022-04-26
excerpt: 'A review of the active repositories in the IIP-Design GitHub account and recommendations regarding archiving outdated code.'
---

## Summary

The [IIP-Design GitHub](https://github.com/IIP-Design) is where the Lab team (and the IIP Design team before it) stores source code for its projects and websites. Over the years during which this account has been used, the number of repositories has proliferated dramatically. The account now contains 156 repositories. These repos include both code written by the team and third party code stored in private repos for use in automated code pipelines. Many of these projects have outlived their usefulness and are no longer being maintained. As such, we recommend archiving the bulk of these old repositories (and deleting or merging a much smaller number).

Benefits of archiving repos:

- The archive status signals to anyone looking at the code that it is no longer being maintained. This reduces the chance for misapplication of old code and keeps expectations on the team regarding support a bit more reasonable.
- Archiving a repo disables security alerts reducing noise and allowing us to focus our attentions on active projects.
- It more clearly delineates what we can/will work on and what we can safely disregard.

If these repos are no longer used, _**why not just delete them?**_ It is reasonable to wonder why we don't just delete this unused code. While this approach would certainly cut down on the clutter, there are some benefits to keeping the old repos around in an archived state. Namely:

- If a third party is referencing one of our public repos they will not suffer a loss of service due to the archiving. A deletion, on the other hand, would cut them off from source code that may be integrated in their build process.
- We may at some future time need to revive a project. It is much easier to reactivate an archived repo than start from scratch.
- The archived repos provide a useful reference. Each one provides useful lessons in approaches both successful and unsuccessful. We may need to solve similar problems in the future in which case it may be useful to review these lessons.
- These old repos, even if not in active use, provide an interesting history of the team's evolution and an indication of the projects we've worked on in the past. Furthermore, they serve as a testament to the invaluable contributions of team members past and present.

Following the recommendations outlined in this guide will not dramatically reduce the number of repositories in the IIP-Design account. However, it will bring great clarity regarding which codebases are under active development.

|          | Active | Archived | Public | Private | Total |
| -------- | ------ | -------- | ------ | ------- | ----- |
| current  | 143    | 13       | 94     | 62      | 156   |
| proposed | 85     | 55       | 83     | 57      | 140   |

## The Archiving Process

Archiving a repo is a pretty straight forward. We begin by adding the following banner to the top of the project's README file:

```md
# :warning: THIS REPO HAS BEEN ARCHIVED :warning:

This project {INSERT REASON FOR ARCHIVING HERE}. As such, this repository has been archived and is no longer being maintained. The code is preserved here as a reference.
```

If a project does not contain a README, add one that contains the banner (and if possible a brief description of the project for future reference). Commit the changes to the README using the commit message `docs: add archive message`.

With this done, go to the repo **Settings** tab and in the **Danger Zone** locate and click the `Archive this repository` button. Acknowledge the warning message to complete the archiving process. Note that once you archive a repo you not be able to make anymore changes to it, so complete any cleanup actions first.

## Repos to Archive

The bulk of the repos recommended for archiving pertain to the various WordPress sites that GPA Lab managed under its previous iteration (as the IIP Design team). Over the course of the past two-three years these sites have been transitioned over to the Web team or external third parties, who have established their own deployment pipeline/source code base.

### Campaigns - 9

The YALI and YLAI campaign sites were transitioned over to the Web team in October of 2020. They have since been moved onto the MWP platform and we do not believe that any of the original codebase is still in use. YSEALI and YTILI were more limited campaigns and the code for them was not extensively used.

- [YALI](https://github.com/IIP-Design/YALI) - The YALI webroot.
- [yali-theme](https://github.com/IIP-Design/yali-theme) - The YALI custom theme (child theme of corona).
- [moxie-yali](https://github.com/IIP-Design/moxie-yali) - An older YALI custom theme (child theme of Parallelus Moxie).
- [gtm4wp](https://github.com/IIP-Design/gtm4wp) - Fork of the GTM plugin adjusted for YALI
- [YLAI](https://github.com/IIP-Design/YLAI) - The YLAI webroot.
- [ylai-theme](https://github.com/IIP-Design/ylai-theme) The YLAI custom theme (child theme of corona).
- [YSEALI](https://github.com/IIP-Design/YSEALI) - The (never released) YSEALI webroot.
- [moxie-yseali](https://github.com/IIP-Design/moxie-yseali) - The YSEALI custom theme (child theme of Parallelus Moxie).
- [ytili-microsite](https://github.com/IIP-Design/ytili-microsite) - An embeddable microsite that was added to a ShareAmerica page

### RFJ - 9

The Rewards for Justice WordPress site was transitioned over to a team within DS on January 19, 2022. This team stood up a brand new WordPress instance and custom theme. As far as we know they are not utilizing any of the legacy code below.

- [RewardsForJustice](https://github.com/IIP-Design/RewardsForJustice) - RFJ site webroot.
- [RewardForJustice-theme](https://github.com/IIP-Design/RewardForJustice-theme) - RFJ custom WordPress theme.
- [rfj-lang](https://github.com/IIP-Design/rfj-lang) - RFJ-oriented WordPress custom post type.
- [rfj-features](https://github.com/IIP-Design/rfj-features) - RFJ-oriented WordPress custom post type.
- [rfj-cpt-wanted](https://github.com/IIP-Design/rfj-cpt-wanted) - RFJ-oriented WordPress custom post type.
- [rfj-basic-page](https://github.com/IIP-Design/rfj-basic-page) - RFJ-oriented WordPress custom post type.
- [rfj-cpt-success-story](https://github.com/IIP-Design/rfj-cpt-success-story) - RFJ-oriented WordPress custom post type.
- [rfj-cpt-act-of-terror](https://github.com/IIP-Design/rfj-cpt-act-of-terror) - RFJ-oriented WordPress custom post type.
- [rfj-cpt-slide](https://github.com/IIP-Design/rfj-cpt-slide) - RFJ-oriented WordPress custom post type.

### ShareAmerica - 9

ShareAmerica was also transitioned over to the Web team in October of 2020. In their initial deployment the Web team mirrored these repos into their own private repos. We do not believe that they are currently referencing and of our repositories in their build pipeline.

- [ShareAmerica](https://github.com/IIP-Design/ShareAmerica) - ShareAmerica webroot.
- [ShareAmerica-Theme](https://github.com/IIP-Design/ShareAmerica-Theme) - ShareAmerica custom child theme.
- [Newspaper](https://github.com/IIP-Design/Newspaper) - Private version of a 3rd party WordPress theme.
- [td-mobile-plugin](https://github.com/IIP-Design/td-mobile-plugin) - Private version of a 3rd a Newspaper theme add-on.
- [td-newsletter](https://github.com/IIP-Design/td-newsletter) - Private version of a 3rd a Newspaper theme add-on.
- [td-composer](https://github.com/IIP-Design/td-composer) - Private version of a 3rd a Newspaper theme add-on.
- [td-social-counter](https://github.com/IIP-Design/td-social-counter) - Private version of a 3rd a Newspaper theme add-on.
- [td-standard-pack](https://github.com/IIP-Design/td-standard-pack) - Private version of a 3rd a Newspaper theme add-on.
- [td-cloud-library](https://github.com/IIP-Design/td-cloud-library) - Private version of a 3rd a Newspaper theme add-on.

### Other WordPress Sites - 9

Several more WordPress sites that were transitioned in the October-November 2020 timeframe. Some are now managed by Web others by external third parties.

- [Translations](https://github.com/IIP-Design/Translations) - The Translations site webroot.
- [translations-theme](https://github.com/IIP-Design/translations-theme) - An adaptation of the WordPress default twentyfifteen theme tailored for the Translations site.
- [America.gov](https://github.com/IIP-Design/America.gov) - Publications/Interactive multisite webroot.
- [America-theme](https://github.com/IIP-Design/America-theme) - The Publications site custom theme (child theme of genesis).
- [interactive-theme](https://github.com/IIP-Design/interactive-theme) - The Interactive site custom theme (child theme of corona).
- [genesis-framework](https://github.com/IIP-Design/genesis-framework) - Private version of a third party starter theme. Used as the parent theme for America-theme.
- [TechCamps](https://github.com/IIP-Design/TechCamps) - The TechCamps webroot.
- [TechCamps-theme](https://github.com/IIP-Design/TechCamps-theme) - The TechCamps site custom theme (child theme of corona).
- [State-Magazine](https://github.com/IIP-Design/State-Magazine) - The State Magazine webroot.

### WPML - 8

WPML is a suite of WordPress plugins used to translate blog content. It was used on ShareAmerica and (until recently) the Content site.

- [sitepress-multilingual-cms](https://github.com/IIP-Design/sitepress-multilingual-cms) - Private version of a 3rd party translation plugin.
- [wpml-cms-nav](https://github.com/IIP-Design/wpml-cms-nav) - Private version of a WPML add-on.
- [wpml-media](https://github.com/IIP-Design/wpml-media) - Private version of a WPML add-on.
- [wpml-sticky-links](https://github.com/IIP-Design/wpml-sticky-links) - Private version of a WPML add-on.
- [wpml-string-translation](https://github.com/IIP-Design/wpml-string-translation) - Private version of a WPML add-on.
- [wpml-translation-analytics](https://github.com/IIP-Design/wpml-translation-analytics) - Private version of a WPML add-on.
- [wpml-translation-management](https://github.com/IIP-Design/wpml-translation-management) - Private version of a WPML add-on.
- [wpml-url-fix](https://github.com/IIP-Design/wpml-url-fix) - Custom plugin that resolves WPML navigation issues for ShareAmerica

### Non-WordPress - 4

The following repos are unrelated to WordPress. They are either early versions of current projects or (in the case of prettier-config) a tool that is no longer used by the team.

- [cdp-app](https://github.com/IIP-Design/cdp-app) - An early client for use with the CDP API. Predecessor to Content Commons.
- [prettier-config](https://github.com/IIP-Design/prettier-config) - A shared Prettier configuration. Can be deprecated since the team no longer uses Prettier.
- [simplesearchreact](https://github.com/IIP-Design/simplesearchreact) - A search interface used to query an early version of the CDP API. Preceded the development of Content Commons.
- [wp-elasticsearch-proxy-server](https://github.com/IIP-Design/wp-elasticsearch-proxy-server) - A server used to handle communication between Elasticsearch and WordPress instances. Predecessor to the cpd-public-api.

### Possibly Archivable - 31

There are a large number of repos that were used in previous team projects. Again most of these repos pertain to WordPress sites that we no longer manage.

- [america-page-protection](https://github.com/IIP-Design/america-page-protection) Plugin developed by the team to add password protected pages to sites using the america-theme (ie. formerly Publications and Interactive).
- [america-screendoor](https://github.com/IIP-Design/america-screendoor) - Plugin developed by the team for YALI.
- [america-plugin-manager](https://github.com/IIP-Design/america-plugin-manager) - Plugin developed by the team to manage WordPress plugin setup.
- [call-to-action](https://github.com/IIP-Design/call-to-action) - Plugin developed by the team to add a call-to-action field to the page edit screens for homepage templates. Used on YALI.
- [cmb2-conditionals](https://github.com/IIP-Design/cmb2-conditionals) - Fork of 3rd party plugin adjusted for YALI.
- [Flexible-Sidebar](https://github.com/IIP-Design/Flexible-Sidebar) - Plugin developed by the team to add a more useful site plugin. Was only used on now archived sites (ie. YALI) but functionality may still be useful.
- [floating-social-bar](https://github.com/IIP-Design/floating-social-bar) - Private version of a 3rd party plugin.
- [formidable](https://github.com/IIP-Design/formidable) - Private version of a 3rd party plugin. Used on ShareAmerica, YALI, and YLAI.
- [formidable-bottombar](https://github.com/IIP-Design/formidable-bottombar) - Developed by the team to embed a formidable form on an element fixed to the bottom of a given page. Used on ShareAmerica.
- [formidable-pro](https://github.com/IIP-Design/formidable-pro) - Private version of a 3rd party plugin.
- [formidable-wpml](https://github.com/IIP-Design/formidable-wpml) - Private version of a 3rd party add-on for the formidable plugin.
- [heading-title-shift](https://github.com/IIP-Design/heading-title-shift) - Plugin developed for ShareAmerica.
- [iip-events](https://github.com/IIP-Design/iip-events) - Plugin developed by the team for the Interactive site.
- [iip-second-author](https://github.com/IIP-Design/iip-second-author) - Plugin developed by the team for ShareAmerica.
- [iip-map](https://github.com/IIP-Design/iip-map) - Plugin developed by the team for YALI.
- [inject-video](https://github.com/IIP-Design/inject-video) - A jQuery plugin that replaces a placeholder image with an embedded YouTube video. Unclear if used anywhere.
- [js-cookie-manager](https://github.com/IIP-Design/js-cookie-manager) - Works with WPML to properly update site cookies when switching languages. Used on ShareAmerica.
- [js_composer](https://github.com/IIP-Design/js_composer) - Private version of a 3rd party plugin used on the ShareAmerica, YALI, and YLAI sites.
- [mailchimp-for-wp-pro](https://github.com/IIP-Design/mailchimp-for-wp-pro) - Private version of a 3rd party plugin used on the YALI site.
- [nextgen-gallery-pro](https://github.com/IIP-Design/nextgen-gallery-pro) - Private version of a 3rd party plugin used on the State Magazine site.
- [node-wordpress](https://github.com/IIP-Design/node-wordpress) - Fork of a 3rd party project that was used to contribute back to source.
- [parallelus-moxie](https://github.com/IIP-Design/parallelus-moxie) - A 3rd party starter theme used as the parent theme for some early IIP Design managed WordPress sites.
- [password-protected](https://github.com/IIP-Design/password-protected) - Private version of a 3rd party plugin. Unclear if ever used.
- [revslider](https://github.com/IIP-Design/revslider) - Private version of a 3rd party plugin used on the State Magazine site.
- [scrollmagic](https://github.com/IIP-Design/scrollmagic) - Private version of a 3rd party plugin. Acquired for use on the Stories/Policy site but never really used.
- [search-filter-pro](https://github.com/IIP-Design/search-filter-pro) - Private version of a 3rd party plugin possibly used on the Publications site.
- [sb-rss-feed-plus](https://github.com/IIP-Design/sb-rss-feed-plus) - Private version of a 3rd party plugin used on the ShareAmerica site.
- [wordpress-seo-premium](https://github.com/IIP-Design/wordpress-seo-premium) - Private version of a 3rd party plugin. Used on ShareAmerica, YALI, and YLAI.
- [wp_core_languages](https://github.com/IIP-Design/wp_core_languages) - Translation files for IIP Design managed WordPress sites.
- [wp-simple-nonce](https://github.com/IIP-Design/wp-simple-nonce) - Fork of 3rd party source to make it work with ShareAmerica (I think).
- [corona-theme](https://github.com/IIP-Design/corona-theme) - It has been a long time since it has been updated but it is still in use by the Courses site. We should hold off on archiving until the Courses site is transitioned to some other theme.

## Repos to Merge - 12

There are two unique repositories in the IIP-Design account. The [lab-test](https://github.com/IIP-Design/lab-test) repo was designed as a single place to test out new ideas and store proofs of concept. The [lab-tools](https://github.com/IIP-Design/lab-tools) repo was intended as a library of internal tooling and resources. There are a number of repositories in our account that were created to demonstrate a proof of concept. As experiments and/or internal tools, these projects were never widely distributed or used in a production context. As such they can safely be merged into one of the two multi-repos above and deleted after merging.

- [lab-test-server](https://github.com/IIP-Design/lab-test-server) - A proof of concept for sending emails using AWS Lambda functions. Can be merge into the lab-test repository.
- [taxonomy-parser](https://github.com/IIP-Design/taxonomy-parser) - A one-off tool used to scrap the state.gov countries taxonomy. Can be merged into the lab-tools repository.
- [create-iip-plugin](https://github.com/IIP-Design/create-iip-plugin) - An incomplete CLI tool used to generate boilerplate files for WordPress templates. Should be merged into the lab-tools repository.
- [cognito-examples](https://github.com/IIP-Design/cognito-examples) - An initial proof of concept for a server using Cognito authentication. Can be merged into the lab-test repository.
- [cdp-graphql-api](https://github.com/IIP-Design/cdp-graphql-api) - An initial experiment at setting up a GraphQL server. Can be merge into the lab-test repository. Already archived.
- [animation_builder](https://github.com/IIP-Design/animation_builder) - A proof of concept for the styled block story telling project (that eventually became the styled-block-builder). Can be merged into the lab-test repository. Already archived.

### Old Dev Setups

We went through various iterations of development setups, some of which never saw the light of day others of which were used but are no longer relevant. It might be worthwhile merging these repos into the lab-tools or lab-test repo as appropriate.

- [design-dev-box](https://github.com/IIP-Design/design-dev-box) - A VCCW-based development environment for the legacy WordPress sites. Already archived.
- [America.gov-Dev-Box](https://github.com/IIP-Design/America.gov-Dev-Box) - An extension of design dev box for development of the america.gov site (ie. Publications and Interactive). Already archived.
- [Share-Docker](https://github.com/IIP-Design/Share-Docker) - An unused proof of concept for running a development ShareAmerica environment in Docker. Already archived.
- [ShareAmerica-Dev-Box](https://github.com/IIP-Design/ShareAmerica-Dev-Box) - An extension of design dev box for development of the ShareAmerica site. Already archived
- [gcx-devbox](https://github.com/IIP-Design/gcx-devbox) A Docker-based setup for running development versions of the Lab-managed WordPress sites.
- [corona](https://github.com/IIP-Design/corona) - Basic WordPress setup intended for development of the corona theme.

## Repos to Delete - 4

Finally, there is a limited number of repositories in our account for which there is little to no value in preserving. Specifically, these were either projects that never got off the ground or are un-modified forks of third party codebases. As such, we recommend fully deleting these repos.

- [cdp-wordpress-plugin](https://github.com/IIP-Design/cdp-wordpress-plugin) - Empty repo
- [react-youtube](https://github.com/IIP-Design/react-youtube) - Unchanged fork of public project
- [S3-Uploads](https://github.com/IIP-Design/S3-Uploads) - Fork of public repo to allow different PHP version, no longer used.
- [wp-es-feeder-publications](https://github.com/IIP-Design/wp-es-feeder-publications) - Intended as an adaptation of the feeder plugin for publications, never made it past the first commit with the initial setup.
