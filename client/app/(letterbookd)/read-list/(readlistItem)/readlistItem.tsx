import './readlistItem.css';
import { Image } from '@mantine/core';

export default function ReadListItem({item}) {
	return (
		<div className="readlist-item-container">
			<div className="readlist-header-container">
				<div className="img">
					<Image 
						src={null}
						h='auto'
						w='auto'
						radius="md"
						fallbackSrc="https://placehold.co/600x400?text=Placeholder"
					/>
				</div>
				<div className="title">{item.BookName}</div>
			</div>
			<div className="readlist-item-body-container">
				<div className=""></div>
				<div className=""></div>
			</div>
		</div>
	)
}
