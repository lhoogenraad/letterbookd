package letterbookd.server;

import letterbookd.server.errors.GeneralError;
import letterbookd.server.User;
import io.jsonwebtoken.*;
import jakarta.servlet.http.HttpServletRequest;
import org.springframework.stereotype.Component;
import org.springframework.http.HttpStatus;

import java.util.Date;
import java.util.List;
import java.util.concurrent.TimeUnit;

@Component
public class JwtUtil {


	private final String secret_key = "mysecretkey";
	private long accessTokenValidity = 60*60*1000;

	private final JwtParser jwtParser;

	private final String TOKEN_HEADER = "Authorization";
	private final String TOKEN_PREFIX = "Bearer ";

	public JwtUtil(){
		this.jwtParser = Jwts.parser().setSigningKey(secret_key);
	}

	public String createToken(User user) {
		Claims claims = Jwts.claims().setSubject(user.getEmail());
		claims.put("firstName",user.getFirstName());
		claims.put("lastName",user.getLastName());
		Date tokenCreateTime = new Date();
		Date tokenValidity = new Date(tokenCreateTime.getTime() + TimeUnit.MINUTES.toMillis(accessTokenValidity));
		return Jwts.builder()
			.setClaims(claims)
			.setExpiration(tokenValidity)
			.signWith(SignatureAlgorithm.HS256, secret_key)
			.compact();
	}

	private Claims parseJwtClaims(String token) {
		return jwtParser.parseClaimsJws(token).getBody();
	}

	public Claims resolveClaims(HttpServletRequest req) throws GeneralError{
		try {
			String token = resolveToken(req);
			if (token != null) {
				return parseJwtClaims(token);
			}
			return null;
		} catch (Exception ex){
			req.setAttribute("invalid", ex.getMessage());
			throw new GeneralError("Unknown token parsing error", HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}

	public String resolveToken(HttpServletRequest request) {

		String bearerToken = request.getHeader(TOKEN_HEADER);
		if (bearerToken != null && bearerToken.startsWith(TOKEN_PREFIX)) {
			return bearerToken.substring(TOKEN_PREFIX.length());
		}
		return null;
	}

	public boolean validateClaims(Claims claims) throws GeneralError {
		try {
			return claims.getExpiration().after(new Date());
		} catch (Exception e) {
			throw new GeneralError("Token is invalid", HttpStatus.UNAUTHORIZED);
		}
	}

	public String getEmail(Claims claims) {
		return claims.getSubject();
	}
}
