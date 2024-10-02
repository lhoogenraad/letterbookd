"use client"

import "./navbar.css"
import { IconBook } from '@tabler/icons-react';
import LogoutButton from './logoutButton';
import { Button } from '@mantine/core';
import Link from 'next/link';
import notify from 'util/notify/notify';
import { useRouter } from "next/navigation";


export default function Navbar() {
	const router = useRouter();
	const notifyNotImplemented = (resourceName: string) => {
		notify.info({message: `${resourceName} not yet implemented, sorry!`});
	}

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
					onClick = {() => notifyNotImplemented('My Readlist')}>My Readlist</Button>
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
