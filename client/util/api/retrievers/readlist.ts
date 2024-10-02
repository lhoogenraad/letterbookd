import api from '../axios';

/**
 * Retrieve readlist for curr user
 */
const getReadList = async () => {
	return await api.get(`/api/read-list`).then((res) => res);
}

/**
 * Add book to read list
 */
const addBookToReadList = async (bookId: number) => {
	return await api.post(`/api/read-list/book/${bookId}`).then((res) => res);
}

/**
 * Add book to read list
 */
const updateReadListItem = async (bookId: number, status: string) => {
	return await api.put(
		`/api/read-list/book/${bookId}`,
		{status}
	).then((res) => res);
}

export default {
	getReadList,
	addBookToReadList,
	updateReadListItem,
}
