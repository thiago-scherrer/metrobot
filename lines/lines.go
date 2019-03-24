package lines

import (
	"regexp"
)

func Blue(body []byte) string {
	bodyResult := string(body)
	reader, _ := regexp.Compile("\"Linha 1-Azul\",\"status\" : \"[a-zA-Z]+\"")
	return reader.FindString(bodyResult)
}

func Green(body []byte) string {
	bodyResult := string(body)
	reader, _ := regexp.Compile("\"Linha 2-Verde\",\"status\" : \"[a-zA-Z]+\"")
	return reader.FindString(bodyResult)
}

func Red(body []byte) string {
	bodyResult := string(body)
	reader, _ := regexp.Compile("\"Linha 3-Vermelha\",\"status\" : \"[a-zA-Z]+\"")
	return reader.FindString(bodyResult)
}

func Lilac(body []byte) string {
	bodyResult := string(body)
	reader, _ := regexp.Compile("\"Linha 5-Lilás\",\"status\" : \"[a-zA-Z]+\"")
	return reader.FindString(bodyResult)
}

func Silver(body []byte) string {
	bodyResult := string(body)
	reader, _ := regexp.Compile("\"Linha 15-Prata\",\"status\" : \"[a-zA-Z]+\"")
	return reader.FindString(bodyResult)
}

func Yellow(body []byte) string {
	bodyResult := string(body)
	reader, _ := regexp.Compile("\"Linha 4-Amarela\",\"status\" : \"[a-zA-Z]+\"")
	return reader.FindString(bodyResult)
}
