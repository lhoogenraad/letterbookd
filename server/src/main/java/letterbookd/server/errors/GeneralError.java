package letterbookd.server.errors;
import org.springframework.http.HttpStatus;

public class GeneralError extends Exception {
	private String message;
	private HttpStatus httpStatus;

	public GeneralError(String message, HttpStatus httpStatus){
		super(message);
		this.message = message;
		this.httpStatus = httpStatus;
	}

	public String getMessage(){
		return this.message;
	}
	
	public HttpStatus getHttpStatus(){
		return this.httpStatus;
	}
}
