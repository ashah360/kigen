package generator

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/ashah360/kigen/verify"
)

// Generate returns a generated code from sorcery stones with a given score
func Generate(score int) string {
	v := verify.Verification{}
	v.Make(score)
	URI := "https://www.freekigames.com/minigame/wizard/checkhighscore"

	resp, err := http.PostForm(URI, url.Values{
		"hash":          {v.HashString()},
		"cipher":        {v.Cipher},
		"gameCheck":     {v.GameString()},
		"score":         {strconv.Itoa(v.Score)},
		"gameName":      {"sorceryStones"},
		"versionNumber": {"2"},
	})

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	q, err := url.ParseQuery(string(body))

	if err != nil {
		panic(err)
	}

	if q["error"][0] != "0" {
		panic("Invalid game")
	}

	return q["code"][0]
}
