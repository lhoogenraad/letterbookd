package letterbookd.server;

import letterbookd.server.errors.GeneralError;

import java.security.SecureRandom;
import java.security.MessageDigest;

import letterbookd.server.User;
import letterbookd.server.UserModel;

public class UserService {
	public static void createUser(User user) throws GeneralError{
		System.out.println("Saving user");

		String passwordHash = getPasswordHash(user.getPassword());
		UserModel.saveUser(
				user.getFirstName(), 
				user.getLastName(), 
				user.getEmail(), 
				passwordHash);
	}

	/**
	 * Returns hashed version of given password
	 */
	private static String getPasswordHash(String password){
		SecureRandom random = new SecureRandom();
		byte[] salt = new byte[16];
		random.nextBytes(salt);

		return "uh oh";
	}
}
