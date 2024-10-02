import api from '../axios';

/**
 * Retrieve all books available
 */
const getBookReviews = async (bookId: number) => {
	return await api.get(`/books/${bookId}/reviews`).then((res) => res);
};

const getBookAverageAndNumberReviews = async (bookId: number) => {
	return await api.get(`/books/${bookId}/reviews/summary`).then((res) => res);
}


export default {
	getBookReviews,
	getBookAverageAndNumberReviews,
}
