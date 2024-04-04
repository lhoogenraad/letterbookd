package letterbookd.server;

import letterbookd.server.errors.GeneralError;
import java.util.concurrent.atomic.AtomicLong;

import org.springframework.web.bind.annotation.*;

import letterbookd.server.User;
import letterbookd.server.UserService;

@RestController
public class UserController {

	@PostMapping("/signup")
	public User user(@RequestBody User body) throws GeneralError {
		UserService.createUser(body);
		return body;
	}

}
