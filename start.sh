#!/usr/bin/env bash

day=$(date +%d)
year=$(date +%Y)

echo "Setting up for day $day..."

tmux new-session -s aoc${day} \; \
  split-window -fh \; \
  resize-pane -x 55 \; \
  send-keys "ram -d day${day}" C-m \; \
  split-window -v \; \
  send-keys "air --build.cmd 'go build -o bin/aoc ./' --build.bin './bin/aoc'" C-m \; \
  select-pane -t 1 \; \
    send-keys "go run cmd/generate.go ${day} && echo 'Press a key to start. Get your input from: https://adventofode.com/${year}/day/${day}' && read -rsn key && n ." C-m \; 


