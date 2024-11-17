import './bookTile.css';
import { Card, Image } from '@mantine/core';

export default function BookTile({book}) {
	return (
		<Card className="book-tile-container select-cursor">
		<div>
			<Image 
				className='book-tile-img'
				src={book.CoverURL}
				w='10rem'
				radius="md"
				fallbackSrc="https://placehold.co/600x400?text=Placeholder"
			/>
		</div>
		<div>
			<div className="book-tile-title">{book.Title}</div>
			<div className="book-tile-author">{book.Author}</div>
			<div className="book-tile-published">{new Date(book.Published).toLocaleDateString()}</div>
		</div>
		</Card>
	)
};
