"use client"

import { Button, Textarea } from '@mantine/core';
import { useForm } from '@mantine/form';
import { useState } from 'react';
import api from 'util/api/api';
import notify from 'util/notify/notify';

export default function EditComment({ comment, closeModal, reloadComments }) {
	const [loading, setLoading] = useState(false);

	const form = useForm({
		mode: 'uncontrolled',
		initialValues: {
			comment: comment.Comment,
		},

		validate: {
			comment: (value) => (!!value ? null : 'Comment required'),
		},
	});

	const editComment = async (commentText: string) => {
		setLoading(true);
		await api.reviews.putReviewComment(comment.ReviewId, comment.Id, commentText)
			.then(() => {
				reloadComments();
				closeModal();
			})
			.catch((err) => notify.error({ message: err.response.data.Message }))
			.finally(() => setLoading(false))
	}

	return (
		<div>
		<form 
			onSubmit={
				form.onSubmit (async (values) => 
					await editComment(values.comment)
			)} 
		>
			<Textarea 
				key={form.key('comment')}
				{...form.getInputProps('comment')}
			/>

			<div style={{
				display: 'flex',
				flexDirection: 'row',
				justifyContent: 'space-around',
				maxWidth: '500px',
				margin: '15px auto 0px auto',
			}}>
				<Button onClick={() => closeModal()} variant="subtle" color="blue">No, go back!</Button>
				<Button 
					type="submit"
					loading={loading} 
					variant="subtle" 
					color="green"
				>
					Submit
				</Button>
			</div>
			</form>
		</div>
	)
}
