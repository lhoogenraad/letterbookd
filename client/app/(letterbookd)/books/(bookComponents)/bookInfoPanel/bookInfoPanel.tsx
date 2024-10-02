import './bookInfoPanel.css';

import { useDisclosure } from '@mantine/hooks';
import { Image, Button, Modal } from '@mantine/core';
import { IconPencilPlus } from '@tabler/icons-react';
import BookReviewAverage from './bookReviewAverage';
import AddBookToReadlistButton from './addBookToReadlistButton';
import AddReview from 'components/reviews/addReview';

export default function BookInfoPanel({ book, reload}) {
	const [opened, { open, close }] = useDisclosure(false);

	const closeModal = () => {
		console.log('close modal called')
		close();
	}

	if (!book) {
		return (
			<div className="book-info-container">
				<div>Rendering. . .</div>
			</div>
		)
	}

	return (
		<div className="book-info-container">
			<div>

				<div className="book-info-item">
					<span className="book-info-subheader">
						Title
					</span>
					<h2 className="book-title">{book.Title}</h2>
				</div>


				<div className="book-info-item">
					<span className="book-info-subheader">
						Author
					</span>
					<div className="book-info-item-content">
						{book.Author}
					</div>
				</div>

				<div className="book-info-item">
					<span className="book-info-subheader">
						Date Published
					</span>
					<div className="book-info-item-content">
						{new Date(book.Published).toLocaleDateString()}
					</div>
				</div>

				<div className="book-info-item">
					<span className="book-info-subheader">
						Synopsis
					</span>
					<div className="book-info-item-content">{book.Synopsis}</div>
				</div>

				<div className="book-info-item">
					<BookReviewAverage bookId={book.Id}/>
				</div>
			</div>

			<div className="book-info-img">
				<Image 
					className='book-tile-img'
					src={null}
					h='auto'
					w='auto'
					radius="md"
					fallbackSrc="https://placehold.co/600x400?text=Placeholder"
				/>

				<div className="book-info-actions-container">
					<AddBookToReadlistButton bookId={book.Id} />
					<Button onClick={open} rightSection={<IconPencilPlus size={20}/>} fullWidth>Create Review</Button>
				</div>
			</div>


		  <Modal 
		  	opened={opened} 
			onClose={close} 
			title={book.Title}
			centered
			size="85%"
			transitionProps={{ transition: 'slide-down' }}
			>
		  		<AddReview book={book} reload={reload} closeModal={closeModal}/>
		  </Modal>
		</div>
	)
}
