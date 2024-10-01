import './review.css';
import { Rating } from '@mantine/core';

export default function BookReview({ review }) {
	return (
		<div>
			<div className="review-header">{review.Title ?? 'PLACEHOLDER HEADER'}</div>
			<div className="review-container">
				<div className="review-metadata-container">
					<div className="review-user">{review.Username}</div>
					<div className="vertical-seperator">|</div>
					<Rating readOnly fractions={2} value={review.Rating / 2} className="review-value" />
				</div>
				<div className="review-description">{review.Description}</div>
			</div>
		</div>
	)
}