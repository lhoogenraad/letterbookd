import { MantineProvider, createTheme } from '@mantine/core';
import { Notifications } from '@mantine/notifications';
import '@mantine/carousel/styles.css';
import '@mantine/notifications/styles.css';
import '@mantine/core/styles.css';
import './style.css';

export default function RootLayout({children,}: {children: React.ReactNode}) {
	const theme = createTheme({
		fontFamily: 'Arial, Helvetica, sans-serif',
		fontFamilyMonospace: 'Monaco, Courier, monospace',
		headings: { fontFamily: 'Greycliff CF, sans-serif' },
	});

	return (
		<html lang="en">
			<body className='body'>
				<MantineProvider theme={theme}>
					<Notifications position="top-center" limit={5} />
					{children}
				</MantineProvider>
			</body>
		</html>
	)
}
