"use client"

import { Button, Group, MantineProvider, TextInput } from '@mantine/core';
import { useForm } from '@mantine/form';
import { useRouter } from 'next/navigation';

import '@mantine/core/styles/global.css';
import '@mantine/core/styles/UnstyledButton.css';
import '@mantine/core/styles/Button.css';
import './style.css'
import notify from '../../util/notify';


function LoginPage() {
	const router = useRouter();
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
		const response = await fetch(
			'http://localhost:8080/api/users/login', {
			method: 'POST',
			body: JSON.stringify({
				email, password
			})
		});

		const responseData = await response.json();

		if (response.ok) {
			router.push('/')
			const authToken = responseData;
			localStorage.setItem("authToken", authToken);
		} else {
			const errorMessage : string = responseData.Message;
			notify.info({
				message: errorMessage,
			})
		}
		
	};

	return (
		<div className="login-container">
			<form className="login-form" onSubmit={form.onSubmit(async (values) => await login(values.email, values.password )) } >
				<TextInput
					label="Email"
					placeholder="your@email.com"
					key={form.key('email')}
					{...form.getInputProps('email')}
				/>

				<TextInput
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
		</div>
	);
}

export default LoginPage;
