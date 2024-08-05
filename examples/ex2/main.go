package main

import (
	"fmt"
	link "parser"
	"strings"
)

var exampleHTML = `
<html>
<head>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body>
  <h1>Social stuffs</h1>
  <div>
    <a href="https://www.twitter.com/AditKSingh">
      Check me out on twitter
      <i class="fa fa-twitter" aria-hidden="true"></i>
    </a>
  </div>
</body>
</html>`

func main() {
	r := strings.NewReader(exampleHTML)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)
}
