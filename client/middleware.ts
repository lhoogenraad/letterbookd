import { NextRequest, NextResponse } from 'next/server';
import { jwtVerify } from 'jose';
import safeRoutes from 'util/auth/safeRoutes';

const SECRET: Uint8Array = new TextEncoder().encode(process.env.JWT_SECRET || "secretpassword");

const redirectToLogin = (request: NextRequest) : NextResponse | undefined => {
	const loginUrl = new URL('/login', request.url);
    return NextResponse.redirect(loginUrl);
};


const safeRoute = (pathname: string) : Boolean => {
	console.log({pathname})
	return safeRoutes.some((route:string) => pathname.startsWith(route));
};


export async function middleware (request: NextRequest): Promise<NextResponse> {
	const pathname = request.nextUrl.pathname;
	if (safeRoute(request.nextUrl.pathname)) {
		return NextResponse.next();
	}
	const token = request.cookies.get('authToken')?.value;

	try {
		await jwtVerify(token, SECRET);
		return NextResponse.next();
	} catch (error) {
		console.error(`Failed to verify token when navigating to ${pathname}`);
		console.error(error.message, '\n\n');
		return redirectToLogin(request);
	}
};


// Config: Apply the middleware only to specific routes
export const config = {
  matcher: ['/book/:path*', '/read-list/:path*', '/dashboard/:path*'],
};
