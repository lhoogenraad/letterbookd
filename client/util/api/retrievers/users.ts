import api from '../axios';


/**
 * Post sign in attempt
 * @param {String} body.email user email
 * @param {String} body.password user password
 */
const signIn = async (body: object) => {
	return await api.post(`/users/login`, body).then((res) => res);
};

export default {
	signIn,
}
