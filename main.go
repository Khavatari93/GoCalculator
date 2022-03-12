package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// Eingabe berechnen
func calculate(x float32, y string, z float32) {
	switch y {
	case "+":
		x += z
	case "-":
		x -= z
	case "*":
		x *= z
	case "/":
		x /= z
	default:
		fmt.Println("Da ist was falsch gelaufen...")
		return
	}

	fmt.Println("Dein Ergebnis ist:", x)
}

// Kontrolle ob int oder string sowie konvertierung von string zu int
func checkNumber(x string) float32 {
	number, err := strconv.ParseFloat(x, 32)
	for err != nil {
		var intCheck = regexp.MustCompile(`[-+]?[0-9]*\.?[0-9]*`)
		for !intCheck.MatchString(x) {
			fmt.Println("Das war keine Nummer, bitte nur Nummern eingeben:")
			fmt.Scanf("%s\n", &x)
		}
		err = nil
	}
	finalNumber := float32(number)
	return finalNumber
}

func main() {
	var number string
	var firstNumber float32
	var operand string
	var secondNumber float32
	var checkOperator = regexp.MustCompile(`\+|\-|\*|\/`)

	// Eingabe vom Benutzer
	fmt.Println("Erste Zahl (Kommastellen als . angeben):")
	fmt.Scanf("%s\n", &number)
	firstNumber = checkNumber(number)

	fmt.Println("Operator (+,-,*,/):")
	fmt.Scanf("%s\n", &operand)

	// Kontrollieren ob ein richtiger Operator benutzt wird
	for !checkOperator.MatchString(operand) {
		fmt.Println("Nur +, -, *, / erlaubt:")
		fmt.Scanf("%s\n", &operand)
	}

	fmt.Println("Zweite Zahl:")
	fmt.Scanf("%s\n", &number)
	secondNumber = checkNumber(number)

	// Ausgabe nach berechnung
	calculate(firstNumber, operand, secondNumber)
}
