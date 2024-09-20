"use client"

import api from 'util/api';
import notify from 'util/notify';
import { useState, useEffect } from 'react';


export default function Books() {
	const [books, setBookList] = useState(null);
	const [loading, setLoading] = useState(true);

	useEffect(() => {

	}, []);


	return (
		<div>
			hello
			<div>{books}</div>
			<div>{loading}</div>
		</div>
	)
};
