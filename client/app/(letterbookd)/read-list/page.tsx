"use client"

import { useState, useEffect } from 'react';
import '@mantine/core/styles.css';
import notify from 'util/notify/notify';
import api from 'util/api/api';
import statusOptions from 'configs/readlistStatusOptions';
import { MultiSelect } from '@mantine/core';
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
	const [searchFilters, setSearchFilters] = useState<string[]>([]);

	const filterReadlist = () => {
		// If no filters selected, return everything
		if(searchFilters.length <= 0) return readList;
		return readList.filter((item: any) => {
			// Filter based on matching status value
			return searchFilters.includes(item.Status);
		})
	}

	const removeItem = (bookId: number) => {
		const newArr = readList.filter((item: any) => {
			return item.BookId != bookId
		});
		setReadList(newArr);
	}

	const filteredReadlist = filterReadlist();

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
		return (
		<div className='readlist-container'>
			<div className='readlist-header-container'>
				<h1>Your readlist</h1>
				<p>It appears you've got no books on your read list!</p>
			</div>
		</div>
		)
	}

	return (
		<div className='readlist-container'>
		<div className='readlist-header-container'>
		<h1>Your readlist</h1>
		<MultiSelect 
			className='status-filter'
			data={statusOptions} 
			onChange={setSearchFilters} 
			label="Status"
		/>
		</div>
		<div className='readlist-items'>
		<div className='readlist-item-count'>{filteredReadlist.length} Items</div>
			{filteredReadlist.map((listItem: any, index: number) => {
				return (
					<ReadListItem
						removeItem={() => removeItem(listItem.BookId)}
						item={listItem} 
						key={index}
					/>
				)
			})}
			</div>
		</div>
	)
};

