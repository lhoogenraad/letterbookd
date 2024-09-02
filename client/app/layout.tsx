import { MantineProvider, createTheme } from '@mantine/core';
import { Notifications } from '@mantine/notifications';

export default function RootLayout({children,}: {children: React.ReactNode}) {
	const theme = createTheme({
		fontFamily: 'Verdana, sans-serif',
		fontFamilyMonospace: 'Monaco, Courier, monospace',
		headings: { fontFamily: 'Greycliff CF, sans-serif' },
	});

	return (
		<html lang="en">
			<body>
				<MantineProvider theme={theme}>
					<Notifications position="top-center" limit={5} />
					{children}
				</MantineProvider>
			</body>
		</html>
	)
}
