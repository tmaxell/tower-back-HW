package main

import (
	"bytes"
	"os"
	"testing"
)

func runUniq(args []string, input string) (string, error) {
	var outBuf bytes.Buffer
	var inBuf bytes.Buffer
	inBuf.WriteString(input)

	// оригинальные
	oldStdin := os.Stdin
	oldStdout := os.Stdout

	// трассировка на буффер
	os.Stdin = &inBuf
	os.Stdout = &outBuf

	defer func() {
		os.Stdin = oldStdin
		os.Stdout = oldStdout
	}()

	os.Args = append([]string{"uniq"}, args...)

	main()

	return outBuf.String(), nil
}
func TestUniqWithoutParams(t *testing.T) {
	input := `I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.
`
	expected := `I love music.

I love music of Kartik.
Thanks.
I love music of Kartik.
`
	output, err := runUniq([]string{}, input)
	if err != nil {
		t.Errorf("Ошибка при запуске программы: %v", err)
	}

	if output != expected {
		t.Errorf("Ожидалось: %q, получено: %q", expected, output)
	}
}
func TestUniqWithCount(t *testing.T) {
	input := `I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.
`
	expected := `3 I love music.
1 
2 I love music of Kartik.
1 Thanks.
2 I love music of Kartik.
`

	output, err := runUniq([]string{"-c"}, input)
	if err != nil {
		t.Errorf("Ошибка при запуске программы: %v", err)
	}

	if output != expected {
		t.Errorf("Ожидалось: %q, получено: %q", expected, output)
	}
}
func TestUniqWithDuplicates(t *testing.T) {
	input := `I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.
`
	expected := `I love music.
I love music of Kartik.
I love music of Kartik.
`

	output, err := runUniq([]string{"-d"}, input)
	if err != nil {
		t.Errorf("Ошибка при запуске программы: %v", err)
	}

	if output != expected {
		t.Errorf("Ожидалось: %q, получено: %q", expected, output)
	}
}
func TestUniqWithUnique(t *testing.T) {
	input := `I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.
`
	expected := `
Thanks.
`

	output, err := runUniq([]string{"-u"}, input)
	if err != nil {
		t.Errorf("Ошибка при запуске программы: %v", err)
	}

	if output != expected {
		t.Errorf("Ожидалось: %q, получено: %q", expected, output)
	}
}
func TestUniqWithIgnoreFields(t *testing.T) {
	input := `We love music.
I love music.
They love music.

I love music of Kartik.
We love music of Kartik.
Thanks.
`
	expected := `We love music.

I love music of Kartik.
Thanks.
`

	output, err := runUniq([]string{"-f", "1"}, input)
	if err != nil {
		t.Errorf("Ошибка при запуске программы: %v", err)
	}

	if output != expected {
		t.Errorf("Ожидалось: %q, получено: %q", expected, output)
	}
}
func TestUniqWithIgnoreChars(t *testing.T) {
	input := `I love music.
A love music.
C love music.

I love music of Kartik.
We love music of Kartik.
Thanks.
`
	expected := `I love music.

I love music of Kartik.
We love music of Kartik.
Thanks.
`

	output, err := runUniq([]string{"-s", "1"}, input)
	if err != nil {
		t.Errorf("Ошибка при запуске программы: %v", err)
	}

	if output != expected {
		t.Errorf("Ожидалось: %q, получено: %q", expected, output)
	}
}
