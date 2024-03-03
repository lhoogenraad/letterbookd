package letterbookd.server;

import java.util.concurrent.atomic.AtomicLong;

import org.springframework.web.bind.annotation.*;

import letterbookd.server.User;

@RestController
public class UserController {

	@PostMapping("/signup")
	public User user(@RequestBody User body) {
		System.out.println(body);
		return body;
	}

}
