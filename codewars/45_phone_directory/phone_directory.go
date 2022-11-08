package phone_directory

import (
	"fmt"
	"strings"
)

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(1)
func Phone(directory, number string) string {
	var matchingRecord Record

	for _, rawRecord := range strings.Split(directory, "\n") {
		currentRecord := newRecord(rawRecord)

		if matchingRecord.Initialized() &&
			matchingRecord.Phone == currentRecord.Phone {
			return fmt.Sprintf("Error => Too many people: %s", number)
		}

		if currentRecord.Phone == number {
			matchingRecord = currentRecord
		}
	}

	if !matchingRecord.Initialized() {
		return fmt.Sprintf("Error => Not found: %s", number)
	}

	return fmt.Sprintf(
		"Phone => %s, Name => %s, Address => %s",
		matchingRecord.Phone,
		matchingRecord.Name,
		matchingRecord.Address,
	)
}

func newRecord(rawRecord string) Record {
	var phone strings.Builder
	var name strings.Builder
	var address strings.Builder
	entry := "address"

	for _, character := range rawRecord {
		switch true {
		case character == '+':
			entry = "phone"
			continue
		case character == '<':
			entry = "name"
			continue
		case character == '>' || (character == ' ' && entry == "phone"):
			entry = "address"
			continue
		}

		switch entry {
		case "address":
			address.WriteRune(character)
		case "name":
			name.WriteRune(character)
		case "phone":
			phone.WriteRune(character)
		}
	}

	return Record{
		initialized: true,
		Phone:       normalizeRecord(phone.String()),
		Name:        normalizeRecord(name.String()),
		Address:     normalizeRecord(address.String()),
	}
}

func normalizeRecord(value string) string {
	var normalizedValue strings.Builder
	var lastCharacter rune

	for _, character := range value {
		if character == ' ' && lastCharacter == ' ' {
			continue
		}

		if (character >= 'a' && character <= 'z') ||
			(character >= 'A' && character <= 'Z') ||
			(character >= '0' && character <= '9') ||
			character == ' ' ||
			character == '-' ||
			character == '\'' ||
			character == '.' {
			normalizedValue.WriteRune(character)
			lastCharacter = character
		}

		if character == '_' && lastCharacter != ' ' {
			normalizedValue.WriteRune(' ')
			lastCharacter = character
		}
	}

	return strings.Trim(normalizedValue.String(), " ")
}

type Record struct {
	initialized bool
	Phone       string
	Name        string
	Address     string
}

func (record *Record) Initialized() bool {
	return record.initialized
}
