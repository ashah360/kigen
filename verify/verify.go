package verify

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Verification is a struct of the game validation payload
type Verification struct {
	Hash      [16]byte
	Cipher    string
	GameCheck [16]byte
	Score     int
}

func createHash(data []byte) [16]byte {
	hash := md5.Sum(data)
	return hash
}

func makeGameCheck(gameName string, score int) [16]byte {
	format := fmt.Sprintf("%s%d", gameName, score)
	hash := createHash([]byte(format))
	return hash
}

func makeCipher(score int) string {
	timestamp := time.Now().UnixNano() / 1e6
	format := fmt.Sprintf("%d%d", score, timestamp)
	hash := createHash([]byte(format))

	result := fmt.Sprintf("%x%s", hash, strings.Repeat("0", 68))

	return result
}

// HashString returns the receiver's stringified hash
func (v *Verification) HashString() string {
	return fmt.Sprintf("%x", v.Hash)
}

// GameString returns the receiver's stringified GameCheck
func (v *Verification) GameString() string {
	return fmt.Sprintf("%x", v.GameCheck)
}

// Make populates the receiver fields with generated game validation data
func (v *Verification) Make(score int) {
	scoreString := strconv.Itoa(score)

	gameCheck := makeGameCheck("sorceryStones", score)
	cipher := makeCipher(score)

	seed := scoreString
	ref := scoreString
	index := 0

	for i := (len(scoreString) - 1); i >= 0; i-- {
		i, err := strconv.Atoi(string(ref[i]))

		if err != nil {
			panic(err)
		}

		index = i

		seed += string(cipher[index])
	}

	hash := createHash([]byte(seed))

	v.Hash = hash
	v.Cipher = cipher
	v.GameCheck = gameCheck
	v.Score = score
}
