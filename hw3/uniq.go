package main

import (
	"flag"
	"fmt"
)

func main() {
	countFlag := flag.Bool("c", false, "Подсчет количество появлений строк")
	doubleFlag := flag.Bool("d", false, "Вывод только дублирующихся строк")
	uniqueFlag := flag.Bool("u", false, "Вывод только уникальных строк")
	ignoreFlag := flag.Bool("i", false, "Игнорирование регистр букв")
	numFieldsFlag := flag.Int("f", 0, "Не учитывать первые num_fields полей")
	numCharsFlag := flag.Int("s", 0, "Не учитывать первые num_chars сиволов")

	flag.Parse()

	if *countFlag && (*doubleFlag || *uniqueFlag) || *doubleFlag && *uniqueFlag {
		fmt.Println("Ошибка: нельзя одновременно использовать флаги -c, -d, -u")
	}
}
