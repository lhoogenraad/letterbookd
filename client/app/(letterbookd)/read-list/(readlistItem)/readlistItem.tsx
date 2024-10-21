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

export default function ReadListItem({ item }) {
	const [status, setStatus] = useState(item.Status)
	const [loading, setLoading] = useState(false);
	const [opened, { open, close }] = useDisclosure(false);

	const closeModal = () => {
		close();
	}

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
						src={null}
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
				<Select
					variant='unstyled'
					className="readlist-status"
					data={statusOptions}
					rightSection={loading ? <Loader size='xs' /> : null}
					value={status}
					disabled={loading}
					onChange={async (val) => await updateStatus(val)}
				/>
				{
					status == 'Read' ?
						<Button onClick={open}>Create review</Button>
						:
						null
				}
				<div className="readlist-date">01/04/2022</div>
		  <Modal 
		  	opened={opened} 
			onClose={close} 
			title={item.BookName}
			centered
			size="85%"
			transitionProps={{ transition: 'slide-down' }}
			>
		  		<AddReview book={{Id: item.BookId}} reload={undefined} closeModal={closeModal}/>
		  </Modal>
			</div>
		</div>
	)
}
