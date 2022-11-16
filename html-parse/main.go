package main

import (
	"fmt"
	"rem-lag/htmlparse/hparse"
	"strings"
)

func main() {

	r := strings.NewReader(exHtml)
	links, err := hparse.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}

var exHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>
`
