package AutoPartsParser

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Запустить программу с выбора режима работы
func StartConsole() {
	log.Println()
	fmt.Print(`Введите режим работы:
-> 1 - Только Avtoto
-> 2 - Только ZZap
-> 3 - Avtoto и ZZap
 > `)
	var Input int
	fmt.Scan(&Input)

	// Выбор режима работы парсера
	var AvtotoIsParse, ZZapIsParse bool
	switch Input {
	case 1:
		AvtotoIsParse = true
	case 2:
		ZZapIsParse = true
	case 3:
		AvtotoIsParse = true
		ZZapIsParse = true
	default:
		fmt.Println("Некорректный ввод. Введите 1, 2 или 3.")
	}

	// Запускаем парсинг
	if AvtotoIsParse || ZZapIsParse {
		ErrorParse := Parse(AvtotoIsParse, ZZapIsParse)
		if ErrorParse != nil {
			log.Println(ErrorParse)
		}
	}

	// "Мягкий" выход из программы
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		exit := scanner.Text()
		if exit == "q" {
			break
		} else {
			fmt.Println("Press 'q' to quit")
		}
	}
}
