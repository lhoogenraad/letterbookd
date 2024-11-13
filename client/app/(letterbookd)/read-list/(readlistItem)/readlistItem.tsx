"use client";

import './readlistItem.css';
import { Image, Select, Loader, Modal, Button } from '@mantine/core';
import Link from 'next/link';
import api from 'util/api/api';
import notify from 'util/notify/notify';
import statusOptions from 'configs/readlistStatusOptions';
import AddReview from 'components/reviews/addReview';
import { useDisclosure } from '@mantine/hooks';
import { useState } from 'react';
import DeleteReadlistItem from './deleteReadlistItem';

export default function ReadListItem({ item, removeItem }) {
	const [status, setStatus] = useState(item.Status)
	const [loading, setLoading] = useState(false);
	const [createReviewOpened, setCreateReviewModal] = useDisclosure(false);
	const [deleteReadlistItemOpened, setDeleteReadlistItemModal] = useDisclosure(false);

	const updateStatus = async (newStatus: string) => {
		// If status value is same as current status, don't send request
		if (newStatus === status) return;
		setLoading(true);
		await api.readlist.updateReadListItem(item.BookId, newStatus)
			.then(() => setStatus(newStatus))
			.catch((err) => notify.error({ message: err?.response?.data?.Message }))
			.finally(() => setLoading(false))
	}

	return (
		<div className="readlist-item-container">
			<div className="readlist-item-header-container">
				<div className="img">
					<Image
						src={item.CoverURL}
						h='5rem'
						radius="md"
						fallbackSrc="https://placehold.co/600x400?text=Placeholder"
					/>
				</div>
				<Link
					style={{ textDecoration: "inherit", color: "inherit" }}
					href={`/books/${item.BookId}`}
				>
					<div className="title">{item.BookName}</div>
				</Link>
			</div>
			<div className="readlist-item-body-container">
				<div style={{display: 'flex', flexDirection: 'row', gap:'2rem'}}>
					<Select
						variant='unstyled'
						className="readlist-status"
						data={statusOptions}
						rightSection={loading ? <Loader size='xs' /> : null}
						value={status}
						disabled={loading}
						onChange={async (val) => await updateStatus(val)}
					/>
					<div className="readlist-date">
						<div className='readlist-date-title'>Date added:</div>
						<div>01/04/2022 <span style={{ fontSize: '0.6rem' }}>(placeholder)</span></div>
					</div>
				</div>
				<div>
				{
					status == 'Read' && !item.UserReviewed ?
						<Button variant='transparent' onClick={setCreateReviewModal.open}>Create review</Button>
						:
						null
				}
				<Button variant='subtle' color="red" onClick={setDeleteReadlistItemModal.open}>Remove book</Button>
				</div>

				<Modal
					opened={createReviewOpened}
					onClose={setCreateReviewModal.close}
					title={item.BookName}
					centered
					size="85%"
					transitionProps={{ transition: 'slide-down' }}
				>
					<AddReview 
						book={{ Id: item.BookId }} 
						reload={undefined} 
						closeModal={setCreateReviewModal.close} 
					/>
				</Modal>


				<Modal
					opened={deleteReadlistItemOpened}
					onClose={setDeleteReadlistItemModal.close}
					title={`Are you sure you want to remove ${item.BookName} from your read list?`}
					centered
					size="40%"
					transitionProps={{ transition: 'slide-down' }}
				>
					<DeleteReadlistItem 
						bookId={item.BookId}
						removeItem={removeItem} 
						closeModal={setDeleteReadlistItemModal.close} 
					/>
				</Modal>
			</div>
		</div>
	)
}
