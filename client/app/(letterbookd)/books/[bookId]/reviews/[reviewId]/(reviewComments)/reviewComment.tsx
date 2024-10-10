"use client";

import './reviewComments.css';
import { Avatar } from '@mantine/core';
import { Button, Modal } from '@mantine/core';
import { useDisclosure } from '@mantine/hooks';
import EditCommentModal from './reviewActionModals/editComment';
import DeleteCommentModal from './reviewActionModals/deleteComment';

export default function ReviewComment({ comment, reloadComments }) {
	const [showDelete, setShowDelete] = useDisclosure(false);
	const [showEdit, setShowEdit] = useDisclosure(false);

	const commentDate = new Date(comment.Timestamp);
	const displayDateTimeString =
		commentDate.toLocaleString('en-NZ', {
			year: 'numeric',
			month: 'numeric',
			day: 'numeric',
			
			hour: 'numeric',
			hour12: true,
			minute: 'numeric',

		});

	return (
		<div className="single-comment-container">
			<div className="comment-content-container">
				<div className={comment.OwnedBy ? "comment-username owner" : "comment-username"}>
					<Avatar />
					{comment.Username}
				</div>
				<div className="comment-content">
					{comment.Comment}
				</div>

			</div>
			<div className="comment-actions-container">
				<div className="comment-timestamp">{displayDateTimeString}</div>
				{comment.OwnedBy ?
					<div>
						<Button onClick={() => setShowEdit.open()} variant="subtle" color="blue">Edit</Button>
						<Button onClick={() => setShowDelete.open()} variant="subtle" color="red">Delete</Button>
					</div>
					: null
				}
			</div>

			<Modal
				opened={showDelete}
				onClose={setShowDelete.close}
				title="Delete comment?"
				centered
				size="50%"
				transitionProps={{ transition: 'slide-down' }}
			>
				<DeleteCommentModal
					reloadComments={reloadComments}
					comment={comment}
					closeModal={setShowDelete.close}
				/>
			</Modal>


			<Modal
				opened={showEdit}
				onClose={setShowEdit.close}
				title="Edit comment"
				centered
				size="50%"
				transitionProps={{ transition: 'slide-down' }}
			>
				<EditCommentModal
					reloadComments={reloadComments}
					comment={comment}
					closeModal={setShowEdit.close}
				/>
			</Modal>
		</div>
	)
};
