"use client";

import { Rating, Tooltip } from '@mantine/core';
import { useState, useEffect } from 'react';
import api from 'util/api/api';
import notify from 'util/notify/notify';

export default function BookReviewAverage({ bookId }) {
	const [loading, setLoading] = useState(false);
	const [avgRating, setAvgRating] = useState(0);
	const [numReviews, setNumReviews] = useState(0);

	const loadBookReviewSummary = (bookId: number) => {
		setLoading(true);
		api.reviews.getBookAverageAndNumberReviews(bookId)
			.then((res) => {
				const { avgRating, numReviews } = res.data;
				setAvgRating(avgRating);
				setNumReviews(numReviews);
			})
			.catch((err) => {
				console.error(err);
				notify.info({ message: `Sorry, we failed to load a the average rating for this book.` })
			})
			.finally(() => setLoading(false))
	};

	useEffect(() => {
		loadBookReviewSummary(bookId);
	}, []);

	if (loading) {
		return <div>Loading average book rating...</div>
	}

	return (
		<div className="book-info-item">
			<div className="book-info-review-aggregate">
				<div className="review-average-rating">
					Average rating:
					<Tooltip label={avgRating / 2}>
						<Rating readOnly fractions={2} value={avgRating / 2} />
					</Tooltip>
				</div>
				<div className="review-number-views">
					{
						numReviews == 0 ? 
						`No reviews yet.` :
						`(Out of ${numReviews} reviews)`
					}
				</div>
			</div>
		</div>
	)
};
