![Logo](examples/logo.png)

# Installation

[Get go](https://golang.org/dl/), then:

```
go get github.com/sebashwa/vixl44
```

# Usage

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

# Key bindings

```
MOVEMENT

h,j,k,l - Move cursor left, down, up, right
w       - Move cursor 5 columns right
b       - Move cursor 5 columns left
<C-D>   - Move cursor 5 rows down
<C-U>   - Move cursor 5 rows up

NORMAL MODE

<Space>, <CR>    - Paint color
s                - Select color under the cursor
f                - Replace color in area (Flood fill)
u                - Undo change
<C-R>            - Redo change
p                - Paste from buffer
c                - Switch to palette mode
<C-V>            - Switch to visual block mode

VISUAL BLOCK MODE

<Space>, <CR> - Paint color
x, d          - Cut area
y             - Copy area
<ESC>         - Switch to normal mode
c             - Switch to palette mode

PALETTE MODE

<Space>, <CR>    - Select color
<ESC>            - Switch to normal mode
<C-V>            - Switch to visual block mode
```

# Commands

```
:w FILENAME<CR>         - Write to FILENAME
:wq FILENAME<CR>        - Write to FILENAME and quit
:exp FILENAME.svg<CR>   - Export to FILENAME in svg format
:exp FILENAMExN.png<CR> - Export to FILENAME in png format with an optional scale factor N
:q<CR>                  - Quit
```

# Todo

- [x] Implement export to SVG
- [x] Implement export to PNG
- [x] Implement a history
- [x] Implement [Flood fill](https://en.wikipedia.org/wiki/Flood_fill)
- [x] Make copy/cut & paste possible
- [ ] Add export for terminal output
