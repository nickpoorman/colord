# colord

Display the color of the selected hex color code instantly using `colord_display`. Additionally, `colord_monitor` can be used to monitor your clipboard for copied #FFF or #FFFFFF color codes and quickly display the color.

![ColorPickerTask](https://github.com/nickpoorman/colord/assets/233103/6d3eadf4-8943-4ea0-8bcf-a52acd972ce0)

I have been using [Zed](#zed-integration) but I need to be able to quickly see the color for the color code when I am doing website design.

## Install / Build

Installing will build the binaries and place them in the the user's `~/bin` directory. It will also create a `colord` bash file to easily launch a `colord` clipboard daemon.

```
make
```

## Remove

```
make clean
```

## Zed Integration

I use this in [Zed](https://zed.dev) to quickly see the color code. Add the following to your Zed config files:

~/.config/zed/tasks.json

```
{
  "label": "HEX Color Diplay",
  "command": "colord_display \"$ZED_SELECTED_TEXT\" 1 50",
  "reveal": "never"
}
```

_colord_display_ arguments: 1 is the number of seconds to display the color, 50 is the window width.

~/.config/zed/keymap.json

```
{
  "context": "Workspace",
  "bindings": {
    "alt-g": ["task::Spawn", { "task_name": "HEX Color Diplay" }]
  }
}
```

## Docs

This program is two parts.

1. A daemon that monitors the clipboard for color codes and launches the `colord_display` program.
2. the `colord_display` program displays the color in the bottom right corner of your screen and then shuts down.

This program is in two parts because I didn't want a application sitting in my macos dock all the time, this way the colord_display` app will terminate after displaying the color and the macos dock item will be closed.
