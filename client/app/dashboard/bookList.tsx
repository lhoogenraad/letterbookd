

export default function BookList({books}) {
	return (
		<div>
		{books.map((book: any) => (
			<div className='book-list-item' key={book.Title}>
				<div className='book-title'>
				{book.Title}
				</div>
				<div className='book-author'>
				{book.Author}
				</div>
				<div className='book-published-date'>
				{new Date(book.Published).toLocaleDateString()}
				</div>
			</div>
		))}
		</div>
	)
}
