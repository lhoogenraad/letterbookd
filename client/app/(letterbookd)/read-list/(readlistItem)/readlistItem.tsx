"use client";

import './readlistItem.css';
import { Image, Select } from '@mantine/core';
import Link from 'next/link';
import api from 'util/api/api';
import notify from 'util/notify/notify';
import statusOptions from 'configs/readlistStatusOptions';
import { useState } from 'react';

export default function ReadListItem({ item }) {
	const [status, setStatus] = useState(item.Status)
	const [loading, setLoading] = useState(false);

	const updateStatus = async (status: string) => {
		setLoading(true);
		await api.readlist.updateReadListItem(item.BookId, status)
			.then(() => setStatus(status))
			.catch((err) => notify.error({ message: err?.response?.data?.Message }))
			.finally(() => setLoading(false))
	}

	return (
		<div className="readlist-item-container">
			<Link
				style={{ textDecoration: "inherit", color: "inherit" }}
				href={`/books/${item.BookId}`}
			>
				<div className="readlist-item-header-container">
					<div className="img">
						<Image
							src={null}
							h='5rem'
							radius="md"
							fallbackSrc="https://placehold.co/600x400?text=Placeholder"
						/>
					</div>
					<div className="title">{item.BookName}</div>
				</div>
			</Link>
			<div className="readlist-item-body-container">
				<Select 
					className="readlist-status"
					data={statusOptions}
					value={status}
					disabled={loading}
					onChange={async (val, _) => await updateStatus(val)}
				/>
				<div className="readlist-date">01/04/2022</div>
			</div>
		</div>
	)
}
