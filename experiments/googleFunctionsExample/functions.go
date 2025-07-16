package googlefunctionsexample

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("HelloWorld", helloWorld)
	functions.HTTP("HelloAlfonso", helloAlfonso)
}

// helloWorld writes "Hello, World!" to the HTTP response.
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!!!")
}

// helloWorld writes "Hello, World!" to the HTTP response.
func helloAlfonso(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Alfonso!")
}