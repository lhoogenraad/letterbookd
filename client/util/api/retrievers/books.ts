import api from '../axios';

/**
 * Retrieve all books available
 */
const getAllBooks = async () => {
	return await api.get('/books').then((res) => res);
};

const getBook = async (id: number) => {
	return await api.get(`/books/${id}`).then((res) => res);
}

export default {
	getAllBooks,
	getBook,
}
