import { MantineProvider, createTheme } from '@mantine/core';
import { Notifications } from '@mantine/notifications';
import '@mantine/carousel/styles.css';
import '@mantine/notifications/styles.css';
import '@mantine/core/styles.css';
import './style.css';

export default function RootLayout({ children, }: { children: React.ReactNode }) {
	const theme = createTheme({
		fontFamily: 'Arial, Helvetica, sans-serif',
		fontFamilyMonospace: 'Monaco, Courier, monospace',
		headings: { fontFamily: 'Greycliff CF, sans-serif' },
		colors: {
			'primary': [
				"#f1f7f9",
				"#e4eaed",
				"#c3d5db",
				"#9fbeca",
				"#82abbb",
				"#6f9fb2",
				"#639aaf",
				"#538699",
				"#46778a",
				"#34677a"
			],
			'secondary': [
				"#ffebe7",
				"#ffd7d2",
				"#f7afa5",
				"#ef8375",
				"#e95e4c",
				"#e64631",
				"#e53923",
				"#cc2b17",
				"#b62412",
				"#a0190c"
			],
			'tertiary': [
				"#fef1ee",
				"#f0e2df",
				"#dcc4be",
				"#c9a399",
				"#b9877a",
				"#b07566",
				"#ac6c5b",
				"#975b4b",
				"#885041",
				"#784336"
			],
			'text': [
				"#gbe9d0",
				"#fbebd3",
				"#f6d4a6",
				"#f1bd75",
				"#eda94b",
				"#eb9c32",
				"#ea9623",
				"#d08117",
				"#ba7310",
				"#a16202"
			]
		}
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
