package sendinglib

import (
	"bufio"
    "bytes"
	"io"
    "log"
	"io/ioutil"
	"strconv"
	"strings"
    "os/exec"
)

const (
	UNORDERED_LIST = ".ul"
	ORDERED_LIST   = ".ol"
	LIST           = ".li"
	IMAGE          = ".img"
	SVG            = ".svg"
	TEXT           = ".txt"
	PRE            = ".pre"
    LANG           = ".[lang]"
)

type slide struct {
	SlideType string
	SlideText string
    Language  string // only for pre
}

const Header = `
<!doctype html> 
<html lang="en"> 
    <head> 
        <meta name="viewport" content="width=device-width"> 
        <style> 
        html {
            height: 100%%; 
        }
        body {
            font-size: calc(1em + 10vmin);
            overflow: hidden;
            margin: 0;
            padding: 0;
        }
        section {
          background-color: white;
          height: 100vh;
          width: 100%%;
          position: absolute;
          z-index: 0;
          display: flex;
          align-items: center;
          justify-content: center;
        }

        pre {
            font-family: monospace;
            font-size: calc(0.2em + 2vmin);
        }

        section:target {
          z-index: 1;
        }

        .start {
            background-color: white;
            height: 100vh;
            width: 100%%;
            position: absolute;
        }
        .next {
            font-size: 1em;
            position: absolute;
            right: 0;
            bottom: 0;
        }
        .prev {
            font-size: 1em;
            position: absolute;
            left: 0;
            bottom: 0;
        }
        </style>
        <meta charset="UTF-8">
        <title>Presentation</title>
    <body>
        <main>
        ` // Have to double percent signs to escape them
const Footer = `
        <div class="start">
            <a href="#0"> start </a>
        </div>
        </main>
    </body>
</html>`

func MakeSection(text string, id int) string {
	return "<section id=\"" + strconv.Itoa(id) + "\">\n" + text + "\n<a class=\"next\" href=\"#" + strconv.Itoa(id+1) + "\">next</a>\n<a class=\"prev\" href=\"#" +
		strconv.Itoa(id-1) +
		"\">prev</a>\n</section>\n"
}

func listEntriesFromString(str string) string {
	var result string

	reader := strings.NewReader(str)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		result += "<li>" + scanner.Text() + "</li>\n"
	}
	return result
}

func UnorderedListFromString(str string) string {
	result := "<ul>\n"
	result += listEntriesFromString(str)
	result += "</ul>\n"
	return result
}

func OrderedListFromString(str string) string {
	result := "<ol>\n"
	result += listEntriesFromString(str)
	result += "</ol>\n"
	return result
}

func PreFromString(str string) string {
	return "<pre>" + str + "</pre>"
}

func LoadSvgFromPath(str string) string {
	content, err := ioutil.ReadFile(str)
	if err != nil {
		panic(err)
	}
	result := "<svg>" + string(content) + "</svg>"
	return result
}

func HighlightLanguage(str string, language string) string {
    cmd := exec.Command("pygmentize", "-lgo", "-fhtml", "-Pnoclasses=True")
    cmd.Stdin = strings.NewReader(str)
    var out bytes.Buffer
    cmd.Stdout = &out
    var stderr bytes.Buffer
    cmd.Stderr = &stderr
    err := cmd.Run()
    if err != nil {
        log.Printf("Ran into error %s, defaulting to no syntax highlighting, stderr: %s", err, stderr.String())
        return PreFromString(str)
    } else {
        return out.String()
    }
}

func ParseLanguage(str string) (bool, string) {
    if str[0] == '.' {
        return true, str[1:len(str)]
    } else {
        return false, ""
    }
}

func appendToLastStringInSlice(arr []slide, text string) {
    currentSlide := &arr[len(arr)-1]
    currentSlide.SlideText += text
    if currentSlide.SlideType != IMAGE && currentSlide.SlideType != SVG {
        currentSlide.SlideText += "\n"
    }
}

func SplitIntoSlides(reader io.Reader) []slide {
	var result []slide
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	new := true
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			if new {
				if text == LIST || text == UNORDERED_LIST {
					result = append(result, slide{SlideType: UNORDERED_LIST})
				} else if text == SVG {
					result = append(result, slide{SlideType: SVG})
				} else if text == ORDERED_LIST {
					result = append(result, slide{SlideType: ORDERED_LIST})
				} else if text == IMAGE {
					result = append(result, slide{SlideType: IMAGE})
				} else if text == PRE {
					result = append(result, slide{SlideType: PRE})
                    scanner.Scan() // gotta rescan here ig
                    text = scanner.Text()
                    exist, value := ParseLanguage(text);
                    if (exist) {
                        result[len(result) - 1].Language = value
                        result[len(result) - 1].SlideType = LANG
                    } else {
                        appendToLastStringInSlice(result, text)
                    }
				} else if text == TEXT {
					result = append(result, slide{SlideType: TEXT})
				} else {
					result = append(result, slide{SlideType: TEXT, SlideText: text + "\n"})
				}
				new = false
			} else {
                appendToLastStringInSlice(result, text);
			}
		} else {
			new = true
		}
	}
	return result
}
