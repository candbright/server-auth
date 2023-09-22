package uuid

import (
	"github.com/google/uuid"
	"math/rand"
	"strconv"
)

func UUID() string {
	random, err := uuid.NewRandom()
	if err != nil {
		return ""
	}
	return random.String()
}

func RandomRegisterCode() string {
	min := 100000
	max := 1000000
	result := rand.Intn(max-min) + min
	return strconv.Itoa(result)
}
