package config

import (
	"strconv"
)

func DefaultValueString(defaultValue, value string) string {
	if value == "" {
		return defaultValue
	}

	return value
}

func DefaultValueIntFromString(defaultValue int, data string) int {
	if data == "" {
		return defaultValue
	}

	dataInt, err := strconv.Atoi(data)
	if err != nil {
		return defaultValue
	}

	return dataInt
}

func DefaultValueBoolFromString(defaultValue bool, data string) bool {
	if data == "" {
		return defaultValue
	}

	dataBool, err := strconv.ParseBool(data)
	if err != nil {
		return defaultValue
	}

	return dataBool
}
