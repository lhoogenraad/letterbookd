package letterbookd.server;

import java.security.SecureRandom;
import java.security.MessageDigest;

import letterbookd.server.User;
import letterbookd.server.UserModel;

public class UserService {

	public static void createUser(User user) {
		System.out.println("Saving user");

		String passwordHash = "uhh123";
		try{
			UserModel.saveUser(
					user.getFirstName(), 
					user.getLastName(), 
					user.getEmail(), 
					passwordHash);
		}catch(Exception e){
			System.out.println("Error encountered");
			e.printStackTrace();
		}
	}

	/**
	 * Returns hashed version of given password
	 */
	private static String getPasswordHash(String password){
		SecureRandom random = new SecureRandom();
		byte[] salt = new byte[16];
		random.nextBytes(salt);
	}
}
