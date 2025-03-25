package helpers

import (
	"crypto/sha256"
	"fmt"
	"log"
	"math/rand"
)

type Debug struct {
	ID string
}

func (c *Debug) Init() {

	randomNum := rand.Intn(10000)
	fmt.Println("Random number:", randomNum)

	// Convert the number to a string and hash it
	hash := sha256.Sum256([]byte(fmt.Sprintf("%d", randomNum)))
	log.Println("hash", hash)
}
