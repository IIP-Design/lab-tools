---
title: GitHub Repository Review
date: 2022-04-26
excerpt: 'A review of the active repositories in the IIP-Design GitHub account and recommendations regarding archiving outdated code.'
---

## Summary

Benefits of archiving repos:

- The archive status signals to anyone looking at the code that it is no longer being maintained. This reduces the chance for misapplication of old code and keeps expectations on the team a bit more reasonable.
- Archiving a repo disables security alerts reducing noise and allowing us to focus our attentions on active projects.
- It more clearly delineates what we are/will work on and what we can safely disregard.

**Why not just delete?** It is reasonable to wonder why we don't just delete this unused code. While this approach would certainly cut down on the clutter, there are some benefits to keeping the old repos around in an archived state. Namely:

- If a third party is referencing one of our public repos they will not suffer a loss of service due to the archiving. A deletion on the other hand would cut them off from the source.
- We may at some future time need to revive a project. It is much easier to reactivate an archived repo than start from scratch.
- The archived repos provide a useful reference. We may need to solve similar problems in the future in which case it may be useful to look back into our past. Each one is a useful lesson in approaches both successful and unsuccessful.
- These old repos, even if not in use provide an interesting history of the team's evolution and an indication of the projects we've worked on in the past. Furthermore, they serve as a testament to the invaluable contributions of team members past and present.

Top line repo stats:

|          | Public | Private | Active | Archived | Total |
| -------- | ------ | ------- | ------ | -------- | ----- |
| current  | 94     | 62      | 143    | 13       | 156   |
| proposed | 92     | 60      | 91     | 61       | 152   |

## The Archiving Process

Archiving a repo is a pretty straight forward. We begin by adding the following banner to the top of the project's README file:

```md
# :warning: THIS REPO HAS BEEN ARCHIVED :warning:

This project {INSERT REASON FOR ARCHIVING HERE}. As such, this repository has been archived and is no longer being maintained. The code is preserved here as a reference.
```

If a project does not contain a README, add one that contains the banner (and if possible a brief description of the project for future reference). Commit the changes to the README using the commit message `docs: add archive message`.

With this done, go to the repo **Settings** tab and in the **Danger Zone** locate and click the `Archive this repository` button. Acknowledge the warning message to complete the archiving process. Note that once you archive a repo you not be able to make anymore changes to it, so complete any cleanup actions first.

## Repos to Archive

The bulk of the repos recommended for archiving pertain to the various WordPress sites that GPA Lab managed under its previous iteration (as the IIP Design Team).

### RFJ - 9

