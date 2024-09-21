"use client"
// Client side rendered for nooow

import "./navbar.css"
import { IconBook } from '@tabler/icons-react';
import LogoutButton from './logoutButton';
import { Button } from '@mantine/core';
import Link from 'next/link';
import notify from 'util/notify/notify';


export default function Navbar() {
	const notifyNotImplemented = (resourceName: string) => {
		notify.info({message: `${resourceName} not yet implemented, sorry!`});
	}

	return (
		<div className="navbar-container">
			<div className="icon">
			<IconBook/>
			</div>

				<Link href="/dashboard">
			<div className="navbar-title">
					Letterbookd
			</div>
				</Link>

			<div className="navbar-btns">
				<Button className="navbar-btn" variant="light" color="gray" 
					onClick = {() => notifyNotImplemented('My Readlist')}>My Readlist</Button>
				<Button className="navbar-btn" variant="light" color="gray"
					onClick = {() => notifyNotImplemented('My Account')}>My Account</Button>
				<LogoutButton />
			</div>
		</div>
	);
};
