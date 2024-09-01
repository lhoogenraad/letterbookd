import { notifications } from "@mantine/notifications";
import { NotificationPosition } from "@mantine/notifications/lib/notifications.store";

const info = ({message, title=""}) => {
	notifications.show({
		message,
		position: 'top-center',
		title,
	})
}

export default {
	info,
};
