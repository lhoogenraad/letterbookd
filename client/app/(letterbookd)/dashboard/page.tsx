import '@mantine/core/styles.css';
import api from 'util/api/server/api';
import './dashboard.style.css';
import { cookies } from 'next/headers';
import { decodeJwt } from 'jose';
import BookTile from '../books/(bookComponents)/bookTile/bookTile';
import BookReview from '../books/(bookComponents)/bookReviewList/bookReview';
import Link from 'next/link';

function getFirstName() {
	const cookieStore = cookies();
	const token = cookieStore.get("authToken")?.value;
	const claims = decodeJwt(token);
	if (!claims) return;

	return claims.firstName;
}

async function getBooks() {
	let books: object[];
	await api.dashboard.getFeaturedBooks()
		.then((res) => books = res.data)
	return books;
};

async function getPopularReviews() {
	let reviews: object[];
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
							`Welcome to letterbookd!`
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
