"use client"

import { useState, useEffect } from 'react';
import '@mantine/core/styles.css';
import notify from 'util/notify/notify';
import api from 'util/api/api';
import './readlist.css';
import ReadListItem from './(readlistItem)/readlistItem';

async function getReadList() {
	let books: object;
	await api.readlist.getReadList()
		.then((res) => books = res.data)
	return books;
};

export default function Dashboard() {
	const [readList, setReadList] = useState(null);
	const [loading, setLoading] = useState(true);

	useEffect(() => {
		// Get read list
		getReadList()
			.then((list) => {
				setReadList(list);
				setLoading(false);
			})
			.catch((err) => notify.error({ message: err?.response?.data?.Message }));
	}, []);


	if (loading) {
		return <p>Loading...</p>
	}
	if (!readList || !readList.length) {
		return <p>It appears you've got no books on your read list!</p>
	}

	return (
		<div className='readlist-container'>
			{readList.map((listItem: any) => {
				return (
					<ReadListItem item={listItem} />
				)
			})}
		</div>
	)
};

