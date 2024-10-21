"use client";

import './readlistItem.css';
import { Image, Select, Loader } from '@mantine/core';
import Link from 'next/link';
import api from 'util/api/api';
import notify from 'util/notify/notify';
import statusOptions from 'configs/readlistStatusOptions';
import { useState } from 'react';

export default function ReadListItem({ item }) {
	const [status, setStatus] = useState(item.Status)
	const [loading, setLoading] = useState(false);

	const updateStatus = async (newStatus: string) => {
		// If status value is same as current status, don't send request
		if (newStatus === status) return;
		setLoading(true);
		await api.readlist.updateReadListItem(item.BookId, newStatus)
			.then(() => setStatus(newStatus))
			.catch((err) => notify.error({ message: err?.response?.data?.Message }))
			.finally(() => setLoading(false))
	}

	return (
		<div className="readlist-item-container">
			<div className="readlist-item-header-container">
				<div className="img">
					<Image
						src={null}
						h='5rem'
						radius="md"
						fallbackSrc="https://placehold.co/600x400?text=Placeholder"
					/>
				</div>
				<Link
					style={{ textDecoration: "inherit", color: "inherit" }}
					href={`/books/${item.BookId}`}
				>
					<div className="title">{item.BookName}</div>
				</Link>
			</div>
			<div className="readlist-item-body-container">
				<Select
					variant='unstyled'
					className="readlist-status"
					data={statusOptions}
					rightSection={loading ? <Loader size='xs' /> : null}
					value={status}
					disabled={loading}
					onChange={async (val) => await updateStatus(val)}
				/>
				{
					status == 'Read' ?
						'Create review here'
						:
						null
				}
				<div className="readlist-date">01/04/2022</div>

			</div>
		</div>
	)
}
