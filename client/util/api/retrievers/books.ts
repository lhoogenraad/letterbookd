import api from '../axios';

/**
 * Retrieve all books available
 */
const getAllBooks = async () => {
	return await api.get('/books').then((res) => res);
};

const getBook = async (id: number) => {
	console.log('getting book')
	return await api.get(`/books/${id}`).then((res) => res);
}

export default {
	getAllBooks,
	getBook,
}
