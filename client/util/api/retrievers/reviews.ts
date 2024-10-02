import api from '../axios';

/**
 * Retrieve all books available
 */
const getBookReviews = async (bookId: number) => {
	return await api.get(`/books/${bookId}/reviews`).then((res) => res);
};

/**
 * Retrieve summary of given book's rating(s)
 * Includes avg rating and number of reviews
 */
const getBookAverageAndNumberReviews = async (bookId: number) => {
	return await api.get(`/books/${bookId}/reviews/summary`).then((res) => res);
}

const submitReview = async ({rating, description, bookId} : {rating: number, description: string, bookId: number}) => {
	return await api.post(
		`/books/${bookId}/reviews`, 
		{description, rating},
	).then((res) => res);
}


export default {
	getBookReviews,
	getBookAverageAndNumberReviews,
	submitReview,
}
