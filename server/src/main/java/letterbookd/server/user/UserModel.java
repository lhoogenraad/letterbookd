package letterbookd.server;

import java.sql.*;
import letterbookd.server.errors.GeneralError;
import org.springframework.http.HttpStatus;

public class UserModel {

	public static void saveUser(
			String firstName, 
			String lastName, 
			String email, 
			String passwordHash) throws GeneralError {
		try{

			Class.forName("com.mysql.cj.jdbc.Driver");

			Connection conn = DriverManager.getConnection(
					"jdbc:mysql://localhost/letterbookd", "root", ""
					);

			String insertStr = """
				INSERT INTO users (first_name, last_name,\
						email, password_hash) values (?, ?, ?, ?) \
				""";
			PreparedStatement insertStmt = conn.prepareStatement(insertStr);

			insertStmt.setString(1, firstName);
			insertStmt.setString(2, lastName);
			insertStmt.setString(3, email);
			insertStmt.setString(4, passwordHash);

			insertStmt.executeUpdate();
		}catch (SQLException e) {
			String errorMessage = e.getMessage();
			System.err.format("SQL State: %s\n%s", e.getSQLState(), errorMessage, "\n");
			if(errorMessage.contains("email_unique")){
				throw new GeneralError(
						"Sorry, a user with that email already exists.",
						HttpStatus.CONFLICT
						);
			}
		} catch (Exception e) {
			throw new GeneralError(
					"Unknown error",
					HttpStatus.INTERNAL_SERVER_ERROR
					);
		}
			}
}
