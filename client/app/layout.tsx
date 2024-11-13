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
		autoContrast: true,
		headings: { fontFamily: 'Greycliff CF, sans-serif' },
		colors: {
			'primary': [
				"#121212",
				"#121212",
				"#121212",
				"#121212",
				"#121212",
				"#121212",
				"#121212",
				"#121212",
				"#121212",
				"#121212",
			],
			'elevate': [
				'rgba(90, 90, 90, 0.165)',
				'rgba(80, 80, 80, 0.165)',
				'rgba(70, 70, 70, 0.165)',
				'rgba(60, 60, 60, 0.165)',
				'rgba(50, 50, 50, 0.165)',
				'rgba(40, 40, 40, 0.165)',
				'rgba(30, 30, 30, 0.165)',
				'rgba(20, 20, 20, 0.165)',
				'rgba(10, 10, 10, 0.165)',
				'rgba(5, 5, 5, 0.165)',
			],
			'secondary': [
				"#f6f6f3",
				"#eaeae8",
				"#d3d4cd",
				"#babdae",
				"#a5a993",
				"#989d82",
				"#919778",
				"#7d8366",
				"#6f7459",
				"#5e6447"
			],
			'tertiary': [
				"#874f41",
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
				'#fbe9d0',
				'#fbebd3',
				'#f6d4a6',
				'#f1bd75',
				'#eda94b',
				'#eb9c32',
				'#ea9623',
				'#d08117',
				'#ba7310',
				'#a16202'
			]
		}
	});

	return (
		<html lang="en">
		<body className='body'>
		<MantineProvider theme={theme} defaultColorScheme="dark">
		<Notifications position="top-center" limit={5} />
		{children}
		</MantineProvider>
		</body>
		</html>
	)
}
