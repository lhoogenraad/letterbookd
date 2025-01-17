"use client"

import { Button } from '@mantine/core';
import { useState } from 'react';
import api from 'util/api/api';
import notify from 'util/notify/notify';

export default function DeleteComment({ comment, closeModal, reloadComments }) {
	const [loading, setLoading] = useState(false);

	const deleteComment = async () => {
		setLoading(true);
		await api.reviews.deleteReviewComment(comment.ReviewId, comment.Id)
			.then(() => {
				reloadComments();
				closeModal();
			})
			.catch((err) => notify.error({ message: err.response.data.Message }))
			.finally(() => setLoading(false))
	}

	return (
		<div style={{
			display: 'flex',
			flexDirection: 'row',
			justifyContent: 'space-around',
			maxWidth: '500px',
			margin: '0 auto',
			backgroundColor: "primary.0",
		}}>
			<Button onClick={() => closeModal()} variant="subtle" color="primary.0">No, go back!</Button>
			<Button onClick={() => deleteComment()} loading={loading} variant="subtle" color="red">Yes please</Button>
		</div>
	)
}
