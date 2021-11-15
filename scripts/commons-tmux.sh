#!/bin/sh

# Ensure that tmux is installed before proceeding.
if ! command -v tmux &> /dev/null
then
  echo ""
  echo "tmux could not be found...exiting the script."
  echo "To install you may want to run 'brew install tmux'."
  echo ""
  exit
fi

# Set the session name.
sn=commons

# Create a new tmux session initialized with the client window.
tmux new -s "$sn" -n 'client' -d

# List of additional windows to open.
windows=( 'server' 'api' 'worker' )

# Iterate through windows array opening each one.
for i in "${!windows[@]}"
do
  # Set the window number base off of the array index.
  # We add one to each index to account for the previously opened client window.
  num=$(($i + 1))

	tmux neww -t "$sn:$num" -n "${windows[$i]}"
done

# List of repos to open.
repos=( 'content-commons-client' 'content-commons-server' 'cdp-public-api' 'cdp-public-api')

# Loop through each window, opening the appropriate
# sub-directory and running the dev server.
for i in "${!repos[@]}"
do
  # Sets the base path to the repo directories based on the user input.
  path=$1

  tmux send-keys -t "$sn:$i" "cd $path${repos[$i]}" C-m

  if [ $i == 3 ];
  then
    # The fourth window (i.e. index 3) is the worker so the dev command is different.
    tmux send-keys -t "$sn:$i" 'npm run dev:worker' C-m
  else
    tmux send-keys -t "$sn:$i" 'npm run dev' C-m
  fi
done

# Select the first (client) window and attach it to the terminal.
tmux select-window -t "$sn:0"
tmux -2 attach-session -t "$sn"