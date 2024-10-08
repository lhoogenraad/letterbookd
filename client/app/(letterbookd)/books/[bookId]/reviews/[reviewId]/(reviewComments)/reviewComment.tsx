import './reviewComments.css';
import { Avatar } from '@mantine/core';
export default function ReviewComment({comment}){
	return (
		<div className="single-comment-container">
			<div className="comment-username">
				<Avatar />
				{comment.Username}
			</div>
			<div className="comment-content">
				{comment.Comment}
			</div>
		</div>
	)
};
