"use client"

import { useState, useEffect } from 'react';
import '@mantine/core/styles.css';
import notify from 'util/notify/notify';
import api from 'util/api/api';
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
		return <p>It appears you've got no books on your read list!</p>
	}

	return (
		<div className='readlist-container'>
		<div className='readlist-header-container'>
		<h2>Your readlist</h2>
		<MultiSelect 
			className='status-filter'
			data={['Read', 'Unread']} 
			onChange={setSearchFilters} 
			label="Status"
		/>
		</div>
		<div className='readlist-items'>
		<div className='readlist-item-count'>{filteredReadlist.length} Items</div>
			{filteredReadlist.map((listItem: any, index: number) => {
				return (
					<ReadListItem item={listItem} key={index}/>
				)
			})}
			</div>
		</div>
	)
};

