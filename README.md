![Logo](logo.svg)

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

s       - Select color under the cursor
<CR>    - Paint color
<Space> - Paint color
u       - Undo change
<C-R>   - Redo change
c       - Switch to palette mode
<C-V>   - Switch to visual block mode

VISUAL BLOCK MODE

<CR>    - Paint color
<Space> - Paint color
<ESC>   - Switch to normal mode
c       - Switch to palette mode

PALETTE MODE

<CR>    - Select color
<Space> - Select color
<ESC>   - Switch to normal mode
<C-V>   - Switch to visual block mode
```

# Commands

```
:w FILENAME<CR>   - Write to FILENAME
:wq FILENAME<CR>  - Write to FILENAME and quit
:exp FILENAME<CR> - Export to FILENAME (SVG or PNG)
:q<CR>            - Quit
```

# Todo

- [x] Implement export to SVG
- [x] Implement export to PNG
- [x] Implement a history
- [ ] Implement [Flood fill](https://en.wikipedia.org/wiki/Flood_fill)
- [ ] Make copy/cut & paste possible
- [ ] Add recording of movements
