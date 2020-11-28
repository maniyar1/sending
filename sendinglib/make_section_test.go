package sendinglib

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestFooter(t *testing.T) {
	want := `
        <div class="start">
            <a href="#0"> start </a>
        </div>
        </main>
    </body>
</html>`
	got := Footer
	if got != want {
		t.Errorf("Footer = %q, want %q", got, want)
	}
}

func TestMakeSectionBasic(t *testing.T) {
	id := 1
	text := "one"
	want := "<section id=\"" + strconv.Itoa(id) + "\">\n" + text + "\n<a class=\"next\" href=\"#" + strconv.Itoa(id+1) + "\">next</a>\n<a class=\"prev\" href=\"#" +
		strconv.Itoa(id-1) +
		"\">prev</a>\n</section>\n"
	got := MakeSection(text, id)
	if got != want {
		t.Errorf("makeSection() = %q, want %q", got, want)
	}
}

func TestSplitIntoSlidesList(t *testing.T) {
	input := `
.li
one
two
three`
	reader := strings.NewReader(input)
	got := SplitIntoSlides(reader)
	want := []slide{{SlideType: UNORDERED_LIST, SlideText: "one\ntwo\nthree\n"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestSplitIntoSlidesUnorderedList(t *testing.T) {
	input := `
.ul
one
two
three`
	reader := strings.NewReader(input)
	got := SplitIntoSlides(reader)
	want := []slide{{SlideType: UNORDERED_LIST, SlideText: "one\ntwo\nthree\n"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestSplitIntoSlidesOrderedList(t *testing.T) {
	input := `
.ol
one
two
three`
	reader := strings.NewReader(input)
	got := SplitIntoSlides(reader)
	want := []slide{{SlideType: ORDERED_LIST, SlideText: "one\ntwo\nthree\n"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestSplitIntoSlidesSvg(t *testing.T) {
	input := `
.svg
path/to/svg.svg`
	reader := strings.NewReader(input)
	got := SplitIntoSlides(reader)
	want := []slide{{SlideType: SVG, SlideText: "path/to/svg.svg"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestSplitIntoSlidesImg(t *testing.T) {
	input := `
.img
path/to/img.png`
	reader := strings.NewReader(input)
	got := SplitIntoSlides(reader)
	want := []slide{{SlideType: IMAGE, SlideText: "path/to/img.png"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestSplitIntoSlidesPre(t *testing.T) {
	input := `
.pre
thing`
	reader := strings.NewReader(input)
	got := SplitIntoSlides(reader)
	want := []slide{{SlideType: PRE, SlideText: "thing\n"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestSplitIntoSlides(t *testing.T) {
	input := "one\n\ntwo\n\nthree"
	reader := strings.NewReader(input)
	got := SplitIntoSlides(reader)
	want := []slide{{SlideType: TEXT, SlideText: "one\n"},
		{SlideType: TEXT, SlideText: "two\n"},
		{SlideType: TEXT, SlideText: "three\n"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("splitIntoSlides(reader) = %q, want %q", got, want)
	}
}

func TestSplitIntoSlidesMultiline(t *testing.T) {
	input := "on\ne\n\ntw\no\n\nthre\ne"
	reader := strings.NewReader(input)
	got := SplitIntoSlides(reader)
	want := []slide{{SlideType: TEXT, SlideText: "on\ne\n"},
		{SlideType: TEXT, SlideText: "tw\no\n"},
		{SlideType: TEXT, SlideText: "thre\ne\n"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("splitIntoSlides(reader) = %q, want %q", got, want)
	}
}

func TestlistEntriesFromString(t *testing.T) {
	input := "one\ntwo\nthree\n"
	want := `<li>one</li>
<li>two</li>
<li>three<li>
`
	got := listEntriesFromString(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestUnorderedListFromString(t *testing.T) {
	input := "one\ntwo\nthree\n"
	want := `<ul>
<li>one</li>
<li>two</li>
<li>three</li>
</ul>
`
	got := UnorderedListFromString(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestOrderedListFromString(t *testing.T) {
	input := "one\ntwo\nthree\n"
	want := `<ol>
<li>one</li>
<li>two</li>
<li>three</li>
</ol>
`
	got := OrderedListFromString(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestPreFromString(t *testing.T) {
	input := "text"
	want := "<pre>" + input + "</pre>"
	got := PreFromString(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestLoadSvgFromPath(t *testing.T) {
	input := "test/rect.svg"
	want := `<svg>  <rect width="300" height="100" style="fill:rgb(0,0,255);stroke-width:3;stroke:rgb(0,0,0)" />
</svg>`
	got := LoadSvgFromPath(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}
