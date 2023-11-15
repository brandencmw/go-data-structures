package utils

import (
	"math/rand"
)

func GetRandomValueOfType(t interface{}) interface{} {
	switch t.(type) {
	case int64:
		return rand.Int63()
	case int32:
		return rand.Int31()
	case int16:
		return int16(rand.Intn(32767))
	case int8, int:
		return int8(rand.Intn(127))
	case uint64:
		return rand.Uint64()
	case uint32:
		return rand.Uint32()
	case uint16:
		return uint16(rand.Intn(65535))
	case uint8, uint:
		return uint16(rand.Intn(255))
	case float64:
		return rand.Float64()
	case float32:
		return rand.Float32()
	case string:
		return generateRandomString()
	case bool:
		return (rand.Int()%2 == 0)
	default:
		return nil
	}
}

func generateRandomString() string {
	const chars string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := []byte{}
	resultLen := rand.Intn(16)
	for i := 0; i < resultLen; i++ {
		result = append([]byte(result), chars[rand.Intn(len(chars))])
	}
	return string(result)
}

func GetZero[T any]() T {
	var zero T
	return zero
}
