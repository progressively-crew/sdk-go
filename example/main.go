package main

import (
	"fmt"
	progressively "progressively/sdk"
)

func main() {
	sdk := progressively.SdkBuilder("valid-sdk-key", "http://localhost:4000").AddField("e", "lol").AddField("hello", 1).Build()
	fmt.Println("fefw", sdk.Evaluate("newHomepage"))
	fmt.Println("fefw", sdk.Evaluate("newFooter"))

}
