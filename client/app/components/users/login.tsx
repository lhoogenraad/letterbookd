"use client"

import { Button, Group, TextInput } from '@mantine/core';
import { useForm } from '@mantine/form';

function LoginPage() {
	const form = useForm({
		mode: 'uncontrolled',
		initialValues: {
			email: '',
			password: '',
		},

		validate: {
			email: (value) => (/^\S+@\S+$/.test(value) ? null : 'Invalid email'),
			password: (value) => (!!value ? null : 'Password required'),
		},
	});

	return (
		<form onSubmit={form.onSubmit((values) => console.log(values))}>
			<TextInput
				withAsterisk
				label="Email"
				placeholder="your@email.com"
				key={form.key('email')}
				{...form.getInputProps('email')}
			/>

			<TextInput
				withAsterisk
				type="password"
				label="Password"
				placeholder="Your password"
				key={form.key('password')}
				{...form.getInputProps('password')}
			/>

			<Group justify="flex-end" mt="md">
				<Button type="submit">Submit</Button>
			</Group>
		</form>
	);
}

export default LoginPage;
