import './reviewComments.css';
export default function ReviewComment({comment}){
	return (
		<div className="single-comment-container">
			<div className="comment-username">
				{comment.Username}
			</div>
			<div className="comment-content">
				{comment.Comment}
			</div>
		</div>
	)
};
