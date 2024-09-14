"use client"

import { useState, useEffect } from 'react';
import notify from 'util/notify/notify';
import api from 'util/api/api';

export default function Dashboard() {
	const [books, setBooks] = useState(null);
	const [loading, setLoading] = useState(true);

	async function getBooks() {
		let books: object;
		await api.books.getAllBooks()
		.then((res) => books = res.data)
		return books;
	};

	useEffect(() => {
		getBooks()
		.then((books) => {
			setBooks(books);
			setLoading(false);
		})
		.catch((err) => notify.info({message: err?.response?.data?.Message}));
	});


	if (loading) return <p>Loading...</p>
		if (!books) return <p>No books available</p>

			return (
				<div>
				<ul>
				{books.map((book: any) => (
					<li>{book.Title}</li>
				))}
				</ul>
				</div>
			)
};

