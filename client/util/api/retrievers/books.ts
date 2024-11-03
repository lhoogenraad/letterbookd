import api from '../axios';

/**
 * Retrieve all books available
 */
const getAllBooks = async (page: number, pageSize: number, filter: string) => {
	return await api.get(`/books?page=${page}&pageSize=${pageSize}?&searchText=${filter}`).then((res) => res);
};

const getBooksCount = async (filter: string) => {
	return await api.get(`/books/count?&searchText=${filter}`).then((res) => res);
}

const getBook = async (id: number) => {
	return await api.get(`/books/${id}`).then((res) => res);
}

export default {
	getAllBooks,
	getBook,
	getBooksCount,
}
