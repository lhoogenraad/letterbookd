"use client";

import api from 'util/api/api';
import notify from 'util/notify/notify';
import { useState, useEffect } from 'react';
import ReviewComment from './(reviewComments)/reviewComment';
import AddComment from './(reviewComments)/addComment';
import ReviewHeart from 'components/reviews/likes/reviewHeart';
import '../../../(bookComponents)/bookReviewList/review.css';
import { Rating } from '@mantine/core';
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

	const loadReview = async () => {
		await api.reviews.getBookReviews(params.bookId)
			.then((res) => {
				const findReview = res.data.find((item) => {
					return item.Id == params.reviewId;
				})
				setReview(findReview)
			})
	};

	const loadComments = async () => {
		await api.reviews.getReviewComments(params.reviewId)
			.then((res) => { 
				const sortedComments = sortComments(res.data ?? []);
				setReviewComments(sortedComments);
			})
	};

	const sortComments = (comments) => {
		return comments.sort((c1, c2) => {
			const d1 = new Date(c1.Timestamp);
			const d2 = new Date(c2.Timestamp);
			return d1 < d2;
		})
	}


	const init = async () => {
		setLoading(true);

		try {
			const initPromises = [loadComments(), loadBook(), loadReview()];
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
		<div className="container">
			<div className="review-container">
				<div className="review-book-title">
					{book.Title}
				</div>
				<div className="review-metadata-container">
					<div className="review-user">{review.Username}</div>
					<div className="vertical-seperator">|</div>
					<Rating readOnly fractions={2} value={review.Rating / 2} className="review-value" />
					<div className="vertical-seperator">|</div>
					<ReviewHeart 
						reviewId={review.Id} 
						liked={review.LikedBy} 
						numLikes={review.NumLikes}
					/>
				</div>
				<div className="review-description">{review.Description}</div>
			</div>

			<div className="comments-container">
				<div className="comment-container">
					<div style={{ marginBottom: '10px' }}>
						{comments.length} Comments
					</div>
					<div className="comment-separator" />
					{comments.map((comment: any, index: number) => (
						<div>
							<ReviewComment comment={comment} key={index} reloadComments={loadComments} />
							<div className="comment-separator" />
						</div>
					))}
					<div>
						<AddComment reload={loadComments} reviewId={review.Id} />
					</div>
				</div>
			</div>
		</div>
	)
}
