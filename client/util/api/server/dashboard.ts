import api from './axios';

const getFeaturedBooks = async () => {
	return await api.get(`/books/featured`).then((res) => res);
};

const getPopularReviews = async () => {
	return await api.get(`/reviews/popular`).then((res) => res);
};

export default {
	getFeaturedBooks,
	getPopularReviews,
};
