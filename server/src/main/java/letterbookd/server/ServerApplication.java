package letterbookd.server;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class ServerApplication {
	private static String url = "http://localhost:8080";
	private static String startMsg = "Letterbookd server now running on ";


	public static void main(String[] args) {
		System.out.println(startMsg + url);

		SpringApplication.run(ServerApplication.class, args);
	}
}
