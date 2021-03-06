# Sending

A port of [sent](https://tools.suckless.org/sent/) for generating js-free html slideshows.
[Demo](https://maniks.net/pres-demo.html)
[Demo Source](sendinglib/test/sample.txt)

Usage:
`sending` - currently takes input through stdin and outputs through stdout


**NO EXAMPLE LINKS EXIST YET**
Example input:
```
one

two

three
```

## Goals
 * Rapid development
 * Easy to parse
 * Easy to remember
 * Single file output (including images)

## Directives
Individual slides can be formatted differently:
text (default):
```
.txt
one

.txt
second slide
```

list (default unordered):
```
.li
one 
two
three

second slide
```

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

Image with alt text:
```
.img
path/example.png
alt text
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

**Requires pygmentize in path** (pip install pygments)

Syntax highlighted text:
```
.pre
.go
if (bolb) {
    fmt.Printf("ohno")
}
```

There is also a special '.link' directive for links, this is the only one that works mid-slide.
```
.txt
Some text
.link
link-text (google)
https://google.com
```

Used as a reference:
[Chen Hui Jing - HTML slides without frameworks, just CSS](https://chenhuijing.com/blog/html-slides-without-frameworks/)
