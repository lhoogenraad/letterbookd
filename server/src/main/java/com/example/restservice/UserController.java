package com.example.restservice;

import java.util.concurrent.atomic.AtomicLong;
import com.example.models.User;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class UserController {

	private static final String template = "Hello, %s!";
	private final AtomicLong counter = new AtomicLong();

	@GetMapping("/user")
	public User user(@RequestParam(value = "name", defaultValue = "World") User body) {
		System.out.println(body);
		return new User("Leon", "Hoogenraad", "le.o.n@outlook.com", "password");
	}
}
