package main

import "github.com/restuhaqza/tdd-in-go/integrations"

func main() {
	r := integrations.HTTPServer()

	// listen and serve on 3000 but showing the log if already listener
	r.Run(":3000")
}
