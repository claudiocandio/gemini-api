package main

import (
	"encoding/json"
	"fmt"

	"github.com/claudiocandio/gemini-api"
)

func main() {

	api := gemini.New(
		false, // if this is false, it will use Gemini Sandox site: <https://api.sandbox.gemini.com>
		// if this is true,  it will use Gemini Production site: <https://api.gemini.com>
		"MyGeminiApiKey",    // GEMINI_API_KEY
		"MyGeminiApiSecret", // GEMINI_API_SECRET
	)

	// check more api methods in private.go & public.go
	accountDetail, err := api.AccountDetail()
	if err != nil {
		fmt.Printf("Error AccountDetail: %s\n", err)
		return
	}
	j, err := json.MarshalIndent(&accountDetail, "", " ")
	if err != nil {
		fmt.Printf("Error MarshalIndent: %s\n", err)
	}

	fmt.Printf("%s", j)
}
