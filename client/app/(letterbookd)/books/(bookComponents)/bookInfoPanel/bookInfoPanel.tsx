import './bookInfoPanel.css';
import { Image, Button } from '@mantine/core';
import { IconPencilPlus } from '@tabler/icons-react';
import BookReviewAverage from './bookReviewAverage';
import AddBookToReadlistButton from './addBookToReadlistButton';

export default function BookInfoPanel({ book }) {

	if (!book) {
		return (
			<div className="book-info-container">
				<div>Rendering. . .</div>
			</div>
		)
	}

	return (
		<div className="book-info-container">
			<div>

				<div className="book-info-item">
					<span className="book-info-subheader">
						Title
					</span>
					<h2 className="book-title">{book.Title}</h2>
				</div>


				<div className="book-info-item">
					<span className="book-info-subheader">
						Author
					</span>
					<div className="book-info-item-content">
						{book.Author}
					</div>
				</div>

				<div className="book-info-item">
					<span className="book-info-subheader">
						Date Published
					</span>
					<div className="book-info-item-content">
						{new Date(book.Published).toLocaleDateString()}
					</div>
				</div>

				<div className="book-info-item">
					<span className="book-info-subheader">
						Synopsis
					</span>
					<div className="book-info-item-content">{book.Synopsis}</div>
				</div>

				<div className="book-info-item">
					<BookReviewAverage bookId={book.Id}/>
				</div>
			</div>

			<div className="book-info-img">
				<Image 
					className='book-tile-img'
					src={null}
					h='auto'
					w='auto'
					radius="md"
					fallbackSrc="https://placehold.co/600x400?text=Placeholder"
				/>

				<div className="book-info-actions-container">
					<AddBookToReadlistButton bookId={book.Id} />
					<Button rightSection={<IconPencilPlus size={20}/>} fullWidth>Create Review</Button>
				</div>
			</div>
		</div>
	)
}
