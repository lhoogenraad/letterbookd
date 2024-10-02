'use client';

import { Rating, Textarea, Button } from '@mantine/core';
import { useForm } from '@mantine/form';
import { useState } from 'react';
import api from 'util/api/api';
import notify from 'util/notify/notify';
import './addReview.css'

export default function AddReview({book, reload, closeModal}) {
	const [submitting, setSubmitting] = useState(false);

	const submitReview = async (bookId: number, rating: number, description: string) => {
		setSubmitting(true);
		api.reviews.submitReview({bookId, rating, description})
		.then(() => {
			reload()
			closeModal()
		})
		.catch((err) => notify.error({message: err.response?.data?.Message ?? err.message}))
		.finally(() => setSubmitting(false));
	}


	const form = useForm({
		mode: 'uncontrolled',
		initialValues: {
			rating: null,
			description: null,
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
					await submitReview(book.Id, parseFloat(values.rating)*2, values.description))}
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
				type="submit"
			>
			Submit
			</Button>
			</form>
		</div>
	)
}
