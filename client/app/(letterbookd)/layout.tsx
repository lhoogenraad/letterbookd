import '@mantine/carousel/styles.css';
import Navbar from '../../components/nav/navbar';
import '@mantine/notifications/styles.css';

export default function RootLayout({children,}: {children: React.ReactNode}) {
	return (
		<html lang="en">
			<body>
				<Navbar />
				{children}
			</body>
		</html>
	)
}
