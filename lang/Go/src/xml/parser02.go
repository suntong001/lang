////////////////////////////////////////////////////////////////////////////
// Porgram: parser02
// Purpose: Go xml parsing demo
// Authors: Tong Sun (c) 2015, All rights reserved
// Credits: https://jan.newmarch.name/go/xml/chapter-xml.html
////////////////////////////////////////////////////////////////////////////

/*
Parse XML
*/

package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := `
<person>
  <name>
    <family> Newmarch </family>
    <personal> Jan </personal>
  </name>
  <email type="personal">
    jan@newmarch.name
  </email>
  <email type="work">
    j.newmarch@boxhill.edu.au
  </email>
</person>
`

	/*
		if len(os.Args) != 2 {
			fmt.Println("Usage: ", os.Args[0], "file")
			os.Exit(1)
		}
		file := os.Args[1]
		bytes, err := ioutil.ReadFile(file)
		checkError(err)
		r := strings.NewReader(string(bytes))
	*/

	r := strings.NewReader(input)

	parser := xml.NewDecoder(r)
	depth := 0
	for {
		token, err := parser.Token()
		if err != nil {
			break
		}
		switch t := token.(type) {
		case xml.StartElement:
			elmt := xml.StartElement(t)
			name := elmt.Name.Local
			printElmt(name, depth)
			depth++
		case xml.EndElement:
			depth--
			elmt := xml.EndElement(t)
			name := elmt.Name.Local
			printElmt(name, depth)
		case xml.CharData:
			bytes := xml.CharData(t)
			printElmt("\""+string([]byte(bytes))+"\"", depth)
		case xml.Comment:
			printElmt("Comment", depth)
		case xml.ProcInst:
			printElmt("ProcInst", depth)
		case xml.Directive:
			printElmt("Directive", depth)
		default:
			fmt.Println("Unknown")
		}
	}
}

func printElmt(s string, depth int) {
	for n := 0; n < depth; n++ {
		fmt.Print("  ")
	}
	fmt.Println(s)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
