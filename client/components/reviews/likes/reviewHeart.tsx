"use client";

import { IconHeartFilled, IconHeart } from '@tabler/icons-react';
import { useState } from 'react';
import api from 'util/api/api';
import './reviewHeart.css';

export default function ReviewHeart({ reviewId, liked, numLikes }) {
	const [loading, setLoading] = useState(false);
	const [likedState, setLiked] = useState(liked);
	const [likes, setLikes] = useState(numLikes);

	const updateLikedCount = (status) => {
		if (status === true) {
			setLikes(likes + 1)
		} else {
			setLikes(likes - 1)
		}
	}

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
		return (
			<div className="heart-container">
				{likedHeart}
				<span>{likes}</span>
			</div>
		);
	} else {
		return (
			<div className="heart-container">
				{unlikedHeart}
				<span>{likes}</span>
			</div>
		);
	}
}
