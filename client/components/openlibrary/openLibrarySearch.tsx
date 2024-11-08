"use client"

import { Button, TextInput } from '@mantine/core';
import { useForm } from '@mantine/form';


export default function OpenLibrarySearch({ submitSearch, loading }: { submitSearch: Function, loading: boolean }) {
	const form = useForm({
		mode: 'uncontrolled',
		initialValues: {
			title: '',
			author: '',
			publisher: '',
		},
	});

	const formSubmit = (values: any) => {
		const { title, author, publisher } = values;
		submitSearch({ title, author, publisher })
	}

	return (
		<div>
			<form onSubmit={formSubmit}>
				<TextInput
					label="Book title"
					placeholder="I, claudius..."
					key={form.key("title")}
					{...form.getInputProps("title")}
				/>
				<TextInput
					label="Author name"
					placeholder="Robet Graves"
					key={form.key("author")}
					{...form.getInputProps("author")}
				/>
				<TextInput
					label="Publisher"
					placeholder="Penguin books"
					key={form.key("publisher")}
					{...form.getInputProps("publisher")}
				/>
				<Button loading={loading} disabled={loading} type="submit">Submit</Button>
			</form>
		</div>
	)

};
