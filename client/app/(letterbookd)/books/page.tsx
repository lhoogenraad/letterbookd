"use client"

import api from 'util/api/api';
import notify from 'util/notify/notify';
import { useState, useEffect } from 'react';
import BookTile from './(bookComponents)/bookTile/bookTile';
import { Input, CloseButton, Select, MultiSelect } from '@mantine/core';
import './books.css';
import Link from 'next/link';

export default function Books() {
	const [books, setBookList] = useState(null);
	const [loading, setLoading] = useState(false);
	const [searchText, setSearchText] = useState('');

	const getBooksList = async () => {
		setLoading(true);
		await api.books.getAllBooks()
			.then((res) => setBookList(res.data))
			.catch(() => notify.info({
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
		getBooksList();
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
						label='Genre'
						data={['Fantasy', 'History', 'Drama']}
					/>

					<MultiSelect 
						clearable
						label='Rating'
						data={['1', '2', '3', '4', '5']}
					/>
				</div>
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
		</div>
	)
};
