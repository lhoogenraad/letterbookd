package letterbookd.server;

import java.sql.*;

public class UserModel {

	public static void saveUser(
			String firstName, 
			String lastName, 
			String email, 
			String passwordHash) throws Exception {
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
			System.err.format("SQL State: %s\n%s", e.getSQLState(), e.getMessage());
		} catch (Exception e) {
			e.printStackTrace();
		}
			}
}
