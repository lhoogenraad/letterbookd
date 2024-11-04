"use client"

import api from 'util/api/api';
import notify from 'util/notify/notify';
import { useState, useEffect } from 'react';
import BookTile from './(bookComponents)/bookTile/bookTile';
import { Input, Select, Pagination } from '@mantine/core';
import './books.css';
import {IconSearch} from '@tabler/icons-react';
import Link from 'next/link';

export default function Books() {
	const PAGE_SIZE = 50;
	const [books, setBookList] = useState([]);
	const [pagCount, setPagCount] = useState(0);
	const [loading, setLoading] = useState(false);
	const [searchText, setSearchText] = useState('');
	const [currPage, setCurrPage] = useState(1);

	const getBooksCount = async () => {
		await api.books.getBooksCount(searchText)
		.then((res) => setPagCount(res.data / PAGE_SIZE))
		.catch((err) => console.error(err))
	};

	const getBooksList = async (page: number, pageSize: number) => {
		setLoading(true);
		await api.books.getAllBooks(page, pageSize, searchText)
			.then((res) => setBookList(res.data ?? []))
			.catch(() => notify.error({
				message: `Failed to load books list for ` +
					`some reason, please try again later!`
			}))
			.finally(() => setLoading(false));
	};

	const init = async () => {
		setCurrPage(1)
		await getBooksCount();
		await getBooksList(1, PAGE_SIZE);
	}


	useEffect(() => {
		init();
	}, []);


	const filterJSX = (
		<div>
			<div className="books-filters-container">
				<Input
					variant='filled'
					className='books-filters-search'
					placeholder='Search books...'
					value={searchText}
					rightSectionPointerEvents="all"
					rightSection = {
						<IconSearch 
							className="search-button"
							aria-label='Search'
							onClick = {() => init()}
							style={{ display: searchText ? undefined : 'none' }}
						/>
					}
					onChange={(ev) => {
						return setSearchText(ev.currentTarget.value);
					}}
				></Input>

				<div className='books-filters-select-container'>
					<Select 
						clearable
						disabled
						label='Genre'
						data={['Fantasy', 'History', 'Drama']}
					/>
				</div>
			</div>

			<div className="pagination-container">
				<Pagination
					color="primary"
					withEdges
					withControls
					value={currPage}
					total={pagCount}
					onChange={(val: number) => {
						getBooksList(val, PAGE_SIZE)
						setCurrPage(val)
					}}
				/>
			</div>
		</div>
	)

	if (loading) {
		return (
			<div className="books-container">
				{filterJSX}
				<div>Loading...</div>
			</div>
		)
	}
	return (
		<div className="books-container">

			{filterJSX}

			<div className="books-list-container">
				{books.map((book: any, index: number) => (
					<Link href={{pathname: `/books/${book.Id}`}}>
					<div className="book-tile" key={index}>
						<BookTile book={book} />
					</div>
					</Link>
				))}
			</div>
			<div className="pagination-container">
				<Pagination
					color="primary"
					withEdges
					withControls
					value={currPage}
					total={pagCount}
					onChange={(val: number) => {
						getBooksList(val, PAGE_SIZE)
						setCurrPage(val)
					}}
				/>
			</div>
		</div>
	)
};
