"use client";

import './review.css';
import { Rating } from '@mantine/core';
import { IconMessages } from '@tabler/icons-react';
import Link from 'next/link';
import ReviewHeart from 'components/reviews/likes/reviewHeart';
import { useState } from 'react';

export default function BookReview({ review, bookId }) {
	const [likes, setLikes] = useState(review.NumLikes);
	const updateLikedCount = (status: boolean) => {
		console.log(status)
		if (status === true) {
			setLikes(likes+1)
		} else {
			setLikes(likes-1)
		}
	}
	return (
		<div>
			<div className="review-header">{review.Title ?? 'PLACEHOLDER HEADER'}</div>
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
						updateLikedCount={updateLikedCount}
					/>
					{likes}
				</div>
				<div className="review-description">{review.Description}</div>
			</div>
		</div>
	)
}
