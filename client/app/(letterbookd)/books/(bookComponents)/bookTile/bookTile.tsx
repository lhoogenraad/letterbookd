import './bookTile.css';
import { Card, Image } from '@mantine/core';

export default function BookTile({book}) {
	return (
		<Card className="book-tile-container select-cursor">
			<Image 
				className='book-tile-img'
				src={book.CoverURL}
				w='8rem'
				radius="md"
				fallbackSrc="https://placehold.co/600x400?text=Placeholder"
			/>
			<div className="book-tile-title">{book.Title}</div>
			<div className="book-tile-author">{book.Author}</div>
			<div className="book-tile-published">{new Date(book.Published).toLocaleDateString()}</div>
		</Card>
	)
};
