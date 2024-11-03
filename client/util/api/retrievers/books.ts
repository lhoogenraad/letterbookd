import api from '../axios';

/**
 * Retrieve all books available
 */
const getAllBooks = async (page: number, pageSize: number) => {
	return await api.get(`/books?page=${page}&pageSize=${pageSize}`).then((res) => res);
};

const getBook = async (id: number) => {
	return await api.get(`/books/${id}`).then((res) => res);
}

export default {
	getAllBooks,
	getBook,
}
