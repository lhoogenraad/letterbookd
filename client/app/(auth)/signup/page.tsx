"use client"

import { Button, Group, TextInput } from '@mantine/core';
import { useForm } from '@mantine/form';
import { useRouter } from 'next/navigation';

import '@mantine/core/styles/global.css';
import '@mantine/core/styles/UnstyledButton.css';
import '@mantine/core/styles/Button.css';
import '../login/style.css'
import notify from 'util/notify/notify';
import api from 'util/api/api';
import React from 'react';
import Link from 'next/link';

import loginUtil from 'util/auth/login';


function Signup() {
	// State
	const [loading, setLoading] = React.useState(false);


	const router = useRouter();

	const form = useForm({
		mode: 'uncontrolled',
		initialValues: {
			email: '',
			password: '',
			firstName: '',
			lastName: '',
		},

		validate: {
			email: (value) => (/^\S+@\S+$/.test(value) ? null : 'Invalid email'),
			password: (value) => (!!value ? null : 'Password required'),
			firstName: (value) => (!!value ? null : 'First name required'),
			lastName: (value) => (!!value ? null : 'Last name required'),
		},
	});

	const signUp = async (email: string, password: string, firstName: string, lastName: string) => {
		setLoading(true);
		await api.users.signUp({ email, password, firstName, lastName })
			.then((res) => {
				loginUtil.saveToken(res.data);
				router.push("/dashboard");
			})
			.catch((err) => {
				notify.error({
					message: err?.response?.data?.Message
						??
						"Something went wrong while signing up. Please try again."
				});
			})
			.finally(() => setLoading(false));
	};

	return (
		<div className="container">
			<form className="form"
				onSubmit={form.onSubmit(async (values) =>
					await signUp(
						values.email,
						values.password,
						values.firstName,
						values.lastName
					))} >
				<div className='header'>Sign up to Shelfd</div>
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

				<Group className='name-input'>
					<TextInput
						className='form-input'
						label="First Name"
						key={form.key('firstName')}
						{...form.getInputProps('firstName')}
					/>

					<TextInput
						className='form-input'
						label="Last Name"
						key={form.key('lastName')}
						{...form.getInputProps('lastName')}
					/>
				</Group>
				{loading}
				<Group justify="flex-end" mt="md">
					<Button color="secondary.8" loading={loading} loaderProps={{ type: 'bars' }} type="submit">Sign up</Button>
				</Group>

				<div className='switch-page-text'>
					Already have an account? <Link href="/login">Login instead</Link>.
				</div>
			</form>
		</div>
	);
}

export default Signup;
