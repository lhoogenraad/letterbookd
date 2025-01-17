"use client"

import api from 'util/api/api';
import notify from 'util/notify/notify';
import { Button } from '@mantine/core';
import { IconTablePlus } from '@tabler/icons-react';
import { useState } from 'react';

export default function BookInfoPanel({ bookId, disabled, setBookAdded }: { bookId: number, disabled: boolean, setBookAdded: Function }) {
	const [loading, setLoading] = useState(false);
	const defaultErrorMessage = `Sorry, something went wrong adding this book to your readlist.`;

	const sendAddBookToReadlistRequest = (bookId: number) => {
		setLoading(true);
		api.readlist.addBookToReadList(bookId, 'Unread')
			.then((res) => {
				notify.info({ message: res.data ?? `Successfully added book` })
				setBookAdded();
			})
			.catch((err) => notify.error({ message: err.response?.data?.Message ?? defaultErrorMessage }))
			.finally(() => setLoading(false))
	};

	return (
		<div>
			<Button
				rightSection={<IconTablePlus size={20} />}
				fullWidth
				disabled={disabled}
				loading={loading}
				onClick={() => sendAddBookToReadlistRequest(bookId)}
				color="secondary"
			>
				Add to my Readlist
			</Button>
		</div>
	)
}
