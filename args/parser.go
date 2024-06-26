package args

import (
	"fmt"
	"os"
	"regexp"
)

const (
	Shadow     = "shadow"
	Standard   = "standard"
	Thinkertoy = "thinkertoy"
)

// DrawInfo holds the text to be drawn, and with which style it is to be drawn
type DrawInfo struct {
	Text  string
	Style string
}

//Takes the flag '--output=file.txt' together with text and style to be printed
func Parse(args []string) []DrawInfo {

	length_of_arguments := len(args)

	flagAndFile := args[0]
	inspectFlagAndFile(flagAndFile)

	args = args[1:]
	length_of_arguments = (length_of_arguments - 1)

	if length_of_arguments < 1 {

		return nil

	} else if length_of_arguments == 1 {

		text := args[0]
		return []DrawInfo{{Text: Escape(text), Style: Standard}}

	} else {

		// Program received a series of texts to be printed, with banner style specified for consecutive texts
		var finalOutput []DrawInfo

		for textPosition := 0; textPosition < length_of_arguments; textPosition += 2 {
			text := args[textPosition]

			// default style is Standard
			style := Standard

			// check if style is provided
			if textPosition+1 < length_of_arguments {

				// style = args[textPosition]
				switch args[textPosition+1] {

				case Standard, Shadow, Thinkertoy:
					style = args[textPosition+1]

				default:
					_, _ = fmt.Fprintf(os.Stderr, "Style argument not recognized! Passed -> %s Expected -> shadow|standard|thinkertoy\n", args[textPosition+1])
					os.Exit(0)
				}

			}
			finalOutput = append(finalOutput, DrawInfo{Text: Escape(text), Style: style})
		}

		return finalOutput
	}
}


//checks if the flag passed is valid --output=file.txt
func inspectFlagAndFile(flagAndFile string) {

	flagPattern := `^--output=w+.txt{1,255}`
	
	

}