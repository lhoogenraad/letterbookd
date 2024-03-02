
package com.example.models;

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

	public String toString() {
		return this.firstName + " " + this.lastName;
	}
}
