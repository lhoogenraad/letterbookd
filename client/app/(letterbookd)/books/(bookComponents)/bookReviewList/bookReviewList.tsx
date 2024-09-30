"use client";

import { useState, useEffect } from 'react';
import api from 'util/api/api';
import notify from 'util/notify/notify';
import BookReview from './bookReview';
import './reviewList.css';

export default function BookReviewList({bookId}) {
	const [reviews, setReviews] = useState(null);
	const [loading, setLoading] = useState(false);

	const getBookReviews = async () => {
		setLoading(true);
		await api.reviews.getBookReviews(bookId)
		.then((res) => setReviews(res.data))
		.catch((err) => notify.info({message: `Failed to load reviews for this book. ${err.message}`}))
		.finally(() => setLoading(false));
	};

	useEffect(() => {
		getBookReviews();
	}, []);

	if (loading || !reviews) {
		return <div>Loading/Rendering...</div>
	}

	return (
		<div className="book-reviews-container">
			{reviews.map((review: any, index: number) => (
				<div key={index}>
					<BookReview review={review}/>
				</div>
			))}
		</div>
	)
}
