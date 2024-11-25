import './review.css';
import { Rating } from '@mantine/core';
import { IconMessages } from '@tabler/icons-react';
import Link from 'next/link';
import ReviewHeart from 'components/reviews/likes/reviewHeart';

export default function BookReview({ review, bookId, showTitle}) {
	return (
		<div>
			{
				showTitle ? <div className="review-header">{review.BookTitle ?? 'Review'}</div>
				:
				<div className="review-header"></div>
			}
			<div className="review-container">
				<div className="review-metadata-container">
					<div className="review-user">{review.Username}</div>
					<div className="vertical-seperator">|</div>
					<Rating readOnly fractions={2} value={review.Rating / 2} className="review-value" />
					<div className="vertical-seperator">|</div>
					<Link 
						style={{ textDecoration: "inherit", color: "inherit" }} 
						href={`/books/${bookId}/reviews/${review.Id}`}
					>
						<div className="review-comment-link">
							<IconMessages />
							{review.NumComments}
						</div>
					</Link>
					<div className="vertical-seperator">|</div>
					<ReviewHeart 
						reviewId={review.Id} 
						liked={review.LikedBy} 
						numLikes={review.NumLikes}
					/>
				</div>
				<div className="review-description">{review.Description}</div>
			</div>
		</div>
	)
}
