"use client"

import './bookTile.css';
import { Card } from '@mantine/core';

export default function BookTile({book}) {
	return (
		/*<div>book name: {book.Title}</div>*/
		<Card className="book-tile-container">
			<div className="book-tile-img">book img</div>
			<div>{book.Title}</div>
			<div>{book.Author}</div>
			<div>{new Date(book.Published).toLocaleDateString()}</div>
		</Card>
	)
};
