import './bookInfoPanel.css';
import { Image, Button } from '@mantine/core';
import { IconTablePlus, IconPencilPlus } from '@tabler/icons-react';
import { Rating, Tooltip } from '@mantine/core';

export default function BookInfoPanel({ book }) {
	// Test data
	const avgRating = 8.3;
	const numReviews = 501;

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
					<div className="book-info-review-aggregate">
						<div className="review-average-rating">
							Average rating: 
								<Tooltip label={avgRating/2}>
								<Rating readOnly fractions={2} value={avgRating/2}/>
								</Tooltip>
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

				<div className="book-info-actions-container">
					<Button rightSection={<IconTablePlus size={20}/>} fullWidth>Add to my Readlist</Button>
					<Button rightSection={<IconPencilPlus size={20}/>} fullWidth>Create Review</Button>
				</div>
			</div>
		</div>
	)
}
