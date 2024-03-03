package letterbookd.server;

import java.sql.*;

public class UserModel {
	private Connection conn = null;

	public void saveUser(String firstName, String lastName, String email, String passwordHash) throws Exception {
		try{
			Class.forName("com.mysql.jdbc.Driver");
		}catch(Exception e){
			throw e;
		}finally{
			close();
		}
	}
}
