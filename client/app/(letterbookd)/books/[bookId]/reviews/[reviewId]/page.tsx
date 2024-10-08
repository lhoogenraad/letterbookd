"use client";

import api from 'util/api/api';
import notify from 'util/notify/notify';
import { useState, useEffect } from 'react';
import ReviewComment from './(reviewComments)/reviewComment';
import './singleReview.css';

export default function Review({ params }: { bookId: number, reviewId: number }) {
	const [book, setBook] = useState(null);
	const [loading, setLoading] = useState(false);
	const [review, setReview] = useState(null);
	const [comments, setReviewComments] = useState(null);

	const loadBook = async () => {
		await api.books.getBook(params.bookId)
			.then((res) => setBook(res.data))
	};

	// const loadReview = async () => {
	// 	await api.books.getBook(params.bookId)
	// 	.then((res) => setReviewComments(res.data))
	// };

	const loadComments = async () => {
		await api.reviews.getReviewComments(params.reviewId)
			.then((res) => setReviewComments(res.data))
	};

	const init = async () => {
		setLoading(true);

		try {
			const initPromises = [loadComments(), loadBook()];
			await Promise.all(initPromises);
		} catch (err) {
			console.error(err);
			notify.error({
				message:
					err?.response?.data?.Message ??
					`Something went wrong. Please try again later.`,
			});
		} finally {
			setLoading(false);
		}
	}

	useEffect(() => {
		init();
	}, []);

	if (loading || !book || !comments) {
		return <div>Loading..</div>
	}

	return (
		<div>
			<div className="comments-container">
				<div className="comment-container">
				<div className="comment-separator" />
					{comments.map((comment: any, index: number) => (
						<div>
							<ReviewComment comment={comment} key={index} />
							<div className="comment-separator" />
						</div>
					))}
				</div>
			</div>
			<div>bookId: {params.bookId}</div>
			<div>reviewId: {params.reviewId}</div>
		</div>
	)
}
