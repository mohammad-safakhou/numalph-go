package numalph_go

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	delimiter = " و "
	zero      = "صفر"
	negative  = "منفی "
	letters   = [][]string{
		{"", "یک", "دو", "سه", "چهار", "پنج", "شش", "هفت", "هشت", "نه"},
		{"ده", "یازده", "دوازده", "سیزده", "چهارده", "پانزده", "شانزده", "هفده", "هجده", "نوزده", "بیست"},
		{"", "", "بیست", "سی", "چهل", "پنجاه", "شصت", "هفتاد", "هشتاد", "نود"},
		{"", "یکصد", "دویست", "سیصد", "چهارصد", "پانصد", "ششصد", "هفتصد", "هشتصد", "نهصد"},
		{"", " هزار", " میلیون", " میلیارد", " بیلیون", " بیلیارد", " تریلیون", " تریلیارد",
			" کوآدریلیون", " کادریلیارد", " کوینتیلیون", " کوانتینیارد", " سکستیلیون", " سکستیلیارد", " سپتیلیون",
			" سپتیلیارد", " اکتیلیون", " اکتیلیارد", " نانیلیون", " نانیلیارد", " دسیلیون", " دسیلیارد"},
	}
	decimalSuffixes = []string{
		"",
		"دهم",
		"صدم",
		"هزارم",
		"ده‌هزارم",
		"صد‌هزارم",
		"میلیونوم",
		"ده‌میلیونوم",
		"صدمیلیونوم",
		"میلیاردم",
		"ده‌میلیاردم",
		"صد‌‌میلیاردم",
	}
)

// prepareNumber pads the input number with zeros to make its length a multiple of 3
// and splits it into an array of 3-digit strings.
func prepareNumber(num string) []string {
	out := num
	length := len(out)
	if length%3 == 1 {
		out = "00" + out
	} else if length%3 == 2 {
		out = "0" + out
	}
	var result []string
	for i := 0; i < len(out); i += 3 {
		result = append(result, out[i:i+3])
	}
	return result
}

// tinyNumToWord converts a 3-digit number into its Persian word representation.
func tinyNumToWord(num string) string {
	parsedInt, err := strconv.Atoi(num)
	if err != nil {
		return ""
	}
	if parsedInt == 0 {
		return ""
	}
	if parsedInt < 10 {
		return letters[0][parsedInt]
	}
	if parsedInt <= 20 {
		return letters[1][parsedInt-10]
	}
	if parsedInt < 100 {
		one := parsedInt % 10
		ten := (parsedInt - one) / 10
		if one > 0 {
			return letters[2][ten] + delimiter + letters[0][one]
		}
		return letters[2][ten]
	}
	one := parsedInt % 10
	hundreds := (parsedInt - (parsedInt % 100)) / 100
	ten := (parsedInt - ((hundreds * 100) + one)) / 10
	out := []string{letters[3][hundreds]}
	secondPart := (ten * 10) + one
	if secondPart == 0 {
		return strings.Join(out, delimiter)
	}
	if secondPart < 10 {
		out = append(out, letters[0][secondPart])
	} else if secondPart <= 20 {
		out = append(out, letters[1][secondPart-10])
	} else {
		out = append(out, letters[2][ten])
		if one > 0 {
			out = append(out, letters[0][one])
		}
	}
	return strings.Join(out, delimiter)
}

// convertDecimalPart converts the decimal part of a number into its Persian word representation.
func convertDecimalPart(decimalPart string) string {
	decimalPart = strings.TrimRight(decimalPart, "0")
	if decimalPart == "" {
		return ""
	}
	if len(decimalPart) > 11 {
		decimalPart = decimalPart[:11]
	}
	length := len(decimalPart)
	suffix := ""
	if length < len(decimalSuffixes) {
		suffix = decimalSuffixes[length]
	}
	return " ممیز " + NumToPersian(decimalPart) + " " + suffix
}

// NumToPersian converts a numerical string into its Persian word representation.
func NumToPersian(input string) string {
	re := regexp.MustCompile(`[^0-9\.-]`)
	input = re.ReplaceAllString(input, "")
	isNegative := false
	floatParse, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return zero
	}
	if floatParse == 0 {
		return zero
	}
	if floatParse < 0 {
		isNegative = true
		input = strings.ReplaceAll(input, "-", "")
	}
	decimalPart := ""
	integerPart := input
	if strings.Contains(input, ".") {
		parts := strings.SplitN(input, ".", 2)
		integerPart = parts[0]
		decimalPart = parts[1]
	}
	if len(integerPart) > 66 {
		return "خارج از محدوده"
	}
	slicedNumber := prepareNumber(integerPart)
	var out []string
	for i, num := range slicedNumber {
		converted := tinyNumToWord(num)
		if converted != "" {
			index := len(slicedNumber) - (i + 1)
			if index < len(letters[4]) {
				converted += letters[4][index]
			}
			out = append(out, converted)
		}
	}
	decimalStr := ""
	if len(decimalPart) > 0 {
		decimalStr = convertDecimalPart(decimalPart)
	}
	negativeStr := ""
	if isNegative {
		negativeStr = negative
	}
	return negativeStr + strings.Join(out, delimiter) + decimalStr
}
