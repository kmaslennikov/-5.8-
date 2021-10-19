package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Загадайте число от 0 до 10")
	fmt.Println("Программа попробует угадать загаданное вами число.")
	fmt.Println("Если это ваше число, введите 'Q'")
	fmt.Println("Если ваше число больше, введите 'L'")
	fmt.Println("Если ваше число меньше, введите 'S'")
	fmt.Println("Для начала игры введите 'start'")
	var menu string

	fmt.Scan(&menu)
	if menu == "start" {
		digit := 5
		fmt.Println("Ваше число:", digit, "?")
		fmt.Scan(&menu)
		if menu == "Q" {
			fmt.Println("Урра!! Ваше число:", digit)
		} else if menu == "L" {
			digit := 7
			fmt.Println("Ваше число:", digit, "?")
			fmt.Scan(&menu)
			if menu == "Q" {
				fmt.Println("Урра!! Ваше число:", digit)
			} else if menu == "L" {
				digit := 9
				fmt.Println("Ваше число:", digit, "?")
				fmt.Scan(&menu)
				if menu == "Q" {
					fmt.Println("Урра!! Ваше число:", digit)
				} else if menu == "L" {
					digit := 10
					fmt.Println("Ваше число:", digit)
				} else if menu == "S" {
					digit := 8
					fmt.Println("Ваше число:", digit)
				} else {
					fmt.Println("Вы ввели неизвестный симвл")
					log.Fatal()
				}
			} else if menu == "S" {
				digit := 6
				fmt.Println("Ваше число:", digit)
			} else {
				fmt.Println("Вы ввели неизвестный симвл")
				log.Fatal()
			}
		} else if menu == "S" {
			digit := 3
			fmt.Println("Ваше число:", digit, "?")
			fmt.Scan(&menu)
			if menu == "Q" {
				fmt.Println("Урра!! Ваше число:", digit)
			} else if menu == "L" {
				digit := 4
				fmt.Println("Ваше число:", digit)
			} else if menu == "S" {
				digit := 1
				fmt.Println("Ваше число:", digit, "?")
				fmt.Scan(&menu)
				if menu == "Q" {
					fmt.Println("Урра!! Ваше число:", digit)
				} else if menu == "L" {
					digit := 2
					fmt.Println("Ваше число:", digit)
				} else if menu == "S" {
					digit := 0
					fmt.Println("Ваше число:", digit)
				} else {
					fmt.Println("Вы ввели неизвестный симвл")
					log.Fatal()
				}
			} else {
				fmt.Println("Вы ввели неизвестный симвл")
				log.Fatal()
			}
		} else {
			fmt.Println("Вы ввели неизвестный симвл")
			log.Fatal()
		}

	}
}
