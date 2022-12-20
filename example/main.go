package main

import (
	"io"
	"net/http"
	"os"
	progressively "progressively/sdk"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	sdk := progressively.SdkBuilder("valid-sdk-key", "http://localhost:4000").Build()

	if sdk.Evaluate("newHomepage") == true {
		io.WriteString(w, "<p>New homepage</p>")
	} else {
		io.WriteString(w, "<p>Old homepage</p>")
	}

}

func main() {
	http.HandleFunc("/", getRoot)
	err := http.ListenAndServe(":3003", nil)

	if err != nil {
		panic(err)
		os.Exit(1)
	}

}
