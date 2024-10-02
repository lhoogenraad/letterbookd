import { notifications } from "@mantine/notifications";

const info = ({message, title="Notification"}) => {
	notifications.show({
		message,
		position: 'top-center',
		title,
		color: 'blue',
		style: {
			backgroundColor: 'whitesmoke',
			color: 'black',
			width: '100%',
		}
	})
}

const error = ({message, title="Error"}) => {
	notifications.show({
		message,
		position: 'top-center',
		title,
		color: 'red',
		style: {
			backgroundColor: 'whitesmoke',
			color: 'black',
			width: '100%',
		}
	})
}

export default {
	info,
	error,
};
