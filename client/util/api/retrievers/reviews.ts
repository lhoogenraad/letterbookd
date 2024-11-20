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

const deleteReviewComment = async (reviewId: number, commentId: number) => {
	return await api.delete(`/reviews/${reviewId}/comments/${commentId}`).then((res) => res);
}

const putReviewComment = async (reviewId: number, commentId: number, comment: string) => {
	return await api.put(`/reviews/${reviewId}/comments/${commentId}`, {comment}).then((res) => res);
}

const setLikeStatus = async (reviewId: number, liked: boolean) => {
	const status = liked ? "like" : "unlike";
	return await api.post(`/reviews/${reviewId}/${status}`).then((res) => res);
}

export default {
	getBookReviews,
	getReviewComments,
	getBookAverageAndNumberReviews,
	submitReview,
	submitReviewComment,
	deleteReviewComment,
	putReviewComment,
	setLikeStatus,
}
