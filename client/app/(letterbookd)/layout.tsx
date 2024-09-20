import { MantineProvider, createTheme } from '@mantine/core';
import { Notifications } from '@mantine/notifications';
import '@mantine/carousel/styles.css';
import Navbar from '../../components/nav/navbar.tsx';
import '@mantine/notifications/styles.css';

export default function RootLayout({children,}: {children: React.ReactNode}) {
	const theme = createTheme({
		fontFamily: 'Arial, Helvetica, sans-serif',
		fontFamilyMonospace: 'Monaco, Courier, monospace',
		headings: { fontFamily: 'Greycliff CF, sans-serif' },
	});

	return (
		<html lang="en">
			<body>
				<Navbar />
				{children}
			</body>
		</html>
	)
}
