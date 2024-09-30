import { Carousel } from '@mantine/carousel';
import Autoplay from 'embla-carousel-autoplay';
import { useRef } from 'react';

export default function BookList({books}) {
	const autoplay = useRef(Autoplay({ delay: 5000 }));


	return (
		<div>
		<Carousel 
		loop
		plugins={[autoplay.current]}
		onMouseEnter={autoplay.current.stop}
		onMouseLeave={autoplay.current.reset}
		slideSize="50%" 
		draggable={false}>
		{books.map((book: any) => (
			<Carousel.Slide
			className='book-list-item' 
			key={book.Title}>
				<div className='book-title'>
				{book.Title}
				</div>
				<div className='book-author'>
				{book.Author}
				</div>
				<div className='book-published-date'>
				{new Date(book.Published).toLocaleDateString()}
				</div>
			</Carousel.Slide>
		))}
		</Carousel>
		</div>
	)
}
