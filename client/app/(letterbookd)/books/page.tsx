"use client"

import api from 'util/api/api';
import notify from 'util/notify/notify';
import { useState, useEffect } from 'react';
import BookTile from './bookTile';


export default  function Books() {
	const [books, setBookList] = useState(null);
	const [loading, setLoading] = useState(false);

	const getBooksList = async () => {
		setLoading(true);
		await api.books.getAllBooks()
			.then((res) => setBookList(res.data))
			.catch(() => notify.info({message: `Failed to load books list for ` + 
										`some reason, please try again later!`}))
			.finally(() => setLoading(false));
	};

	useEffect(() => {
		getBooksList();
	}, []);

	if (loading || !books) {
		return <div>Loading..</div>
	}

	return (
		<div>
			{books.map((book: any, index: number) => <BookTile key={index} book={book} />)}
		</div>
	)
};
