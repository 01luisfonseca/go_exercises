package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	weekday, _ := readInput("Ingrese día de la semana")
	hour, _ := readInput("Ingrese Hora")
	minute, _ := readInput("Ingrese Minuto")
	description, _ := readInput("Ingrese descripcion")

	if !checkWeekDay(weekday) {
		fmt.Println("No es un día de la semana. ", weekday)
		return
	}

	if !checkWeekDay(weekday) {
		fmt.Println("No es un día de la semana. ", weekday)
		return
	}

	fmt.Println(weekday)
	fmt.Println(hour)
	fmt.Println(minute)
	fmt.Println(description)
}

func readInput(message string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message + ": ")
	text, err := reader.ReadString('\n')
	text = strings.Replace(text, "\r", "", -1)
	text = strings.Replace(text, "\n", "", -1)
	return text, err
}

func checkWeekDay(weekDay string) bool {
	newWeekDay := strings.ToLower(weekDay)
	switch newWeekDay {
	case
		"lunes", "martes", "miercoles", "jueves", "viernes", "sabado", "domingo":
		return true
	}
	return false
}

func checkValidDateInfo(hour, minute string) bool {
	intHour, err := strconv.ParseInt(hour, 10, 0)
	intMin, err2 := strconv.ParseInt(minute, 10, 0)
	if err != nil || err2 != nil {
		return false
	}
	return IntBetween(intHour, 0, 24) && IntBetween(intMin, 0, 59)
}

func IntBetween(i, min, max int) bool {
	if (i >= min) && (i <= max) {
		return true
	} else {
		return false
	}
}
