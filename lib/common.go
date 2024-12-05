package common

import "unicode"

func Filter[T any](slice []T, fn func(T) bool) []T {
	var result []T

	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}

	return result
}

func Map[T any, TOut any](slice []T, fn func(T) TOut) []TOut {
	var result []TOut

	for _, v := range slice {
		result = append(result, fn(v))
	}

	return result
}

func IsWhitespace(str string) bool {
	for _, v := range str {
		if !unicode.IsSpace(v) {
			return false
		}
	}
	return true;
}
