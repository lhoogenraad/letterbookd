import api from '../axios';

/**
 * Retrieve all books available
 */
const getBookReviews = async (bookId: number) => {
	return await api.get(`/books/${bookId}/reviews`).then((res) => res);
};


export default {
	getBookReviews,
}
