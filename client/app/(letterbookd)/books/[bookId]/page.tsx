'use client';

import './books.style.css';
import api from 'util/api/api';
import notify from 'util/notify/notify';
import { useState, useEffect } from 'react';
import BookInfoPanel from '../(bookComponents)/bookInfoPanel/bookInfoPanel';
import BookReviews from '../(bookComponents)/bookReviewList/bookReviewList';

export default function Book({ params }: { params: { bookId: number } }) {
	// Uhhhh
	const [updateKey, setUpdateKey] = useState(1);
	const [book, setBook] = useState(null);
	const [loading, setLoading] = useState(false);

	const refresh = async () => {
		console.log("refresh called")
		await loadBook();
		setUpdateKey(updateKey + 1);
	}

	const loadBook = async () => {
		setLoading(true);
		await api.books.getBook(params.bookId)
			.then((res) => setBook(res.data))
			.catch((error) => {
				console.error(error)
				notify.error({ message: `Failed to load this book` })
			})
			.finally(() => setLoading(false));
	};

	useEffect(() => {
		loadBook();
	}, []);


	if (loading || !book) {
		return (
			<div className="book-container">
				<div>Loading...</div>
			</div>
		)
	}

	return (
		<div className="book-container">
			<BookInfoPanel key={updateKey} book={book} reload={refresh}/>
			<BookReviews key={updateKey} bookId={book.Id}/>
		</div>
	)
}
