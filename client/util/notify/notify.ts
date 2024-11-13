import { notifications } from "@mantine/notifications";

const info = ({message, title="Notification"}) => {
	notifications.show({
		message,
		position: 'top-center',
		title,
		color: 'blue',
		style: {
			backgroundColor: 'primary.0',
			color: 'text.0',
			width: '100%',
		}
	})
}

const error = ({message="Unexpected error occured", title="Error"}) => {
	notifications.show({
		message,
		position: 'top-center',
		title,
		color: 'red',
		style: {
			backgroundColor: 'primary.0',
			color: 'text.0',
			width: '100%',
		}
	})
}

export default {
	info,
	error,
};
