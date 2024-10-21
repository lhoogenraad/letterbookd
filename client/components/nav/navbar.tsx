"use client"

import "./navbar.css"
import { IconBook } from '@tabler/icons-react';
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
			<div className="icon">
			<IconBook/>
			</div>
			<Link style={{textDecoration: "inherit", color: "inherit"}}  href="/dashboard">
				<div className="navbar-title">
						Letterbookd
				</div>
			</Link>

			<div className="navbar-btns">
				<Button 
					className="navbar-btn" 
					variant="gradient"
					gradient={{from: 'black', to: 'darkblue', deg: 30}}
					onClick = {() => goTo('/read-list')}>My Readlist</Button>
				<Button 
					className="navbar-btn" 
					variant="gradient"
					gradient={{from: 'black', to: 'darkred', deg: 250}}
					onClick = {() => goTo('/books')} >Search Books</Button>
				<LogoutButton />
			</div>
		</div>
	);
};
