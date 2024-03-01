package transpose

import (
	"strings"
)

// Transpose транспонирует входные строки.
func Transpose(rows []string) []string {
	longest := 0
	// Находим максимальную длину строки среди всех входных строк.
	for _, r := range rows {
		if len(r) > longest {
			longest = len(r)
		}
	}

	// Создаем матрицу рун для хранения транспонированного результата.
	runeRes := make([][]rune, longest)
	for i := range runeRes {
		runeRes[i] = make([]rune, len(rows))
	}

	// Дополняем каждую строку нулевыми символами до максимальной длины и транспонируем.
	for i, r := range rows {
		rows[i] += strings.Repeat("\x00", longest-len(r))
		// Транспонируем матрицу рун.
		for j := 0; j < longest; j++ {
			runeRes[j][i] = rune(rows[i][j])
		}
	}

	res := make([]string, len(runeRes))

	// Преобразуем транспонированные строки обратно в строки, в обратном порядке (!)
	// заменяя нулевые символы пробелами и удаляя лишние пробелы справа.
	for i, r := range runeRes {
		pad := false
		for j := len(r) - 1; j >= 0; j-- {
			// Если мы встречаем ненулевой символ, начинаем дополнение пробелами.
			if !pad && r[j] != '\x00' {
				pad = true
			} else if pad && r[j] == '\x00' {
				runeRes[i][j] = ' '
			}
		}
		// Удаляем все нулевые символы с правой стороны строки (которые мы изначально пропустили когда определяли pad).
		res[i] = strings.TrimRightFunc(string(r), func(r rune) bool {
			return r == '\x00'
		})
	}

	return res
}
