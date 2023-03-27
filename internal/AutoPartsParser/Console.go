package AutoPartsParser

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func StartConsole() {
	log.Println()
	fmt.Println(`Введите режим работы:
-> 1 - Только Avtoto
-> 2 - Только ZZap
-> 3 - Avtoto и ZZap`)
	var Input int
	fmt.Scan(&Input)

	// Выбор режима работы парсера
	var AvtotoIsParse, ZZapIsParse bool
	switch Input {
	case 1:
		AvtotoIsParse = true
		break
	case 2:
		ZZapIsParse = true
		break
	case 3:
		AvtotoIsParse = true
		ZZapIsParse = true
		break
	default:
		fmt.Println("Некорректный ввод. Введите 1, 2 или 3.")
		break
	}

	// Запускаем парсинг
	if !AvtotoIsParse || !ZZapIsParse {
		ErrorParse := Parse(AvtotoIsParse, ZZapIsParse)
		if ErrorParse != nil {
			log.Println(ErrorParse)
		}
	}

	// "Мягкий" выход из программы
	fmt.Println("Press 'q' to quit")
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