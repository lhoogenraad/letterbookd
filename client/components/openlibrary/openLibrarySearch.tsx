"use client"

import { Button } from '@mantine/core';
import { useForm } from '@mantine/form';


export default function OpenLibrarySearch ({submitSearch, loading}: {submitSearch: Function, loading: boolean}) {
	const form = useForm({
		mode: 'uncontrolled',
		initialValues: {
			title: '',
			author: '',
			publisher: '',
		},
	});

	const formSubmit = (values: any) => {
		const {title, author, publisher} = values;
		submitSearch({title, author, publisher})
	}

	return (
		<div>
			<form onSubmit={formSubmit}>
			<Button loading={loading} disabled={loading} type="submit">Submit</Button>
			</form>
		</div>
	)

};
