"use client";

import { IconBooks, IconRosetteDiscountCheckFilled } from '@tabler/icons-react';
import { Rating, Tooltip } from '@mantine/core';
import { useState, useEffect } from 'react';
import api from 'util/api/api';
import notify from 'util/notify/notify';

export default function BookReviewAverage({ bookId }) {
	const [loading, setLoading] = useState(false);
	const [stats, setStats] = useState({});

	const loadBookReviewSummary = (bookId: number) => {
		setLoading(true);
		api.reviews.getBookAverageAndNumberReviews(bookId)
			.then((res) => {
				setStats(res.data)
			})
			.catch((err) => {
				console.error(err);
				notify.error({ message: `Sorry, we failed to load a the average rating for this book.` })
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
					<Tooltip label={stats.avgRating / 2}>
						<Rating readOnly fractions={2} value={stats.avgRating / 2} />
					</Tooltip>
				</div>
				<div className="review-number-views">
					{
						stats.numReviews == 0 ?
							`No reviews yet.` :
							`(Out of ${stats.numReviews} reviews)`
					}
				</div>
				<div className="readlist-stats">
					<Tooltip label={"Number of completed reads"}>
						<div className="readlist-stats-item">
							<IconRosetteDiscountCheckFilled />
							<div>{stats.numCompletedReads}</div>
						</div>
					</Tooltip>

					<Tooltip label={"Number of times this book was added to someone's readlist"}>
						<div className="readlist-stats-item">
							<IconBooks />
							<div>{stats.numReadlistOccurences}</div>
						</div>
					</Tooltip>
				</div>
			</div>
		</div>
	)
};
