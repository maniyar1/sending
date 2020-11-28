package main

import (
    "maniks.net/sending/sendinglib"
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    slides := sendinglib.SplitIntoSlides(reader)
    fmt.Printf(sendinglib.Header)
    for i, slide := range slides {
        if (slide.SlideType == sendinglib.UNORDERED_LIST) {
            str := sendinglib.UnorderedListFromString(slide.SlideText);
            fmt.Printf(sendinglib.MakeSection(str, i))
        } else if (slide.SlideType == sendinglib.ORDERED_LIST) {
            str := sendinglib.OrderedListFromString(slide.SlideText);
            fmt.Printf(sendinglib.MakeSection(str, i))
        } else if (slide.SlideType == sendinglib.PRE) {
            str := sendinglib.PreFromString(slide.SlideText);
            fmt.Printf(sendinglib.MakeSection(str, i))
        } else if (slide.SlideType == sendinglib.SVG) {
            str := sendinglib.LoadSvgFromPath(slide.SlideText);
            fmt.Printf(sendinglib.MakeSection(str, i))
        } else if (slide.SlideType == sendinglib.IMAGE) {
        } else {
            fmt.Printf(sendinglib.MakeSection(slide.SlideText, i))
        }
    }
    fmt.Printf(sendinglib.Footer)
}
