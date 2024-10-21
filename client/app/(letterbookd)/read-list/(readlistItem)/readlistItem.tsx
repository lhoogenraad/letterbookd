"use client";

import './readlistItem.css';
import { Image, Select } from '@mantine/core';
import Link from 'next/link';

export default function ReadListItem({ item }) {
	return (
		<div className="readlist-item-container">
			<Link
				style={{ textDecoration: "inherit", color: "inherit" }}
				href={`/books/${item.BookId}`}
			>
				<div className="readlist-item-header-container">
					<div className="img">
						<Image
							src={null}
							h='5rem'
							radius="md"
							fallbackSrc="https://placehold.co/600x400?text=Placeholder"
						/>
					</div>
					<div className="title">{item.BookName}</div>
				</div>
			</Link>
			<div className="readlist-item-body-container">
				<div className="readlist-status">{item.Status}</div>
				<div className="readlist-date">01/04/2022</div>
			</div>
		</div>
	)
}
