---
layout: page
title: How to Contribute To This Site
image: /assets/covers/contribute-cover.png
image_credit: Photo by Art Lasovsky on Unsplash
---

## Setup

The source files for this site reside in the GPA Lab Tools [GitHub repo](https://github.com/IIP-Design/lab-tools). Clone this Lab Tools repository to your local computer: `git clone git@github.com:IIP-Design/lab-tools.git`.

There are a number of other tools within this repository. Everything relevant to the site is found in the `docs` directory. You can navigate there using the command `cd lab-tools/docs`.

At any time, you can test your changes by running a local version of this site. To start the development site run command `bundle exec jekyll serve` from the `lab-tools/docs` directory. You can then navigate to localhost:4040 in your browser to see the site. To stop the development site, simple enter `Ctrl + c` in the terminal window where the dev server is running.

To suggest changes to the site, create a new branch (preferably named in a way that indicate which page you are adding/editing). When committing your changes, precede the commit message with a `docs:` identifier. Once you have finished, push up your branch and open a pull request to the `main` branch.

## General

Each page on the site is created from a markdown file. Generally speaking, you can use any features supported by [GitHub flavored markdown](https://docs.github.com/en/get-started/writing-on-github/getting-started-with-writing-and-formatting-on-github/basic-writing-and-formatting-syntax).

Special care should be taken when linking to other pages on the site. The urls are transformed when GitHub pages are built so by default, relative links will not work. Rather, you should use a standard markdown link augmented with liquid template markup to ensure proper linking. For example for the link `[link text](/dev_meetings/2022/05/26/#heading-on-the-page)` replace the contents of the parenthesis with `'/dev_meetings/2022/05/26/#heading-on-the-page' | relative_url` wrapped in two pairs of curly brackets. While a bit unwieldy, this format is the most robust.

Images or other static files should be uploaded to the `_docs/assets` directory under a `year/month` sub-directory matching the date of the meeting. They can then be embedded into the page using the method for relative links described above. Please ensure that you add alt text to any images or graphics added to the page.

## Adding to the Dev Meetings Notes

Create a new page. Dev meetings pages are found in the `docs/_dev_meetings` directory and organized according to the date on which they occurred. To create a new dev meeting page, add a markdown file in the appropriate year/month directory. The file name should be the day of the month on which the meeting occurred. Single digit months and days should be preceded with a 0. For example the notes from June 2, 2022 would have have the path `docs/_dev_meetings/2022/06/02.md`.

Add the top of the file, add the following YAML frontmatter fields:

- **title** - The title of the page which in this case is the date of the meeting in "Month DD, YYYY" format.
- **tags** - An array of topics covered in the meeting. There is no set list of allowed tags, but if possible favor reusing previous tags to adding new ones where appropriate.
- **date** - The date on which the meeting occurred in "YYYY-MM-DD" format. This may seem redundant, given the title field, but is useful for keeping the entries in order.
- **excerpt** - A very brief (ideally 120-220 characters) description of the meeting, when possible highlighting the most important topic of discussion.

The result should look something like this:

```yml
---
title: June 02, 2022
tags: [Topic A, Topic B]
date: 2022-06-06
excerpt: "Minutes from the June 02, 2022 GPA Lab developer's meeting. In which we have a good time."
---
```

The minutes are not intended to be a word-for-word transcription of the dev team's meetings. They are intended to identify topics of discussion, illustrate the team's thinking process, and above all record the decisions made by the team.

Since this site is public, individuals should be identified by their first names only (affiliation if from outside the team or for purposes of disambiguation is okay too). Any sensitive information or details should be omitted. If security is vulnerabilities are discussed, these notes should be excluded from the record until the vulnerability is resolved.

The page H1 is generated automatically from the title in the frontmatter and does not need to be added. Whenever possible the content of the meeting notes should be broken up into a few logically distinct sections. Each section should be indicated by an H2 identifying the topic of discussion.

Paragraphs with a description of the matter and a summation of the next steps to be taken are encouraged at the beginning and end of the section, respectively. Links to external resources or other dev meetings where the topic was discussed are extremely helpful, as are example code block where relevant.

Follow up notes or updates to the notes can be added after their publication. However, they should be preceded with an italicized, bolded, and dated update flag (for instance: `**_Update (06/02/2022):_** Here is my update.`)

### Deployment Checklists

Deployment checklists are a very common result of the dev meeting. They are constructed prior to a production release by when the whole team meets and verbally walks through the steps required for a successful deployment. These checklists follow a specific format outlined in the [Deployment Checklist Template]({{ 'sysadmin-cookbook/template-deployment-checklist' | relative_url }}).
