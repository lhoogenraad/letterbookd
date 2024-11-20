"use client";

import { IconHeartFilled, IconHeart } from '@tabler/icons-react';
import { useState } from 'react';
import api from 'util/api/api';
import './reviewHeart.css';

export default function ReviewHeart({ reviewId, liked, updateLikedCount}) {
	const [loading, setLoading] = useState(false);
	const [likedState, setLiked] = useState(liked);

	const sendLike = async (status: boolean) => {
		// Is this a bad idea?
		if (loading) return;
		setLoading(true);
		await api.reviews.setLikeStatus(reviewId, status)
			.then(() => {
				setLiked(status)
				updateLikedCount(status)
			})
			.catch((err) => console.error(err))
			.finally(() => setLoading(false));

	}

	const likedHeart = (
		<IconHeartFilled
			className="review-heart"
			color="red" 
			onClick={() => sendLike(false)} 
		/>
	)

	const unlikedHeart = (
		<IconHeart 
			className="review-heart"
			color="red" 
			onClick={() => sendLike(true)} 
		/>
	)

	if (likedState === true) {
		return likedHeart;
	} else {
		return unlikedHeart;
	}
}
