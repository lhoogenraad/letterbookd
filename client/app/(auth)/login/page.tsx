"use client"

import { Button, Group, TextInput } from '@mantine/core';
import { useForm } from '@mantine/form';
import { useRouter } from 'next/navigation';

import '@mantine/core/styles/global.css';
import '@mantine/core/styles/UnstyledButton.css';
import '@mantine/core/styles/Button.css';
import './style.css'
import notify from 'util/notify/notify';
import api from 'util/api/api';
import React from 'react';
import Link from 'next/link';
import auth from 'util/auth/token';


function LoginPage() {
	// State
	const [loading, setLoading] = React.useState(false);


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
		console.log("OK we are logging in")
		setLoading(true);
		await api.users.signIn({email, password})
		.then((res) => {
			localStorage.setItem("authToken", res.data);
			auth.saveToken(res.data);
			console.log("saving token to cookies")
			router.push("/dashboard");
		})
		.catch((err) => {
			notify.error({
				message: err?.response?.data?.Message
				??
				"Something went wrong while logging you in. Please try again."
			});
		})
		.finally(() => setLoading(false));
	};

	return (
		<div className="container">
			<form className="form" onSubmit={form.onSubmit(async (values) => await login(values.email, values.password )) } >
				<div className='header'>Login to Shelfd</div>

				<TextInput
					className='form-input'
					label="Email"
					placeholder="your@email.com"
					key={form.key('email')}
					{...form.getInputProps('email')}
				/>

				<TextInput
					className='form-input'
					type="password"
					label="Password"
					placeholder="Your password"
					key={form.key('password')}
					{...form.getInputProps('password')}
				/>

				<Group justify="flex-end" mt="md">
					<Button color="secondary.8" loaderProps={{ type: 'bars' }} loading={loading} type="submit">Login</Button>
				</Group>

				<div className='switch-page-text'>
					Don't have an account? <Link href="/signup">Signup here</Link>.
				</div>
			</form>
		</div>
	);
}

export default LoginPage;
