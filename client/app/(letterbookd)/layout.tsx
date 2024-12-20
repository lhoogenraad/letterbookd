import '@mantine/carousel/styles.css';
import Navbar from 'components/nav/navbar';
import Footer from 'components/footer/footer';
import '@mantine/notifications/styles.css';
import './app.style.css';

export default function RootLayout({ children, }: { children: React.ReactNode }) {
	return (
		<div className="app-container">
			<div className="upper-content">
				<Navbar />
				<div className='app'>
					{children}
				</div>
			</div>
			<div className="lower-content">
				<Footer />
			</div>
		</div>
	)
}
