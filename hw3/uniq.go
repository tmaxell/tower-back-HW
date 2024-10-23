package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
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
		flag.Usage()
		os.Exit(1)
	}
	// тут будем читать файлик
	var input *os.File
	var err error
	if flag.NArg() > 0 {
		input, err = os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка отркытия файла: %v\n", err)
			os.Exit(1)
		}
		defer input.Close()
	} else {
		input = os.Stdin
	}
	var output *os.File
	if flag.NArg() > 1 {
		output, err = os.Create(flag.Arg(1))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка создания файла: %v\n", err)
			os.Exit(1)
		}
		defer output.Close()
	} else {
		output = os.Stdout
	}

	scanner := bufio.NewScanner(input)
	linesCount := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		if *ignoreFlag {
			line = strings.ToLower(line)
		}
		line = processLine(line, *numFieldsFlag, *numCharsFlag)
		if line != "" {
			lineCount[line]++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка чтения: %v\n", err)
		os.Exit(1)
	}
}
