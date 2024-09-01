import { createTheme, MantineProvider } from '@mantine/core';

const theme = createTheme({
  /** Put your mantine theme override here */
});

export default function Page() {
  return (
    <MantineProvider theme={theme}>
		<div>Application</div>
    </MantineProvider>
  );
};
