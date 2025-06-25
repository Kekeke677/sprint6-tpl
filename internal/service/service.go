package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func AutoDetectAndConvert(data string) (string, error) {
	if data == "" {
		return "", errors.New("input data is empty")
	}

	trimmedData := strings.TrimSpace(data)

	isMorse := true
	for _, char := range trimmedData {
		if char != '.' && char != '-' && char != ' ' && char != '\n' && char != '\r' {
			isMorse = false
			break
		}
	}

	if !strings.ContainsAny(trimmedData, ".-") {
		isMorse = false
	}

	if isMorse {
		return morse.ToText(trimmedData), nil
	} else {
		return morse.ToMorse(trimmedData), nil
	}
}
