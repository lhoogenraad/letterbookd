"use client"

import { usePathname } from 'next/navigation'
import "./navbar.css"
import hideNavbarRoutes from './routesToHideNavbar.tsx';
import LogoutButton from './logoutButton.tsx';

export default function Navbar() {
	const pathname = usePathname();

	if (hideNavbarRoutes.includes(pathname)) {
		return;
	}

	return (
		<div className="navbar-container">
			<div className="icon">
				icon
			</div>

			<div className="navbar-title">
				Letterbookd
			</div>

			<div className="navbar-btns">
				<div>btn1</div>
				<div>btn2</div>
				<LogoutButton />
			</div>
		</div>
	);
};
