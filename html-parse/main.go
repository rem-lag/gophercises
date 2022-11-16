package main

import (
	"fmt"
	"rem-lag/htmlparse/links"
	"strings"
)

func main() {

	r := strings.NewReader(exHtml)
	links, err := links.Parse(r)
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
