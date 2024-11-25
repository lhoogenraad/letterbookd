"use client"

import "./navbar.css"
import { useRouter } from 'next/navigation'
import { Button } from '@mantine/core';
import cookies from 'util/auth/token';

export default function LogoutButton() {
	const router = useRouter();
	/**
	 * Clear auth token localStorage and redirect
	 * user to the login page.
	 */
	const logout = () => {
		localStorage.removeItem("authToken");
		cookies.deleteCookie("authToken");
		router.push('/login');
	}

	return (
		<Button 
			className="logout-btn"
			variant="subtle"
			color="primary.0"
			onClick={() => logout()}
		>
		Logout
		</Button>
	)
};
