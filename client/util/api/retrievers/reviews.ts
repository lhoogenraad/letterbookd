import api from '../axios';

/**
 * Retrieve all reviews for book
 */
const getBookReviews = async (bookId: number) => {
	return await api.get(`/books/${bookId}/reviews`).then((res) => res);
};

/**
 * Retrieve review comments for a given review
 */
const getReviewComments = async (reviewId: number) => {
	return await api.get(`/reviews/${reviewId}/comments`).then((res) => res);
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

const submitReviewComment = async (reviewId: number, comment: string) => {
	return await api.post(
		`/reviews/${reviewId}/comments`,
		{comment},
	);
}


export default {
	getBookReviews,
	getReviewComments,
	getBookAverageAndNumberReviews,
	submitReview,
	submitReviewComment,
}
