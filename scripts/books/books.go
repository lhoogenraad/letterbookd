package books

import (
	"scripts/structs"
	"scripts/util"
)

func GetBookOpenLibIdMap() ( map[string]int, error) {
	var books = make(map[string]int)
	selectQuery := `SELECT ol_id, id FROM books WHERE ol_id IS NOT NULL;`

	rows, err := util.DB.Query(selectQuery);

	if err != nil {return books, err}

	for rows.Next() {
		var olId string
		var bookId int
		err := rows.Scan(&olId, &bookId)
		if err != nil {return books, err}
		
		books[olId] = bookId
	}

	return books, nil
}
