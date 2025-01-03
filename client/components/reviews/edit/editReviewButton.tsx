"use client"

import { IconEdit } from '@tabler/icons-react'
import { useDisclosure } from '@mantine/hooks';
import { Button } from '@mantine/core';
import { Modal } from '@mantine/core';
import EditReviewModal from './editReviewModal';
import '../addReview.css';

export default function EditReviewButton({ review, book, updateReview}) {
	const [opened, { open, close }] = useDisclosure(false);

	return (
		<div>
			<Button variant="subtle" onClick={open}>Edit</Button>
			<Modal
				opened={opened}
				onClose={close}
				title={book.Title}
				centered
				size="85%"
				transitionProps={{ transition: 'slide-down' }}
			>
				<EditReviewModal 
					book={book} 
					review={review} 
					closeModal={close} 
					updateReview={updateReview}
				/>
			</Modal>
		</div>
	)
}
