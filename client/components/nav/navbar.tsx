"use client"
// Client side rendered for nooow

import "./navbar.css"
import { IconBook } from '@tabler/icons-react';
import LogoutButton from './logoutButton.tsx';
import { Button } from '@mantine/core';
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

			<div className="navbar-title">
				Letterbookd
			</div>

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
