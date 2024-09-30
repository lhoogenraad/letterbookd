import './bookInfoPanel.css';
import { Image } from '@mantine/core';

export default function BookInfoPanel({ book }) {
	// Test data
	const avgRating = 4.6;
	const numReviews = 501;

	if (!book) {
		return (
			<div className="book-container">
				<div>Rendering. . .</div>
			</div>
		)
	}

	return (
		<div className="book-container">
			<div className="book-info-container">

				<div className="book-info-item">
					<span className="book-info-subheader">
						Title
					</span>
					<h2 className="book-title">{book.Title}</h2>
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
					<div className="book-info-review-aggregate">
						<div className="review-average-rating">
							Average rating: {avgRating}
						</div>
						<div className="review-number-views">
							(Out of {numReviews} reviews)
						</div>
					</div>
				</div>
			</div>

			<div className="book-info-img">
				<Image 
					className='book-tile-img'
					src={null}
					w='30vw'
					radius="md"
					fallbackSrc="https://placehold.co/600x400?text=Placeholder"
				/>
			</div>
		</div>
	)
}
