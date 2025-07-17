package googlefunctions

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("HelloWorld", helloWorld)
	functions.HTTP("HelloAlfonso", helloAlfonso)
	functions.HTTP("getSingInQr", retrieveLogInQr)
}

// helloWorld writes "Hello, World!" to the HTTP response.
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, Sumar(1, 4))
}

// helloWorld writes "Hello, World!" to the HTTP response.
func helloAlfonso(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Alfonso!")
}

func retrieveLogInQr (w http.ResponseWriter, r *http.Request) {
	getAndRetrieveLogInQr()
}