package main

import (
	"fmt"
	bfconfluence "github.com/kentaro-m/blackfriday-confluence"
	bf "github.com/russross/blackfriday/v2"
	"io/ioutil"
)

func main() {
	renderer := &bfconfluence.Renderer{}
	extensions := bf.CommonExtensions
	md := bf.New(bf.WithRenderer(renderer), bf.WithExtensions(extensions))
	input, err := ioutil.ReadFile("./test_uml.md")
	if err != nil {
		fmt.Printf("read file error %s", err)
		return
	}
	ast := md.Parse(input)
	renderer.Flags = bfconfluence.InformationMacros
	output := renderer.Render(ast) // h1. sample text
	fmt.Printf("%s\n", output)
}
