package utils

/*
 * Dependencies
 */

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

/*
 * Shuffle utility
 */

func Shuffle[T any](array []T) []T {
	a := array

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(array), func(i, j int) { a[i], a[j] = a[j], a[i] })

	return a
}

/*
 * Get unique value utility
 */

func GetUniqueValue() (string, error) {
	b := make([]byte, 5)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%X", b), nil
}

/*
 * JSON stringify
 */

func StringifyJSON(data any) string {
	val, _ := json.MarshalIndent(data, "", "    ")

	return string(val)
}
