"use client"

import { Button } from "@mantine/core";
import Link from "next/link";
import "./not-found.css";

export default function Custom404() {

	return (
		<div className="not-found-container">
			<div className="not-found-text">
				<div>404</div>
				<div className="divider" />
				<div>Sorry, we couldn't find the page you're looking for</div>
			</div>
			<div className="buttons">
				<Link href="/dashboard">
					<Button variant='default' color="secondary.9">Dashboard</Button>
				</Link>
			</div>
		</div>
	)
}
