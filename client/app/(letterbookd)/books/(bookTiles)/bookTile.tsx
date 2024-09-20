"use client"

import './bookTile.css';
import { Card } from '@mantine/core';

export default function BookTile({book}) {
	return (
		<Card className="book-tile-container">
			<div className="book-tile-img">book img</div>
			<div className="book-tile-title">{book.Title}</div>
			<div className="book-tile-author">{book.Author}</div>
			<div className="book-tile-published">{new Date(book.Published).toLocaleDateString()}</div>
		</Card>
	)
};
