import { notifications } from "@mantine/notifications";

const info = ({message, title=""}) => {
	notifications.show({
		message,
		position: 'top-center',
		title,
		style: {
			backgroundColor: '#222',
			color: '#eee',
			width: "50%",
			margin: "0 auto",
			padding: "1rem",
			marginTop: "1rem",
		}
	})
}

export default {
	info,
};
