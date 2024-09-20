import "./navbar.css"
import { IconBook } from '@tabler/icons-react';
import LogoutButton from './logoutButton.tsx';


export default function Navbar() {
	return (
		<div className="navbar-container">
			<div className="icon">
			<IconBook/>
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
