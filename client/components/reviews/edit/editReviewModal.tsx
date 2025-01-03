"use client"

import api from "util/api/api";
import notify from "util/notify/notify";
import { Rating, Textarea, Button } from '@mantine/core';
import { useForm } from '@mantine/form';
import { useState } from 'react';

export default function EditReviewModal({review, book, closeModal, updateReview}) {
	const [submitting, setSubmitting] = useState(false);
	const submitReviewEdit = async (bookId: number, rating: number, description: string) => {
		setSubmitting(true);
		api.reviews.editReview({reviewId: review.Id, rating, description})
		.then(() => {
			console.log("THEN CALLED")
			updateReview();
			closeModal();
		})
		.catch((err) => notify.error({message: err.response?.data?.Message ?? err.message}))
		.finally(() => setSubmitting(false));
	}

	const form = useForm({
		mode: 'uncontrolled',
		initialValues: {
			rating: review.Rating,
			description: review.Description,
		},

		validate: {
			rating: (value) => (!!value ? null : 'Please provide a rating.'),
			description: (value) => (!!value || parseFloat(value) < 1 ? null : 'Please provide a review description.'),
		},
	});

	return (
		<div className="review-container">
			<form 
				className="review-form"
				onSubmit={form.onSubmit(async (values) => 
					await submitReviewEdit(book.Id, parseFloat(values.rating)*2, values.description))}
			>

			<Rating 
				size='lg'
				fractions={2} 
				{...form.getInputProps('rating')}
				key={form.key('rating')} 
			/>
			<Textarea
				key={form.key('description')}
				{...form.getInputProps('description')}
				placeholder="Your review here..."
				autosize
				minRows={5}
				maxRows={8}
		  	/>

			<Button 
				className="add-review-button"
				loading={submitting} 
				loaderProps={{ type: 'bars' }} 
				type="Update"
			>
			Submit
			</Button>
			</form>
		</div>
	)
}
