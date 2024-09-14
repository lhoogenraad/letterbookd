"use client"

import { useState, useEffect } from 'react';
import { NextPageContext } from 'next';
import notify from 'util/notify/notify';
import api from 'util/api/api';
import './dashboard.style.css';

async function getBooks() {
	let books: object;
	await api.books.getAllBooks()
	.then((res) => books = res.data)
	return books;
};

export default function Dashboard() {
	const [books, setBooks] = useState(null);
	const [loading, setLoading] = useState(true);

	useEffect(() => {
		getBooks()
		.then((books) => {
			setBooks(books);
			setLoading(false);
		})
		.catch((err) => notify.info({ message: err?.response?.data?.Message }));
	}, []);


	if (loading) { 
		return <p>Loading...</p>
	}
	if (!books) { 
		return <p>No books available</p>
	}

	return (
		<div className='book-list-container'>
		{books.map((book: any) => (
			<div className='book-list-item'>
			<div className='book-title'>
			{book.Title}
			</div>
			<div className='book-author'>
			{book.Author}
			</div>
			<div className='book-published-date'>
			{new Date(book.Published).toLocaleDateString()}
			</div>
			</div>
		))}
		</div>
	)
};

