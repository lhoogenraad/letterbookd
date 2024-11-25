"use client";

import Link from 'next/link';
import './footer.css';
import { EmailIcon, GithubIcon, LinkedinIcon } from "next-share";
import notify from 'util/notify/notify';

export default function Footer() {
	const leonEmail = `le.o.n@outlook.com`;

	const copyEmail = () => {
		navigator.clipboard.writeText(leonEmail);
		notify.info({message: `Copied Leon's email to clipboard!`});
	}

	return (
		<div className="footer-container">
			<div className="primary-content">
				This application was made as a passion project by Leon Hoogenraad
				<div className="social-links">
					<Link href="https://www.linkedin.com/in/leon-hoogenraad-444b521ab">
						<LinkedinIcon size={26} round />
					</Link>
					<Link href="https://github.com/lhoogenraad">
						<GithubIcon size={26} round />
					</Link>
					<div onClick={copyEmail} className="copy-button">
						<EmailIcon size={26} round />
					</div>
				</div>
			</div>
			<div className="secondary-content">
				©️ Letterbookd
			</div>
		</div>
	);
};
