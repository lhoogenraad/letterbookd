import '@mantine/core/styles.css';
import api from 'util/api/server/api';
import './dashboard.style.css';
import BookList from '../books/(bookComponents)/bookList/bookList';


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
	const firstName = "First name uhhh";

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
			{JSON.stringify(reviews)}
			{JSON.stringify(books)}
		</div>
	)
};
