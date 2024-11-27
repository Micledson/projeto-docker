package validator

import (
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func IsTextBlank(text string) bool {
	text = strings.TrimSpace(text)

	return len(text) <= 0
}

func IsEmailValid(email string) bool {
	email = strings.ToLower(email)

	re := regexp.MustCompile(`^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)

	return re.MatchString(email)
}

func IsCPFValid(cpf string) bool {
	if cpf == "00000000000" || cpf == "11111111111" || cpf == "22222222222" ||
		cpf == "33333333333" || cpf == "44444444444" || cpf == "55555555555" ||
		cpf == "66666666666" || cpf == "77777777777" || cpf == "88888888888" ||
		cpf == "99999999999" || utf8.RuneCountInString(cpf) != 11 {
		return false
	}

	var digits [11]int
	a := []rune(cpf)
	for i := 0; i < 11; i++ {
		digits[i] = int(a[i]) - 48
	}

	var sm, expected10, expected11 int
	p := 10
	for i := 0; i < 9; i++ {
		sm += digits[i] * p
		p--
	}

	r := 11 - (sm % 11)
	if r == 10 || r == 11 {
		expected10 = 0
	} else {
		expected10 = r
	}

	sm = 0
	p = 11
	for i := 0; i < 10; i++ {
		sm += digits[i] * p
		p--
	}

	r = 11 - (sm % 11)
	if r == 10 || r == 11 {
		expected11 = 0
	} else {
		expected11 = r
	}
	if expected10 == digits[9] && expected11 == digits[10] {
		return true
	} else {
		return false
	}
}

func IsCNPJValid(cnpj string) bool {
	b := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	re := regexp.MustCompile(`[^\d]`)
	c := re.ReplaceAllString(cnpj, "")

	if len(c) != 14 {
		return false
	}

	if matched, _ := regexp.MatchString(`^0{14}$`, c); matched {
		return false
	}

	n := 0
	for i := 0; i < 12; i++ {
		num, _ := strconv.Atoi(string(c[i]))
		n += num * b[i+1]
	}
	if d1 := n % 11; c[12] != strconv.Itoa((func() int {
		if d1 < 2 {
			return 0
		}
		return 11 - d1
	})())[0] {
		return false
	}

	n = 0
	for i := 0; i <= 12; i++ {
		num, _ := strconv.Atoi(string(c[i]))
		n += num * b[i]
	}
	if d2 := n % 11; c[13] != strconv.Itoa((func() int {
		if d2 < 2 {
			return 0
		}
		return 11 - d2
	})())[0] {
		return false
	}

	return true
}
