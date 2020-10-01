package main

import "fmt"

func main() {
	var dia int
	var mes int
	fmt.Println()
	fmt.Scan(&dia)
	fmt.Scan(&mes)

	if mes == 1 && dia >= 21 && dia <= 31 || mes == 2 && dia <= 19 {
		fmt.Println("acuario")
	} else if mes == 2 && dia >= 20 && dia <= 29 || mes == 3 && dia <= 20 {
		fmt.Println("piscis")
	} else if mes == 3 && dia >= 21 && dia <= 31 || mes == 4 && dia <= 20 {
		fmt.Println("aries")
	} else if mes == 4 && dia >= 21 && dia <= 30 || mes == 5 && dia <= 21 {
		fmt.Println("tauro")
	} else if mes == 5 && dia >= 22 && dia <= 31 || mes == 6 && dia <= 21 {
		fmt.Println("geminis")
	} else if mes == 6 && dia >= 22 && dia <= 30 || mes == 7 && dia <= 22 {
		fmt.Println("cancer")
	} else if mes == 7 && dia >= 23 && dia <= 31 || mes == 8 && dia <= 23 {
		fmt.Println("leo")
	} else if mes == 8 && dia >= 24 && dia <= 31 || mes == 9 && dia <= 23 {
		fmt.Println("virgo")
	} else if mes == 9 && dia >= 24 && dia <= 30 || mes == 10 && dia <= 23 {
		fmt.Println("libra")
	} else if mes == 10 && dia >= 24 && dia <= 31 || mes == 11 && dia <= 22 {
		fmt.Println("escorpio")
	} else if mes == 11 && dia >= 23 && dia <= 30 || mes == 12 && dia <= 21 {
		fmt.Println("sagitario")
	} else if mes == 12 && dia >= 22 && dia <= 31 || mes == 1 && dia <= 20 {
		fmt.Println("capricornio")
	}
}
