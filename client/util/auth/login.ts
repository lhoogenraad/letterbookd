import auth from 'util/auth/token';

const saveToken = (authToken: string) => {
	localStorage.setItem("authToken", authToken);
	auth.saveToken(authToken);
	console.info("Saving token to cookies")
}

export default {
	saveToken,
}
