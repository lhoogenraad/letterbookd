"use client"

import { usePathname } from 'next/navigation'
import "./navbar.css"
import { useRouter } from 'next/navigation'
import { Button } from '@mantine/core';

export default function LogoutButton() {
	const router = useRouter();
	/**
	 * Clear auth token localStorage and redirect
	 * user to the login page.
	 */
	const logout = () => {
		localStorage.removeItem("authToken");
		router.push('/login');
	}

	return (
		<Button 
			variant="filled"
			onClick={() => logout()}
		>
		Logout
		</Button>
	)
};
