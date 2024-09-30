'use client';

import './books.style.css';
import api from 'util/api/api';
import notify from 'util/notify/notify';
import { useState, useEffect } from 'react';

export default function Book({ params }: { params: { bookId: number } }) {
	const [book, setBook] = useState(null);
	const [loading, setLoading] = useState(false);

	const loadBook = async () => {
		setLoading(true);
		await api.books.getBook(params.bookId)
			.then((res) => setBook(res.data))
			.catch((error) => {
				console.error(error)
				notify.info({ message: `Failed to load this book` })
			})
			.finally(() => setLoading(false));
	};

	useEffect(() => {
		loadBook();
	}, []);


	if (loading) {
		return (
			<div className="book-container">
				<div>Loading...</div>
			</div>
		)
	}

	return (
		<div className="book-container">
			Param: {params.bookId}
			Book: {JSON.stringify(book)}
		</div>
	)

}
