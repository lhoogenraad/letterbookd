"use client";

import { useState, useEffect } from 'react';
import api from 'util/api/api';
import notify from 'util/notify/notify';

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

	return (
		<div>
		<div>
		bookId: {bookId}
		reviews: {JSON.stringify( reviews )}
		</div>
		</div>
	)
}
