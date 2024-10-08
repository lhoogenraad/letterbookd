"use client";

import { useState, useEffect } from 'react';
import api from 'util/api/api';
import notify from 'util/notify/notify';
import BookReview from './bookReview';
import './reviewList.css';

export default function BookReviewList({bookId}) {
	const [reviews, setReviews] = useState([]);
	const [loading, setLoading] = useState(false);

	const getBookReviews = async () => {
		setLoading(true);
		await api.reviews.getBookReviews(bookId)
		.then((res) => {
			if (res.data) setReviews(res.data)
		})
		.catch((err) => notify.error({message: `Failed to load reviews for this book. ${err.message}`}))
		.finally(() => setLoading(false));
	};

	useEffect(() => {
		getBookReviews();
	}, []);

	if (loading) {
		return <div>Loading/Rendering Book reviews...</div>
	}

	return (
		<div className="book-reviews-container">
			{reviews.map((review: any, index: number) => (
				<div key={index}>
					<BookReview review={review} bookId={bookId}/>
				</div>
			))}
		</div>
	)
}
