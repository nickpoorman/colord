# colord

Monitors your clipboard and when it finds a #FFF or #FFFFFF color code, quickly displays the color code.

I have been using Zed but I need to be able to quickly see the color for the color code when I am doing website design.

## Install

Installing will build the binaries and then add a plist file for the program to start up on boot. Macos will prompt you to let you know a "Login item" is being added.

```
make
```

## Remove

```
make clean
```

## Docs

This program is two parts. 1. A daemon that monitors the clipboard for color codes and launches the `colord_display` program. 2. the `colord_display` program displays the color in the bottom right corner of your screen and then shuts down. This program is in two parts because I didn't want a application sitting in my Macos dock all the time, this way the colord_display` app will terminate after displaying the color and remove the macos dock item will be removed.
