"use client"

import api from 'util/api/api';
import notify from 'util/notify/notify';
import { useState, useEffect } from 'react';
import BookTile from './(bookComponents)/bookTile/bookTile';
import { Input, CloseButton, Select, Pagination } from '@mantine/core';
import './books.css';
import Link from 'next/link';

export default function Books() {
	const PAGE_SIZE = 50;
	const [books, setBookList] = useState(null);
	const [loading, setLoading] = useState(false);
	const [searchText, setSearchText] = useState('');
	const [currPage, setCurrPage] = useState(1);

	const getBooksList = async (page: number, pageSize: number) => {
		setLoading(true);
		await api.books.getAllBooks(page, pageSize)
			.then((res) => setBookList(res.data))
			.catch(() => notify.error({
				message: `Failed to load books list for ` +
					`some reason, please try again later!`
			}))
			.finally(() => setLoading(false));
	};

	const filterBookList = (books: Array<object>, searchText: string) => {
		if (!searchText || searchText === "") return books;

		return books.filter((book: object) => {
			let bookHasMatching = false;
			Object.values(book).forEach((bookAttribute) => {

				try {
					if (bookAttribute.toString().toLowerCase().includes(searchText.toLowerCase())) {
						bookHasMatching = true;
						return;
					}
				} catch (error) {
					console.error(error.message);
					console.error("Failed to filter on book", book);
				}
			});
			return bookHasMatching;
		});
	};

	const filteredBooks = filterBookList(books, searchText);

	useEffect(() => {
		getBooksList(1, PAGE_SIZE);
	}, []);


	if (loading || !books) {
		return <div>Loading..</div>
	}
	return (
		<div className="books-container">

			<div className="books-filters-container">
				<Input
					variant='filled'
					className='books-filters-search'
					placeholder='Search books...'
					value={searchText}
					rightSectionPointerEvents="all"
					rightSection = {
						<CloseButton 
							aria-label='Clear input'
							onClick = {() => setSearchText('')}
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
					withEdges
					withControls
					value={currPage}
					total={10}
					onChange={(val: number) => {
						getBooksList(val, PAGE_SIZE)
						setCurrPage(val)
					}}
				/>
			</div>

			<div className="books-list-container">
				{filteredBooks.map((book: any, index: number) => (
					<Link href={{pathname: `/books/${book.Id}`}}>
					<div className="book-tile" key={index}>
						<BookTile book={book} />
					</div>
					</Link>
				))}
			</div>
			<div className="pagination-container">
				<Pagination
					withEdges
					withControls
					value={currPage}
					total={10}
					onChange={(val: number) => {
						getBooksList(val, PAGE_SIZE)
						setCurrPage(val)
					}}
				/>
			</div>
		</div>
	)
};
