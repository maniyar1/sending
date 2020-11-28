# Sending

A port of [sent](https://tools.suckless.org/sent/) for generating minimal-js (or js-free) html slideshows.
Usage:
`sending` - currently takes input through stdin and outputs through stdout



**NO EXAMPLE LINKS EXIST YET**
Example input:
```
one

two

three
```

Example output:
[output link](shrug)

Individual slides can be formatted differently:
text (default):
```
.txt
one

.txt
second slide
```
([output](shrug))

list (default unordered):
```
.li
one 
two
three

second slide
```
([output](shrug))

unordered list:
```
.ul
one 
two
three

second slide
```

ordered list:
```
.ol
one
two
three
```

Image:
```
.img
path/example.png
```

Svg:
```
.svg
path/example.svg
```

pre-formatted text:
```
.pre
if (bolb) {
    thing();
} 

.txt
second slide
```

**Requires pygmentize in path**
Syntax highlighted text:
```
.pre
.go
if (bolb) {
    fmt.Printf("ohno")
}
```

([output](shrug))

([output](shrug))

Used as a reference:
[Chen Hui Jing - HTML slides without frameworks, just CSS](https://chenhuijing.com/blog/html-slides-without-frameworks/)
