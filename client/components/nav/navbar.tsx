"use client"

import "./navbar.css"
import { IconBooks } from '@tabler/icons-react';
import LogoutButton from './logoutButton';
import { Button } from '@mantine/core';
import Link from 'next/link';
import { useRouter } from "next/navigation";


export default function Navbar() {
	const router = useRouter();

	const goTo = (url:string) => {
		router.push(url);
	}

	return (
		<div className="navbar-container">
			<Link style={{textDecoration: "inherit", color: "inherit"}}  href="/dashboard">
				<div className="icon">
				<IconBooks size={38} />
					<div className="navbar-title">
							Shelfd
					</div>
				</div>
			</Link>

			<div className="navbar-btns">
				<Button 
					variant="subtle"
					className="navbar-btn" 
					color="primary.0"
					onClick = {() => goTo('/read-list')}>My Readlist</Button>
				<Button 
					variant="subtle"
					className="navbar-btn" 
					color="primary.0"
					onClick = {() => goTo('/books')} >Search Books</Button>
				<LogoutButton />
			</div>
		</div>
	);
};
