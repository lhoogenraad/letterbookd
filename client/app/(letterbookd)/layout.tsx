import '@mantine/carousel/styles.css';
import Navbar from '../../components/nav/navbar';
import '@mantine/notifications/styles.css';
import './app.style.css';

export default function RootLayout({ children, }: { children: React.ReactNode }) {
	return (
		<div className='app'>
			<Navbar />
			{children}
		</div>
	)
}
