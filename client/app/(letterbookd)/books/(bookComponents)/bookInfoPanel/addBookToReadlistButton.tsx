"use client"

import api from 'util/api/api';
import notify from 'util/notify/notify';
import { Button } from '@mantine/core';
import { IconTablePlus } from '@tabler/icons-react';
import { useState } from 'react';

export default function BookInfoPanel({ bookId }: { bookId: number }) {
	const [loading, setLoading] = useState(false);
	const defaultErrorMessage = `Sorry, something went wrong adding this book to your readlist.`;

	const sendAddBookToReadlistRequest = (bookId: number) => {
		setLoading(true);
		api.readlist.addBookToReadList(bookId, 'Unread')
			.then((res) => notify.info({message: res.data ?? `Successfully added book`}))
			.catch((err) => notify.error({ message: err.response?.data?.Message ?? defaultErrorMessage }))
			.finally(() => setLoading(false))
	};

	return (
		<div>
			<Button
				rightSection={<IconTablePlus size={20} />}
				fullWidth
				loading={loading}
				onClick={() => sendAddBookToReadlistRequest(bookId)}
			>
				Add to my Readlist
			</Button>
		</div>
	)
}
