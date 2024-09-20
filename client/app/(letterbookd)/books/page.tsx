"use client"

import api from 'util/api/api';
import notify from 'util/notify/notify';
import { useState, useEffect } from 'react';
import BookTile from './(bookTiles)/bookTile';
import { Input, CloseButton } from '@mantine/core';
import './books.css';

export default function Books() {
	const [books, setBookList] = useState(null);
	const [searchText, setSearchText] = useState('');
	const [loading, setLoading] = useState(false);

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
			Object.values(book).every((bookAttribute) => {
				if (bookAttribute.toString().toLowerCase().includes(searchText.toLowerCase())) {
					bookHasMatching = true;
					return;
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
					placeholder='Search books...'
					value={searchText}
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
				/>
			</div>

			<div className="books-list-container">
				{filteredBooks.map((book: any, index: number) => (
					<div className="book-tile" key={index}>
						<BookTile book={book} />
					</div>
				))}
			</div>
		</div>
	)
};
