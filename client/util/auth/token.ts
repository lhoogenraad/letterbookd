import cookie from 'js-cookie'; 

const saveToken = (token: string) => {
	cookie.set('authToken', token, { expires: 1, secure: true, sameSite: 'strict' });
}

export default {
	saveToken,
}
