![Logo](examples/logo.png)

# vixl44

Create pixel art inside your terminal using vim movements

## Installation

[Get go](https://golang.org/dl/), then:

```
go get github.com/sebashwa/vixl44
```

## Usage

```
~/go/bin/vixl44 [OPTIONS] [FILENAME]

FILENAME

    the name of your file

OPTIONS

-c, --cols
    number of columns, default is 20, 0 means full width, ignored if name of existing file given
-r, --rows
    number of rows, default is 20, 0 means full height, ignored if name of existing file given
```

## Keybindings

### Movement

Keys                                                | Action
----------------------------------------------------|----------------------------------
<kbd>h</kbd>,<kbd>j</kbd>,<kbd>k</kbd>,<kbd>l</kbd> | Move cursor left, down, up, right
<kbd>w</kbd>                                        | Move cursor 5 columns right
<kbd>b</kbd>                                        | Move cursor 5 columns left
<kbd>Ctrl-d</kbd>                                   | Move cursor 5 rows down
<kbd>Ctrl-u</kbd>                                   | Move cursor 5 rows up

### Normal Mode

Keys                                                | Action
----------------------------------------------------|----------------------------------
<kbd>Space</kbd>, <kbd>Return</kbd>                 | Paint color
<kbd>s</kbd>                                        | Select color under the cursor
<kbd>f</kbd>                                        | Replace color in area ([Flood fill](https://en.wikipedia.org/wiki/Flood_fill))
<kbd>u</kbd>                                        | Undo change
<kbd><Ctrl-r</kbd>                                  | Redo change
<kbd>p</kbd>                                        | Paste from buffer
<kbd>c</kbd>                                        | Switch to palette mode
<kbd>Ctrl-v</kbd>                                   | Switch to visual block mode

### Visual Block Mode

Keys                                                | Action
----------------------------------------------------|----------------------------------
<kbd>Space</kbd>, <kbd>Return</kbd>                 | Paint color
<kbd>x</kbd>, <kbd>d</kbd>                          | Cut area
<kbd>y</kbd>                                        | Copy area
<kbd>Esc</kbd>                                      | Switch to normal mode
<kbd>c</kbd>                                        | Switch to palette mode

### Palette Mode

Keys                                                | Action
----------------------------------------------------|----------------------------------
<kbd>Space</kbd>, <kbd>Return</kbd>                 | Select color
<kbd>Esc</kbd>                                      | Switch to normal mode
<kbd>Ctrl-v</kbd>                                   | Switch to visual block mode

## Commands

```
:w FILENAME<CR>         - Write to FILENAME
:wq FILENAME<CR>        - Write to FILENAME and quit
:exp FILENAME.svg<CR>   - Export to FILENAME in svg format
:exp FILENAMExN.png<CR> - Export to FILENAME in png format with an optional scale factor N
:exp FILENAME.ansi<CR>  - Export ANSI escape sequences to FILENAME (use i.e. cat FILENAME to view it)
:q<CR>                  - Quit
```

## License

This project is free software: You can redistribute it and/or modify it under the terms of the [GNU General Public License](https://www.gnu.org/licenses/gpl.html), either version 3 of the License, or (at your option) any later version.
