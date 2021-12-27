---
title: Deployment Checklist Template
---

## Deployment Checklist

**Update Window:** Wednesday January 3, 2:00 PM **_[Date of the deployment]_**

**Personnel:** **_[Team members (first name only) available to support the deployment]_**

**Content Freeze:** Four hours (2:00 PM - 6:00 PM) **_[Length of time and hours (in EST) of content freeze if required]_**

**Order of Operation:**

1. Update changelogs for **_[impacted repos]_**
1. Tag the releases
   - **_[Repo to be tagged]_** : **_[release type (i.e. major, minor, patch)]_** -> **_[release semantic version]_**
   - Client: patch -> 5.11.3
1. Update AWS permissions for **_[resource name]_** to allow **_[reason why change is required]_**
1. Take snapshots for backup ( **_[List services to backup - ex. Prisma DB, Elasticsearch]_** )
1. Set the [requisite environmental variables](#envvars) on the **_[environment (i.e. client, server)]_**
1. Run the builds via Jenkins:
   - [**_List the jobs to be run in the order in which they should be executed._**]
   - [**_Each job should be it's own bullet point. If run with parameter specify here._**]
1. Run Prisma **_[operation (i.e. deploy, seed)]_** via the the Jenkins job.
1. Any operations to be run directly from GraphQL or Kibana **_[explain the changes and if possible include code snippets below]_**
1. Update the `ui.json` file in the static assets to **_[purpose of change]_**.
1. Any changes to the Cloudflare settings.
1. Any clean up operations that should be performed.

**Code Snippets:**

**_[Any code snippets that maybe useful during the course of the deployment. This allows the deployment manager to just copy and paste the requisite code reducing the chance of errors. These snippets should ALWAYS be verified before running on a production environment. Also, DO NOT include any sensitive values in the snippet.]_**

<details id="envvars" open><summary>Sample Environmental Variables to Add</summary>
  <pre><code class="language-bash hljs">
SAMPLE_ENVAR=
  </code></pre>
</details>