- [RewardsForJustice](https://github.com/IIP-Design/RewardsForJustice) - RFJ site webroot.
- [RewardForJustice-theme](https://github.com/IIP-Design/RewardForJustice-theme) - RFJ custom WordPress theme.
- [rfj-lang](https://github.com/IIP-Design/rfj-lang) - RFJ-oriented WordPress custom post type.
- [rfj-features](https://github.com/IIP-Design/rfj-features) - RFJ-oriented WordPress custom post type.
- [rfj-cpt-wanted](https://github.com/IIP-Design/rfj-cpt-wanted) - RFJ-oriented WordPress custom post type.
- [rfj-basic-page](https://github.com/IIP-Design/rfj-basic-page) - RFJ-oriented WordPress custom post type.
- [rfj-cpt-success-story](https://github.com/IIP-Design/rfj-cpt-success-story) - RFJ-oriented WordPress custom post type.
- [rfj-cpt-act-of-terror](https://github.com/IIP-Design/rfj-cpt-act-of-terror) - RFJ-oriented WordPress custom post type.
- [rfj-cpt-slide](https://github.com/IIP-Design/rfj-cpt-slide) - RFJ-oriented WordPress custom post type.

### Campaigns - 9

- [YALI](https://github.com/IIP-Design/YALI) - The YALI webroot.
- [yali-theme](https://github.com/IIP-Design/yali-theme) - The YALI custom theme (child theme of corona).
- [moxie-yali](https://github.com/IIP-Design/moxie-yali) - An older YALI custom theme.
- [gtm4wp](https://github.com/IIP-Design/gtm4wp) - Fork of the GTM plugin adjusted for YALI
- [YLAI](https://github.com/IIP-Design/YLAI) - The YLAI webroot.
- [ylai-theme](https://github.com/IIP-Design/ylai-theme) The YLAI custom theme (child theme of corona).
- [YSEALI](https://github.com/IIP-Design/YSEALI) - The (never released) YSEALI webroot.
- [moxie-yseali](https://github.com/IIP-Design/moxie-yseali)
- [ytili-microsite](https://github.com/IIP-Design/ytili-microsite)

### ShareAmerica - 9

- [ShareAmerica](https://github.com/IIP-Design/ShareAmerica) - ShareAmerica webroot.
- [ShareAmerica-Theme](https://github.com/IIP-Design/ShareAmerica-Theme) - ShareAmerica custom child theme.
- [Newspaper](https://github.com/IIP-Design/Newspaper) - Private version of 3rd party WordPress theme.
- [td-mobile-plugin](https://github.com/IIP-Design/td-mobile-plugin) - Private version of 3rd a Newspaper theme add-on.
- [td-newsletter](https://github.com/IIP-Design/td-newsletter) - Private version of 3rd a Newspaper theme add-on.
- [td-composer](https://github.com/IIP-Design/td-composer) - Private version of 3rd a Newspaper theme add-on.
- [td-social-counter](https://github.com/IIP-Design/td-social-counter) - Private version of 3rd a Newspaper theme add-on.
- [td-standard-pack](https://github.com/IIP-Design/td-standard-pack) - Private version of 3rd a Newspaper theme add-on.
- [td-cloud-library](https://github.com/IIP-Design/td-cloud-library) - Private version of 3rd a Newspaper theme add-on.

### Other WordPress - 9

- [Translations](https://github.com/IIP-Design/Translations) - The translations site webroot.
- [translations-theme](https://github.com/IIP-Design/translations-theme)
- [America.gov](https://github.com/IIP-Design/America.gov)
- [America-theme](https://github.com/IIP-Design/America-theme)
- [interactive-theme](https://github.com/IIP-Design/interactive-theme)
- [genesis-framework](https://github.com/IIP-Design/genesis-framework)
- [TechCamps](https://github.com/IIP-Design/TechCamps)
- [TechCamps-theme](https://github.com/IIP-Design/TechCamps-theme)
- [State-Magazine](https://github.com/IIP-Design/State-Magazine)

### WPML - 8

WPML is a suite of WordPress plugins used to translate blog content. It was used on ShareAmerica and (until recently) the Content site.

- [sitepress-multilingual-cms](https://github.com/IIP-Design/sitepress-multilingual-cms) - Private version of 3rd party product
- [wpml-cms-nav](https://github.com/IIP-Design/wpml-cms-nav) - Private version of 3rd party product
- [wpml-media](https://github.com/IIP-Design/wpml-media) - Private version of 3rd party product
- [wpml-sticky-links](https://github.com/IIP-Design/wpml-sticky-links) - Private version of 3rd party product
- [wpml-string-translation](https://github.com/IIP-Design/wpml-string-translation) - Private version of 3rd party product
- [wpml-translation-analytics](https://github.com/IIP-Design/wpml-translation-analytics) - Private version of 3rd party product
- [wpml-translation-management](https://github.com/IIP-Design/wpml-translation-management) - Private version of 3rd party product
- [wpml-url-fix](https://github.com/IIP-Design/wpml-url-fix) - Custom plugin that resolves WPML navigation issues for ShareAmerica

### Other - 4

- [cdp-app](https://github.com/IIP-Design/cdp-app) - An early client for use with the CDP API. Predecessor to Content Commons.
- [prettier-config](https://github.com/IIP-Design/prettier-config) - A share Prettier configuration. Can be deprecated since the team no longer uses Prettier.
- [simplesearchreact](https://github.com/IIP-Design/simplesearchreact) - A search interface used to query the CDP API. Preceded the development of Content Commons.
- [wp-elasticsearch-proxy-server](https://github.com/IIP-Design/wp-elasticsearch-proxy-server) - A server used to handle communication between Elasticsearch and WordPress instances. Predecessor to the cpd-public-api.

### Possibly - 31

- [america-page-protection](https://github.com/IIP-Design/america-page-protection)
- [america-screendoor](https://github.com/IIP-Design/america-screendoor)
- [america-plugin-manager](https://github.com/IIP-Design/america-plugin-manager)
- [call-to-action](https://github.com/IIP-Design/call-to-action)
- [cmb2-conditionals](https://github.com/IIP-Design/cmb2-conditionals) - Fork of 3rd party plugin adjusted for YALI
- [Flexible-Sidebar](https://github.com/IIP-Design/Flexible-Sidebar) - Was only used on now archived sites but functionality may still be useful
- [floating-social-bar](https://github.com/IIP-Design/floating-social-bar)
- [formidable](https://github.com/IIP-Design/formidable) - Private version of 3rd party product
- [formidable-bottombar](https://github.com/IIP-Design/formidable-bottombar)
- [formidable-pro](https://github.com/IIP-Design/formidable-pro) - Private version of 3rd party product
- [formidable-wpml](https://github.com/IIP-Design/formidable-wpml) - Private version of 3rd party product
- [heading-title-shift](https://github.com/IIP-Design/heading-title-shift)
- [iip-events](https://github.com/IIP-Design/iip-events)
- [iip-second-author](https://github.com/IIP-Design/iip-second-author)
- [iip-map](https://github.com/IIP-Design/iip-map)
- [inject-video](https://github.com/IIP-Design/inject-video)
- [js-cookie-manager](https://github.com/IIP-Design/js-cookie-manager) - Something for Share (?)
- [js_composer](https://github.com/IIP-Design/js_composer) - Private version of 3rd party product
- [mailchimp-for-wp-pro](https://github.com/IIP-Design/mailchimp-for-wp-pro) - Private version of 3rd party product
- [nextgen-gallery-pro](https://github.com/IIP-Design/nextgen-gallery-pro) - Private version of 3rd party product
- [node-wordpress](https://github.com/IIP-Design/node-wordpress) - Fork that was used to contribute back to source
- [parallelus-moxie](https://github.com/IIP-Design/parallelus-moxie)
- [password-protected](https://github.com/IIP-Design/password-protected)
- [revslider](https://github.com/IIP-Design/revslider) - Private version of 3rd party product
- [scrollmagic](https://github.com/IIP-Design/scrollmagic) - Private version of 3rd party product.
- [search-filter-pro](https://github.com/IIP-Design/search-filter-pro) - Private version of 3rd party product
- [sb-rss-feed-plus](https://github.com/IIP-Design/sb-rss-feed-plus)
- [wordpress-seo-premium](https://github.com/IIP-Design/wordpress-seo-premium) - Private version of 3rd party product
- [wp_core_languages](https://github.com/IIP-Design/wp_core_languages)
- [wp-simple-nonce](https://github.com/IIP-Design/wp-simple-nonce) - Fork of 3rd party source to make it work with Share (?)
- [corona-theme](https://github.com/IIP-Design/corona-theme)

## Repos to Delete - 4

- [cdp-wordpress-plugin](https://github.com/IIP-Design/cdp-wordpress-plugin) - Empty repo
- [react-youtube](https://github.com/IIP-Design/react-youtube) - Unchanged fork of public project
- [S3-Uploads](https://github.com/IIP-Design/S3-Uploads) - Fork of public repo to allow different PHP version, no longer used.
- [wp-es-feeder-publications](https://github.com/IIP-Design/wp-es-feeder-publications) - Intended as an adaptation of the feeder plugin for publications, never made it past the first commit with the initial setup.

## Repos to Merge - 12

into lab-test or lab-tools

- [lab-test-server](https://github.com/IIP-Design/lab-test-server) - lab test
- [taxonomy-parser](https://github.com/IIP-Design/taxonomy-parser) - lab tools/lab test?
- [create-iip-plugin](https://github.com/IIP-Design/create-iip-plugin) - lab tools
- [cognito-examples](https://github.com/IIP-Design/cognito-examples) - lab test
- [cdp-graphql-api](https://github.com/IIP-Design/cdp-graphql-api) - lab test
- [animation_builder](https://github.com/IIP-Design/animation_builder) - lab test

### Old Dev Setups

We went through various iterations of development setups, some of which never saw the light of day.

- [design-dev-box](https://github.com/IIP-Design/design-dev-box) - lab tools?
- [America.gov-Dev-Box](https://github.com/IIP-Design/America.gov-Dev-Box) - lab tools?
- [Share-Docker](https://github.com/IIP-Design/Share-Docker) - lab test - described as a proof of concept?
- [ShareAmerica-Dev-Box](https://github.com/IIP-Design/ShareAmerica-Dev-Box) - lab tools?
- [gcx-devbox](https://github.com/IIP-Design/gcx-devbox)
- [corona](https://github.com/IIP-Design/corona) - Basic WordPress setup intended for development of the corona theme.
