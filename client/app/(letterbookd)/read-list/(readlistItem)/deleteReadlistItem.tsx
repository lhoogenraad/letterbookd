"use client"

import { Button } from '@mantine/core';
import { useState } from 'react';
import api from 'util/api/api';
import notify from 'util/notify/notify';

export default function DeleteReadlistItem({ bookId, closeModal, removeItem }) {
	const [loading, setLoading] = useState(false);

	const deleteReadlistItem = async () => {
		setLoading(true);
		await api.readlist.removeReadlistItem(bookId)
			.then(() => {
				closeModal();
				removeItem();
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
		}}>
			<Button onClick={() => closeModal()} variant="subtle" color="primary.0">No, go back!</Button>
			<Button onClick={() => deleteReadlistItem()} loading={loading} variant="subtle" color="red">Yes please</Button>
		</div>
	)
}
