"use client"

import { Button, Group, MantineProvider, TextInput } from '@mantine/core';
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

	const login = async (email: string, password: string) => {
		const data = await fetch(
			'http://localhost:8080/api/users/login', {
			method: 'POST',
			body: JSON.stringify({
				email, password
			})
		});
		const authToken = await data.json();
		localStorage.setItem("authToken", authToken);
	};

	return (
		<MantineProvider>
			<form onSubmit={form.onSubmit(async (values) => await login(values.email, values.password )) } >
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
		</MantineProvider>
	);
}

export default LoginPage;
