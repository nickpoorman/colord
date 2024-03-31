# colord

Monitors your clipboard and when it finds a #FFF or #FFFFFF color code, quickly displays the color code.

I have been using Zed but I need to be able to quickly see the color for the color code when I am doing website design.

## Install

Installing will build the binaries and place them in the the user's `~/bin` directory. It will also create a `colord` bash file to launch the `colord` daemon.

```
make
```

## Remove

```
make clean
```

## Docs

This program is two parts.

1. A daemon that monitors the clipboard for color codes and launches the `colord_display` program.
2. the `colord_display` program displays the color in the bottom right corner of your screen and then shuts down.

This program is in two parts because I didn't want a application sitting in my macos dock all the time, this way the colord_display` app will terminate after displaying the color and the macos dock item will be closed.
