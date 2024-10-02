'use client';


import { Rating } from '@mantine/core';
import { useState } from 'react';
import api from 'util/api/api';
import notify from 'util/notify/notify';

export default function AddReview({book}) {
	const [submitting, setSubmitting] = useState(false);
	const [rating, setRating] = useState(null);
	const [description, setDescription] = useState(null);

	const submitReview = () => {
		setSubmitting(true);
		api.reviews.submitReview({bookId: book.Id, rating, description})
		.then((res) => console.log(res.data))
		.catch((err) => notify.error({message: err.response?.data?.Message}))
		.finally(() => setSubmitting(false));
	}

	return (
		<div className="review-container">
			Hey 
			{JSON.stringify(book)}
		</div>
	)
}
