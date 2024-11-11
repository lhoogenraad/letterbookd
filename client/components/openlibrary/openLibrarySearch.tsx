"use client"

import { Button, TextInput } from '@mantine/core';
import { useForm } from '@mantine/form';

import './openLibrarySearch.css';


export default function OpenLibrarySearchForm({ submitSearch, loading }: { submitSearch: Function, loading: boolean }) {
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
		<div className="open-library-form-container">
			<form className="open-library-form" onSubmit={form.onSubmit((vals: any) => formSubmit(vals))}>
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
				<div className="submit-btn">
					<Button loading={loading} disabled={loading} type="submit">Submit</Button>
				</div>
			</form>
		</div>
	)

};
