package letterbookd.server;

public record User(
		String firstName,
		String lastName,
		String email,
		String password) { }
