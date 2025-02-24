package main

import (
	"encoding/hex"
	"fmt"
	"os"
)

func decodeFromHex(hexString string) ([]rune, error) {
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return []rune{}, err
	}

	return []rune(string(bytes)), nil
}

func decodeRunes(data, key []rune, kelp int) []rune {
	result := make([]rune, len(data))
	for i := 0; i < len(data); i++ {
		result[i] = (data[i] - rune(kelp)) ^ key[i]
	}
	return result
}

func decodeString(encodedString string) (string, error) {
	decodedBytes, err := decodeFromHex(encodedString)
	if err != nil {
		return "", fmt.Errorf("error decoding hex string: %v", err)
	}

	totalLength := len(decodedBytes)

	kelpIndex := totalLength / 2
	kelp := int(decodedBytes[kelpIndex])

	encodedData := decodedBytes[:kelpIndex]
	encodedKey := decodedBytes[kelpIndex+1:]

	decodedKey := decodeRunes(encodedKey, encodedData, kelp)
	decodedData := decodeRunes(encodedData, decodedKey, kelp)

	decodedString := string(decodedData)

	return decodedString, nil
}

func main() {
	data, err := os.ReadFile("./kelpfile")
	if err != nil {
		fmt.Println(err)
		return
	}

	str, err := decodeString(string(data))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(str)
}

