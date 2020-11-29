package main

import (
	"bufio"
	"fmt"
	"maniks.net/sending/sendinglib"
	"os"
)

func init() {

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	slides := sendinglib.SplitIntoSlides(reader)
	fmt.Println(sendinglib.Header)
	for i, slide := range slides {
		if slide.SlideType == sendinglib.UNORDERED_LIST {
			str := sendinglib.UnorderedListFromString(slide.SlideText)
			fmt.Println(sendinglib.MakeSection(str, i))
		} else if slide.SlideType == sendinglib.ORDERED_LIST {
			str := sendinglib.OrderedListFromString(slide.SlideText)
			fmt.Println(sendinglib.MakeSection(str, i))
		} else if slide.SlideType == sendinglib.PRE {
			str := sendinglib.PreFromString(slide.SlideText)
			fmt.Println(sendinglib.MakeSection(str, i))
		} else if slide.SlideType == sendinglib.LANG {
			str := sendinglib.HighlightLanguage(slide.SlideText, slide.Language)
			fmt.Println(sendinglib.MakeSection(str, i))
		} else if slide.SlideType == sendinglib.SVG {
			str := sendinglib.LoadSvgFromPath(slide.SlideText)
			fmt.Println(sendinglib.MakeSection(str, i))
		} else if slide.SlideType == sendinglib.IMAGE {
			fmt.Println(sendinglib.MakeSection(slide.SlideText, i)) // Image is a special case that is handled in the library because of alt-text
		} else {
			str := sendinglib.ParagraphFromString(slide.SlideText)
			fmt.Println(sendinglib.MakeSection(str, i))
		}
	}
	fmt.Println(sendinglib.Footer)
}
