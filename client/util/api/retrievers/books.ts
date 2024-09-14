import api from '../axios';

/**
 * Retrieve all books available
 */
const getAllBooks = async () => {
	return await api.get('/books').then((res) => res);
};

export default {
	getAllBooks,
}
