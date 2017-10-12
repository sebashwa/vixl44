![Logo](logo.svg)

Create pixel art inside your terminal using vim movements

# Installation
[Get go](https://golang.org/dl/), then:
```
go get github.com/sebashwa/vixl44
```

# Usage
```
~/go/bin/vixl44 [FILENAME] [OPTIONS]

FILENAME

    the name of your file

OPTIONS

-c, --cols
    number of columns, 0 means full width, ignored if filename given (default 20)
-r, --rows
    number of rows, 0 means full height, ignored if filename given (default 20)
-f, --filename
    the name of your file
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
:w [FILENAME]<CR>  - Write to [FILENAME] 
:wq [FILENAME]<CR> - Write to [FILENAME] and quit
:q<CR>             - Quit
```

# Todo
- [ ] Implement export to svg
- [ ] Add a "draw mode"
- [ ] Make copy/cut&paste possible
