package letterbookd.server.errors;

public class GeneralError extends Exception {
	private int status;
	private String message;
	public GeneralError(String message, int status){
		super(message);
		this.message = message;
		this.status = status;
	}

	public String getMessage(){
		return this.message;
	}

	public int getStatus(){
		return this.status;
	}
}
