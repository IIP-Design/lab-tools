---
layout: page_full
title: Commons tmux Script
image: /assets/covers/lab-tools-cover.png
---

## Purpose

This bash [script](https://github.com/IIP-Design/lab-tools/blob/main/scripts/commons-tmux.sh) allows a user to start up an existing Commons development site with a single command.

Specifically, the script uses [tmux](https://github.com/tmux/tmux/wiki) (a terminal multiplexing application) to initiate a new terminal session. Within this terminal session it opens four windows - client, server, API, and worker. In each window it navigates to the the appropriate directory and starts the development server. Once started the user can manage all four processes from within a single terminal session.

![A terminal window showing the user in the Projects/CDP/content-commons-client directory with the application running on localhost:3000. A bar at the bottom of the terminal window indicates the session is called commons, and that there are four available windows: 0 - client, 1 - server, 2 - api, and 3 - worker]({{ '/assets/2021/11/tmux.png' | relative_url }})

A couple assumptions must be true in order for the script to work:

1. The user has tmux installed on their machine ([notes on tmux](#tmux)).
1. The user has all of the Commons repositories ([client](https://github.com/IIP-Design/content-commons-client), [server](https://github.com/IIP-Design/content-commons-server), and [API](https://github.com/IIP-Design/cdp-public-api)) cloned into the same parent directory.

## Setup and Usage

1. Download the script. The easiest way to get the script is to simply pull it down as part of the GPA Lab tools repo by running the command

   ```bash
   git clone git@github.com:IIP-Design/lab-tools.git
   ```

   Within this repo, the script is found in the `scripts` directory in the file `commons-tmux.sh`. To get just the script and not the full repo, you can copy the [contents of the bash file](https://github.com/IIP-Design/lab-tools/blob/main/scripts/commons-tmux.sh) into a `.sh` file anywhere on your machine.

2. Change the permissions on the script to make it executable by running the following commands (from the `lab-tools` repo root - modify the path if running from elsewhere):

   ```bash
   chmod u+x scripts/commons-tmux.sh
   ```

3. You can now run the script. The script accepts one optional argument which is the path to the parent directory for all of the Commons repos.

   ```bash
   ./scripts/commons-tmux.sh <path/to/commons-repos/>
   ```

   If you prefer, you can place the script in the same directory as your Commons repos in which case you will not need to pass a path argument to the script. An even better approach is to alias the script in your bash profile (see step 4).

4. **(Optional)** We recommend setting an alias in your `.bashrc` file to more easily run the script. An alias provides several benefits by allowing you define the path to both the script and the Commons repos within the alias.Thereafter, you can start the tmux Commons session from any directory in your file system.

   For example assuming that the `lab-tools` repo is in `projects` directory and all the Commons repos are in a `cdp` sub-directory under `projects`, the alias might look like:

   ```bash
   alias commons='. $HOME/projects/lab-tools/scripts/commons-tmux.sh ~/projects/cdp/'
   ```

   The user can now run `commons` from anywhere to quickly start up all the commons development servers.

## tmux

This script depends on the presence of tmux on the user's system. The easiest way to install tmux on a Mac is via the command line using homebrew:

```bash
brew install tmux

tmux -V # To verify successful installation
```

Users will also need to familiarize themselves with the tmux commands in order to work effectively within the tmux environment. Commands directed at tmux itself (such as switching between windows) must be preceded by the input `ctrl + b`. You can detach from a tmux session, leaving it running in the background by typing, `ctrl + b`, `d`. To see what sessions are running in the background, type `tmux ls` in your terminal window. To reattach to a tmux session, enter `tmux a <session_number>` or `tmux a -t <session_name>`. To close a tmux window, type `exit` or `ctrl + d` (without the tmux prefix) when all windows are closed the session will close automatically.

The [tmux commands guide](https://tmuxguide.readthedocs.io/en/latest/tmux/tmux.html) cover the basics of application. Much more information can be found at the official [tmux docs wiki](https://github.com/tmux/tmux/wiki). Users can also customize their tmux environment with a `.tmux.conf` file.
