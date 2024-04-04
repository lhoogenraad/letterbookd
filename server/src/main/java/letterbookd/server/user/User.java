package letterbookd.server;

public class User {
	private String firstName;
	private String lastName;
	private String email;
	private String password;

	public User (String firstName, String lastName, String email, String password) {
		this.firstName = firstName;
		this.lastName = lastName;
		this.email = email;
		this.password = password;
	}


	public String getFirstName(){
		return this.firstName;
	}


	public String getLastName(){
		return this.lastName;
	}


	public String getEmail(){
		return this.email;
	}

	public String getPassword(){
		return this.password;
	}
}
