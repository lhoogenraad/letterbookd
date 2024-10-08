"use client";

import './addComment.css';

import { Textarea, Button } from '@mantine/core';
import { useForm } from '@mantine/form';

import notify from 'util/notify/notify';
import api from 'util/api/api';
import { useState } from 'react';

export default function AddComment({ reviewId, reload }) {
	const [submitting, setSubmitting] = useState(false);

	const form = useForm({
		mode: 'uncontrolled',
		initialValues: {
			comment: null,
		},
		validate: {
			comment: (value) => (!!value ? null : 'Please leave a comment'),
		},
	});

	const submitComment = async (comment) => {
		setSubmitting(true);
		await api.reviews.submitReviewComment(reviewId, comment)
			.then((res) => {
				notify.info({message: res.data});
				form.setFieldValue('comment', null)
				reload();
			})
			.catch((err) => notify.error(err.response.data.Message))
			.finally(() => setSubmitting(false));
	}

	return (
		<div className="add-comment-container">
			<form 
				onSubmit={form.onSubmit(async (values) => await submitComment(values.comment))}
				className="add-comment-form"
			>
			<Textarea 
				label="Comment" 
				placeholder="Leave your comment here..."
				key={form.key('comment')}
				{...form.getInputProps('comment')}
			/>
			<Button 
				loading={submitting}
				loaderProps={{ type: 'bars' }} 
				type="submit">
			Submit
			</Button>
			</form>
		</div>
	)
}
