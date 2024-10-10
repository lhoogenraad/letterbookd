"use client"

import { Button } from '@mantine/core';
import { useState } from 'react';
import api from 'util/api/api';
import notify from 'util/notify/notify';

export default function DeleteComment({ comment, closeModal }) {
	const [loading, setLoading] = useState(false);

	const deleteComment = async () => {
		setLoading(true);
		await api.reviews.deleteReviewComment(comment.ReviewId, comment.Id)
			.then(() => closeModal())
			.catch((err) => notify.error({ message: err.response.data.Message }))
			.finally(() => setLoading(false))
	}

	return (
		<div style={{
			display: 'flex',
			flexDirection: 'row',
			justifyContent: 'space-around'
		}}>
			<Button onClick={() => closeModal()} variant="subtle" color="blue">No, go back!</Button>
			<Button onClick={() => deleteComment()} loading={loading} variant="subtle" color="red">Yes please</Button>
		</div>
	)
}
