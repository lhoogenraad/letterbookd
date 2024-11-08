import api from '../axios';

/**
 * Retrieve all books available
 */
const searchForBookOL = async ({ title, author, publisher }) => {
	const url = buildSearchOLURL({ title, author, publisher })
	console.log(url)
	// return await api.get(url).then((res) => res);
};

const buildSearchOLURL = ({ title, author, publisher }) => {
	let url = `/books/search/open-lbrary`
	let f = '?'
	
	if (title) {
		url += `${f}title=${title}`;
		f = '&'
	}
	
	if (author) {
		url += `${f}author=${author}`;
		f = '&'
	}
	
	if (publisher) {
		url += `${f}publisher=${publisher}`;
		f = '&'
	}
	url = encodeURI(url);
	return url
}

export default {
	searchForBookOL,
}
