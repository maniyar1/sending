package sendinglib

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
    "encoding/base64"
    "io/ioutil"
    "fmt"
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
test/example.png
alt text`
    data, format := LoadBase64FromPath("test/example.png")
    alt := "alt text"
	reader := strings.NewReader(input)
	got := SplitIntoSlides(reader)
	want := []slide{{SlideType: IMAGE, SlideText: AddImgTags(data, format, alt)}}
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

func TestSplitIntoSlidesLanguage(t *testing.T) {
	input := `
.pre
.rust
thing`
	reader := strings.NewReader(input)
	got := SplitIntoSlides(reader)
    want := []slide{{SlideType: LANG, SlideText: "thing\n", Language: "rust"}}
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

func TestCreateLink(t *testing.T) {
    linkText := "link text";
    linkRef := "https://google.com"
    want := "<a href=\"https://google.com\">link text</a>"
    got := createLink(linkText, linkRef);
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestSplitIntoSlidesLink(t *testing.T) {
	input :=`
.link
link text
https://google.com`
	reader := strings.NewReader(input)
	got := SplitIntoSlides(reader)
    want := []slide{{SlideType: TEXT, SlideText: "<a href=\"https://google.com\">link text</a>\n"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("splitIntoSlides(reader) = %q, want %q", got, want)
	}
}

func TestSplitIntoSlidesTextWithLink(t *testing.T) {
	input :=`
text
.link
link text
https://google.com`
	reader := strings.NewReader(input)
	got := SplitIntoSlides(reader)
    want := []slide{{SlideType: TEXT, SlideText: "text\n<a href=\"https://google.com\">link text</a>\n"}}
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

func TestParagraphFromString(t *testing.T) {
	input := "text"
	want := "<p>" + input + "</p>"
	got := PreFromString(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestHighlightLanguage(t *testing.T) { // just make sures it can run...
    input :=`if (bolb) {
    fmt.Printf("ohno")
}`
    lang := "go"
    HighlightLanguage(input, lang);
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

func TestLoadBase64FromPathPng(t *testing.T) {
	input := "test/example.png"
	content, _ := ioutil.ReadFile(input)
    want := base64.StdEncoding.EncodeToString(content);
	data, format := LoadBase64FromPath(input)
	if !reflect.DeepEqual(data, want) && format != "png" {
		t.Errorf("got %q\n\n\n want %q", data, want)
	}
}

func TestLoadBase64FromPathJpg(t *testing.T) {
	input := "test/example.jpg"
	content, _ := ioutil.ReadFile(input)
    want := base64.StdEncoding.EncodeToString(content);
	data, format := LoadBase64FromPath(input)
	if !reflect.DeepEqual(data, want) && format != "jpg" {
		t.Errorf("got %q\n\n\n want %q", data, want)
	}

	input = "test/example.jpeg"
	data, format = LoadBase64FromPath(input)
	if !reflect.DeepEqual(data, want) && format != "jpg" {
		t.Errorf("got %q\n\n\n want %q", data, want)
	}
}

func TestAddImgTags(t *testing.T) {
    input := "test/example.jpeg"
    data, format := LoadBase64FromPath(input)
    alt := "example"
    want := fmt.Sprintf("<img src=\"data:image/%s;base64, %s\" alt=\"%s\"/>", format, data, alt);
    got := AddImgTags(data, format, alt)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestParseLanguage(t *testing.T) {
    input := ".rust"
    want := "rust"
    exists, got := ParseLanguage(input)
	if !reflect.DeepEqual(got, want) && exists {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestParseLanguageNotLanguage(t *testing.T) {
    input := "yes it is me"
    want := false
    got, _ := ParseLanguage(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %t, want %t", got, want)
	}
}
