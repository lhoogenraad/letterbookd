"use client"

import { useState, useEffect } from 'react';
import jwt from 'jsonwebtoken';
import '@mantine/core/styles.css';
import notify from 'util/notify/notify';
import { IconBook } from '@tabler/icons-react';
import { Button } from '@mantine/core';
import api from 'util/api/api';
import './dashboard.style.css';
import { Carousel } from '@mantine/carousel';
import BookList from './bookList';

async function getBooks() {
	let books: object;
	await api.books.getAllBooks()
	.then((res) => books = res.data)
	return books;
};

export default function Dashboard() {
	const [books, setBooks] = useState(null);
	const [loading, setLoading] = useState(true);
	const [firstName, setFirstName] = useState(undefined);

	useEffect(() => {
		// Get books
		getBooks()
		.then((books) => {
			setBooks(books);
			setLoading(false);
		})
		.catch((err) => notify.info({ message: err?.response?.data?.Message }));

		// Get username
		const claims = jwt.decode(localStorage.getItem("authToken"))
		setFirstName(claims.firstName);
	}, []);


	if (loading) { 
		return <p>Loading...</p>
	}
	if (!books) { 
		return <p>No books available</p>
	}

	return (
		<div className='dashboard-container'>
		<div className='navbar-container'>
		<div>
			<IconBook/>
		</div>
		<div className='font-family'>
			Letterbookd
		</div>
		<div className='navbar-action-buttons'>
			<Button variant="filled">My Read List</Button>
			<Button>My Account</Button>
			<Button>Sign Out</Button>
		</div>
		</div>
		<div className='title-container'>
			<h1 className='title'>
				{
					firstName ?
					`Welcome back, ${firstName}!` :
					`Welcome to letterbookd!`
				}
			</h1>
		</div>
		<div className='book-list-container'>
		<div className="font-family">
		Trending books
		</div>
		<BookList books={books}/>
		</div>
		</div>
	)
};

