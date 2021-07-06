---
layout: page
title: SysAdmin Cookbook - Jenkins
---

## Failed Builds

To troubleshoot the reason for a failed build:

1. In the Jenkins dashboard for the service that failed to build, find the "Build History" list in the left-hand column.
1. From this list, find your failed build (generally the most recent one) click on

Below we describe a couple of common issues that may cause a build to fail.

### Pull Access From ECR Denied

**The Problem:**

Upon

```text
Step 1/9 : FROM ************.dkr.ecr.us-east-1.amazonaws.com/content_commons_server:base
pull access denied for ************.dkr.ecr.us-east-1.amazonaws.com/content_commons_server, repository does not exist or may require 'docker login'
```

**The Fix:**

1. SSH to the Jenkins server (via the bastion server).
1. Switch to the Jenkins user with the following commands:

```bash
sudo -i
su jenkins
```

1. Log into ECR with the following command:

```bash
$(aws ecr get-login --no-include-email --region us-east-1)
```

**The Explanation:**

We store our customized Docker images in a private AWS Elastic Container Registry (ECR).

### No Space Left on Device

**The Problem:**

```bash
npm ERR! Error: ENOSPC: no space left on device, write
```

**The Fix:**

1. SSH to the Jenkins server (via the bastion server).
1. Switch to the root user with the following command:

```bash
sudo -i
```

1. Remove dangling Docker images by running:

```bash
docker image prune
```

1. If no space is reclaimed or the issue persists, remove all unused data using:

```bash
docker system prune -a
```

1. If the issue still persists, re-run the system prune command with the `--volumes` flag to also delete orphaned volumes:

```bash
docker system prune -a --volumes
```

Note that if you run system prune, the connection to ECR may be closed any you may need to log back in following the [instructions above](#pull-access-from-ecr-denied).

**The Explanation:**
