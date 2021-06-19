package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Interpreter interface {
	Handle()
}

type ForPrintInterpreter struct {
	Text  []string
	Index int32
}

func (f *ForPrintInterpreter) Handle() {
	start, _ := strconv.Atoi(f.Text[f.Index])
	f.Index++
	end, _ := strconv.Atoi(f.Text[f.Index])
	for ; start < end; start++ {
		fmt.Println(start)
	}
	bei := &BeginEndInterpreter{
		Text:  f.Text,
		Index: f.Index + 1,
	}
	bei.Handle()
}

type BeginEndInterpreter struct {
	Text  []string
	Index int32
}

func (b *BeginEndInterpreter) Handle() {
	if b.Text[b.Index] == "Begin" {
		fi := &ForPrintInterpreter{
			Text:  b.Text,
			Index: b.Index + 1,
		}
		fi.Handle()
	} else {
		return
	}
}

func main() {
	fmt.Println("老板：设计一门自己的语言，用什么方案好")
	fmt.Println("你：好的")
	text := strings.Split("Begin 1 10 End", " ")
	bei := &BeginEndInterpreter{
		Text:  text,
		Index: 0,
	}
	bei.Handle()
}
