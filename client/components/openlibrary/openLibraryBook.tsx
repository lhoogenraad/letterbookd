"use client"

import { Button, Image } from "@mantine/core"
import "./openLibraryBook.css";

export default function OpenLibraryBook({ olBook, submitAdd, loading }) {
	return (
		<div>
			<div className="ol-book-container">
				<div className="attributes-container">
					<div className="book-attribute">
						<div className="attribute-subtitle">Title</div>
						<div>{olBook.Title}</div>
					</div>

					<div className="book-attribute">
						<div className="attribute-subtitle">Author</div>
						<div>{olBook.Author}</div>
					</div>

					<div className="book-attribute">
						<div className="attribute-subtitle">Published</div>
						<div>{new Date(olBook.Published).toLocaleDateString()}</div>
					</div>

					<div className="book-attribute">
						<div className="attribute-subtitle">Synopsis</div>
						<div>{olBook.Synopsis == "" ? "No synopsis available" : olBook.Synopsis}</div>
					</div>

				</div>

				<Image
					src={olBook.CoverURL}
					h='300'
					w='200'
					fallbackSrc="https://placehold.co/600x400?text=Placeholder"
				/>
			</div>
			<Button
				color="secondary.6"
				variant="filled"
				loading={loading}
				onClick={() => submitAdd(olBook.OlID)}>
				Attempt to add
			</Button>
		</div>
	)
}
