package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fileName := "something.csv"
	data := [][]string{ // 先放入標頭
		{"competitor_id", "competitor_name", "competitor_abbreviation", "competitor_country", "competitor_country_code", "player_id", "player_full_name", "player_name", "player_nationality", "player_country_code", "player_type", "player_date_of_birth"},
	}

	writeAll(fileName, data)
}
func writeAll(fileName string, content [][]string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		file.Close()
	}()

	file.WriteString("\xEF\xBB\xBF") // 寫入UTF-8 BOM，防止中文亂碼

	// 寫資料到csv檔案
	w := csv.NewWriter(file)

	w.WriteAll(content)

	w.Flush() // 把在buffer緩存中的所有資料輸出
}
