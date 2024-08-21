package flags

import (
	"flag"
	"strings"
)

func FlagHandler() (int, []string) {
	// Объявляем флаги
	configPtr := flag.Bool("config", false, "Конфигурация")
	parsePtr := flag.String("parse", "", "Ссылки для парсинга")
	setPtr := flag.String("set", "", "параметры")

	// Парсим аргументы командной строки
	flag.Parse()

	if *configPtr {
		if *setPtr != "" {
			//config + set
			return 210, strings.Split(*setPtr, ":")
		}
		//config
		return 200, nil
	} else if *parsePtr != "" {
		//parse
		return 100, strings.Split(*parsePtr, ",")
	} else {
		// Если не был использован ни один флаг, выводим Usage
		return 0, nil
	}
}
