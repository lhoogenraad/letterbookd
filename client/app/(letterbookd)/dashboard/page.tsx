"use client";

import '@mantine/core/styles.css';
import api from 'util/api/server/api';
import './dashboard.style.css';
import jwt from "jsonwebtoken";
import BookTile from '../books/(bookComponents)/bookTile/bookTile';
import BookReview from '../books/(bookComponents)/bookReviewList/bookReview';
import Link from 'next/link';
import { useEffect, useState } from 'react';

export const dynamic = 'force-dynamic';

type book = {
	Id: number,
	Title: string,
	Author: string,
	Published: string,
	NumPages: string,
	CoverURL: string,
	Synopsis: string,
	OnUserReadList: string,
}

type review = {
	BookId: number,
}


async function getBooks(): Promise<book[]> {
	let books: book[];
	await api.dashboard.getFeaturedBooks()
		.then((res) => {
			console.log(res)
			books = res.data
		})
	return books;
};

async function getPopularReviews(): Promise<review[]> {
	let reviews: review[];
	await api.dashboard.getPopularReviews()
		.then((res) => reviews = res.data)
	return reviews;
};

export default function Dashboard() {
	function getFirstName() {
		// Get username
		const claims = jwt.decode(localStorage.getItem("authToken"))
		setFirstName(claims.firstName)
	}

	const [books, setBooks] = useState([]);
	const [reviews, setReviews] = useState([]);
	const [firstName, setFirstName] = useState(undefined);

	const init = async () => {
		try {
			getFirstName();
			const booksData = await getBooks();
			const reviewsData = await getPopularReviews();
			console.log(booksData, reviewsData)
			setBooks(booksData);
			setReviews(reviewsData);
		} catch (err) {
			console.error(err);
		}
	};

	useEffect(() => {
		console.log(localStorage.getItem("authToken"))
		init()
	}, []);

	return (
		<div className='dashboard-container'>
			<div className='title-container'>
				<h1 className='dashboard-title'>
					{
						firstName ?
							`Welcome back, ${firstName}!` :
							`Welcome to Shelfd!`
					}
				</h1>
			</div>
			<div className="dashboard-data-columns">
			<div className="data-column book-data-column">
			<h3 className="data-column-title">Here's what we're reading</h3>
				{
					books.map((book) => (
					<Link 
						href={{ pathname: `/books/${book.Id}` }}
						style={{textDecoration: "inherit", color: "inherit"}}
					>
						<BookTile book={book} />
					</Link>
					))
				}
			</div>
			<div className="data-column">
			<h3 className="data-column-title">Popular reviews this week</h3>
				{
					reviews.map((review) => (
						<BookReview showTitle={true} review={review} bookId={review.BookId} />
					))
				}
			</div>
			</div>
		</div>
	)
};
