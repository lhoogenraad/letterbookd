import { MantineProvider, createTheme } from '@mantine/core';

export default function RootLayout({
	children,
}: {
	children: React.ReactNode
}) {

	const theme = createTheme({
		fontFamily: 'Verdana, sans-serif',
		fontFamilyMonospace: 'Monaco, Courier, monospace',
		headings: { fontFamily: 'Greycliff CF, sans-serif' },
	});

	return (
		<html lang="en">
			<body>
				<MantineProvider theme={theme}>
					{children}
				</MantineProvider>
			</body>
		</html>
	)
}
