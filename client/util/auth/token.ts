import cookie from 'js-cookie'; 

const saveToken = (token: string) => {
	cookie.set('authToken', token, { expires: 1, secure: true, sameSite: 'strict' });
}

const deleteCookie = (cookieName: string) => {
	cookie.remove(cookieName, {path: ''});
}

export default {
	saveToken,
	deleteCookie,
}
