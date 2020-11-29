package sendinglib

import (
	"bufio"
    "bytes"
	"io"
    "log"
    "fmt"
	"io/ioutil"
	"strconv"
	"strings"
    "os/exec"
    "encoding/base64"
)

const (
	UNORDERED_LIST = ".ul"
	ORDERED_LIST   = ".ol"
	LIST           = ".li"
	IMAGE          = ".img"
	SVG            = ".svg"
	TEXT           = ".txt"
	PRE            = ".pre"
    LINK           = ".link"
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
        .next {
            visibility: hidden;
        }
        .prev {
            visibility: hidden;
        }
        @media (min-width: 600px) {
        * {
            font-family: sans-serif;
            margin: 0;
            padding: 0;
            border: 0;
        }
        html {
            height: 100%; 
        }
        body {
            font-size: calc(1em + 10vmin);
            overflow: hidden;
        }

        p {
            margin: 0.5em;
        }

        section {
          background-color: white;
          height: 100vh;
          width: 100%;
          position: absolute;
          z-index: 0;
          display: flex;
          align-items: center;
          justify-content: center;
        }
        
        ul,ol {
            font-size: 0.75em;
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
            width: 100%;
            position: absolute;
        }

        .next {
            visibility: visible;
            font-size: 1em;
            position: absolute;
            right: 0;
            bottom: 0;
        }
        .prev {
            visibility: visible;
            font-size: 1em;
            position: absolute;
            left: 0;
            bottom: 0;
        }
    }
    @media print {
        .next {
            visibility: hidden;
        }
        .prev {
            visibility: hidden;
        }

        section {
            position: initial;
            z-index: intial;
        }

        section:target {
          z-index: initial;
        }
    }
        </style>
        <script>
        window.addEventListener("keydown", function (event) {
          if (event.defaultPrevented) {
            return; // Do nothing if the event was already processed
          }

          switch (event.key) {
            case "Down": 
            case "ArrowDown":
              prev();
              break;
            case "Up": 
            case "ArrowUp":
              next();
              break;
            case "Left": 
            case "ArrowLeft":
              prev();
              break;
            case "Right": // IE/Edge specific value
            case "ArrowRight":
              next();
              break;
            default:
              return; // Quit when this doesn't handle the key event.
          }

          // Cancel the default action to avoid it being handled twice
          event.preventDefault();
        }, true);

        function next() {
            if(window.location.hash) {
                  let hash = window.location.hash.substring(1); //Puts hash in variable, and removes the # character
                  window.location = "#" + (parseInt(hash) + 1);
                  // hash found
            } else {
                  window.location = "#" + 0;
            }
        }

        function prev() {
            if(window.location.hash) {
                  let hash = window.location.hash.substring(1); //Puts hash in variable, and removes the # character
                  window.location = "#" + (parseInt(hash) - 1);
                  // hash found
            } else {
                  window.location = "#" + 0;
            }
        }
        </script>
        <meta charset="UTF-8">
        <title>Presentation</title>
    <body>
        <main>
        ` // Have to double percent signs to escape them
const Footer = `
        <div class="start">
            <a href="#0" class="next"> start </a>
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

func ParagraphFromString(str string) string {
	return "<p>" + str + "</p>"
}

func QuoteFromString(str string) string {
	return "<blockquote>" + str + "</blockquote>"
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

func LoadBase64FromPath(str string) (string, string) {
	content, err := ioutil.ReadFile(str)
	if err != nil {
		panic(err)
	}
    result := base64.StdEncoding.EncodeToString(content)
    ending := strings.ToLower(str[len(str) - 3:len(str)])
    if ending == "png" {
        return result, "png"
    } else if ending == "jpg" || ending == "peg" { // last three letters so jpeg = peg lol
        return result, "jpg"
    } else {
        log.Printf("File format unknown (is this a jpg or png?), assuming png")
        return result, "png"
    }
}

func AddImgTags(data string, format string, alt string) string {
    return fmt.Sprintf("<img src=\"data:image/%s;base64, %s\" alt=\"%s\"/>", format, data, alt)
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

func createLink(linkText string, linkRef string) string {
    return fmt.Sprintf("<a href=\"%s\">%s</a>", linkRef, linkText)
}

func getLinkFromScanner(scanner *bufio.Scanner) string {
        scanner.Scan() // gotta rescan here ig
        linkText := scanner.Text() // next line has to be file path
        scanner.Scan()
        linkRef := scanner.Text()
        return createLink(linkText, linkRef); // blank alt-text default
}

func appendToLastStringInSlice(arr []slide, text string, scanner *bufio.Scanner) {
    currentSlide := &arr[len(arr)-1]
    if text == LINK {
        arr[len(arr) - 1].SlideText += getLinkFromScanner(scanner)
    } else {
        currentSlide.SlideText += text
    }
    if currentSlide.SlideType != IMAGE && currentSlide.SlideType != SVG && currentSlide.SlideType != LINK {
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
                    scanner.Scan() // gotta rescan here ig
                    filePath := scanner.Text() // next line has to be file path
                    data, format := LoadBase64FromPath(filePath)
                    scanner.Scan()
                    alt := scanner.Text() // next line is either blank, or the alt text, either is fine
                    result[len(result) - 1].SlideText = AddImgTags(data, format, alt); // blank alt-text default
                    continue // continue because sort of messed with the loop, restart for next label
				} else if text == PRE {
					result = append(result, slide{SlideType: PRE})
                    scanner.Scan() // gotta rescan here ig
                    text = scanner.Text()
                    exist, value := ParseLanguage(text);
                    if (exist) {
                        result[len(result) - 1].Language = value
                        result[len(result) - 1].SlideType = LANG
                    } else {
                        appendToLastStringInSlice(result, text, scanner)
                    }
				} else if text == TEXT {
					result = append(result, slide{SlideType: TEXT})
				} else if text == LINK {
                    result = append(result, slide{SlideType: TEXT, SlideText: getLinkFromScanner(scanner) + "\n"})
				} else {
					result = append(result, slide{SlideType: TEXT, SlideText: text + "\n"})
				}
				new = false
			} else {
                appendToLastStringInSlice(result, text, scanner);
			}
		} else {
			new = true
		}
	}
	return result
}
