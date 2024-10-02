'use client';


import { Rating } from '@mantine/core';

export default function AddReview({book}) {
	return (
		<div className="review-container">
			Hey 
			{JSON.stringify(book)}
		</div>
	)
}
