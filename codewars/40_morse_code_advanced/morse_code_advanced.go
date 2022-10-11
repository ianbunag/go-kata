package morse_code_advanced

import (
	"strings"
)

var MORSE_CODE = map[string]string{
	"-...":      "B",
	"-.-":       "K",
	".-..":      "L",
	".--":       "W",
	"....-":     "4",
	"---...":    ":",
	"-.-.":      "C",
	".---":      "J",
	"--..":      "Z",
	".-.-.":     "+",
	"--.":       "G",
	".----.":    "'",
	"-.-.--":    "!",
	"-.--.-":    ")",
	"-.-.-.":    ";",
	"...":       "S",
	".----":     "1",
	"..---":     "2",
	"-..-.":     "/",
	"-.--.":     "(",
	".-...":     "&",
	"-..-":      "X",
	".-.-.-":    ".",
	"...-..-":   "$",
	"---":       "O",
	"--.-":      "Q",
	"-.--":      "Y",
	"---..":     "8",
	"----.":     "9",
	"..--..":    "?",
	"..--.-":    "_",
	".-":        "A",
	"--":        "M",
	"-----":     "0",
	"--...":     "7",
	"..-.":      "F",
	".-.":       "R",
	"--..--":    ",",
	".-..-.":    "\"",
	"-..":       "D",
	"-....-":    "-",
	"...-":      "V",
	".":         "E",
	"..":        "I",
	"-.":        "N",
	"-":         "T",
	"...--":     "3",
	"...---...": "SOS",
	"-....":     "6",
	".--.-.":    "@",
	".--.":      "P",
	"..-":       "U",
	"....":      "H",
	".....":     "5",
	"-...-":     "=",
}

const (
	INITIAL = '0' - 1
	DOWN    = '0'
	UP      = '1'
)

// Average time complexity: O(n log n)
// Worst time complexity:   O(n^2)
// Space complexity:        O(log n)
func DecodeBits(bits string) string {
	currentSequence := INITIAL
	signalCounts := []int{0}
	baseSignal := 0

	for _, bit := range bits {
		if currentSequence == INITIAL {
			if bit == DOWN {
				continue
			}

			currentSequence = UP
		}

		signalCountsLen := len(signalCounts)

		if bit == currentSequence {
			signalCounts[signalCountsLen-1] += 1
		}

		if bit != currentSequence {
			lastCount := signalCounts[signalCountsLen-1]

			if baseSignal == 0 || baseSignal > lastCount {
				baseSignal = lastCount
			}

			signalCounts = append(signalCounts, 1)
			currentSequence = bit
		}
	}

	if baseSignal == 0 {
		baseSignal = signalCounts[0]
	}

	var builder strings.Builder
	bitmap := createBitmap(baseSignal)

	for index, count := range signalCounts {
		if index%2 == 0 {
			builder.WriteString(bitmap[UP][count])
			continue
		}

		builder.WriteString(bitmap[DOWN][count])
	}

	return builder.String()
}

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(1)
func DecodeMorse(morseCode string) string {
	var decoder MorseDecoder

	for _, bit := range morseCode {
		decoder.Receive(bit)
	}

	return decoder.Commit().String()
}

func createBitmap(baseline int) map[rune]map[int]string {
	return map[rune]map[int]string{
		UP: {
			baseline:     "1",
			baseline * 3: "111",
		},
		DOWN: {
			baseline:     "0",
			baseline * 3: "000",
			baseline * 7: "0000000",
		},
	}
}

type MorseDecoder struct {
	up       int
	down     int
	sentence strings.Builder
	word     strings.Builder
	signal   strings.Builder
}

func (decoder *MorseDecoder) Receive(signal rune) {
	if signal == DOWN {
		if decoder.up > 0 {
			decoder.commitSignal()
		}

		decoder.down += 1
	}

	if signal == UP {
		if decoder.down > 0 {
			decoder.commitSpace()
		}

		decoder.up += 1
	}
}

func (decoder *MorseDecoder) commitSignal() {
	if decoder.up == 1 {
		decoder.signal.WriteString(".")
	}

	if decoder.up == 3 {
		decoder.signal.WriteString("-")
	}

	decoder.up = 0
}

func (decoder *MorseDecoder) commitSpace() {
	if decoder.down == 3 {
		decoder.commitLetter()
	}

	if decoder.down == 7 {
		decoder.commitWord()
	}

	decoder.down = 0
}

func (decoder *MorseDecoder) commitLetter() {
	decoder.commitSignal()
	decoder.word.WriteString(MORSE_CODE[decoder.signal.String()])
	decoder.signal.Reset()
}

func (decoder *MorseDecoder) commitWord() {
	decoder.commitLetter()

	if decoder.sentence.Len() > 0 {
		decoder.sentence.WriteString(" ")
	}

	decoder.sentence.WriteString(decoder.word.String())
	decoder.word.Reset()
}

func (decoder *MorseDecoder) Commit() *MorseDecoder {
	decoder.commitWord()

	return decoder
}

func (decoder *MorseDecoder) String() string {
	return decoder.sentence.String()
}
