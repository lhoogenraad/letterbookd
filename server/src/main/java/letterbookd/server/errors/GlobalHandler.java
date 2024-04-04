package letterbookd.server;

import letterbookd.server.errors.GeneralError;
import org.springframework.web.bind.annotation.*;
import org.springframework.http.ResponseEntity;
import org.springframework.http.HttpStatus;
import org.springframework.core.Ordered;
import org.springframework.core.annotation.Order;
import java.util.*;

@ControllerAdvice
@Order(Ordered.HIGHEST_PRECEDENCE)
public class GlobalHandler {
	@ExceptionHandler({GeneralError.class})
	protected ResponseEntity<Object> handleGeneralError(GeneralError err) {
		System.out.println("GENERAL GLOBAL ERROR HANDLER CALLED");
		System.err.println(err.getHttpStatus() + ": " + err.getMessage());
		
		return ResponseEntity
			.status(err.getHttpStatus())
			.body(err.getMessage());
	}
}
