package models

import (
	"server/internal/resources"
	"server/internal/tools"
)

func GetReadListItems(userId int) ([]resources.ReadListItem, error) {
	var queryString string = `
		SELECT
		read_list_items.id,
		read_list_items.user_id,
		books.id,
		books.name,
		read_list_items.status
		FROM read_list_items
		JOIN books
			ON books.id = read_list_items.book_id;`

	rows, err := tools.DB.Query(queryString)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var readListItems []resources.ReadListItem

	for rows.Next() {
		var item resources.ReadListItem
		if err := rows.Scan(
			&item.Id,
			&item.UserId,
			&item.BookId,
			&item.BookName,
			&item.Status,
		); err != nil {
			return readListItems, err
		}
		readListItems = append(readListItems, item)
	}
	if err = rows.Err(); err != nil {
		return readListItems, err
	}
	return readListItems, nil
}
