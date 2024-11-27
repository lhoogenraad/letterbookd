import '@mantine/core/styles.css';
import api from 'util/api/server/api';
import './dashboard.style.css';
import { cookies } from 'next/headers';
import { decodeJwt } from 'jose';
import BookTile from '../books/(bookComponents)/bookTile/bookTile';
import BookReview from '../books/(bookComponents)/bookReviewList/bookReview';
import Link from 'next/link';

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

function getFirstName() {
	const cookieStore = cookies();
	const token = cookieStore.get("authToken")?.value;
	const claims = decodeJwt(token);
	if (!claims) return;

	return claims.firstName;
}

async function getBooks(): Promise<book[]> {
	let books: book[];
	await api.dashboard.getFeaturedBooks()
		.then((res) => books = res.data)
	return books;
};

async function getPopularReviews(): Promise<review[]> {
	let reviews: review[];
	await api.dashboard.getPopularReviews()
		.then((res) => reviews = res.data)
	return reviews;
};

export default async function Dashboard() {
	const books = await getBooks();
	const reviews = await getPopularReviews();
	const firstName = getFirstName();

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
