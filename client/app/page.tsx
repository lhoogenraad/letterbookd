import { MantineProvider, createTheme } from '@mantine/core';

const theme = createTheme({
  fontFamily: 'Verdana, sans-serif',
  fontFamilyMonospace: 'Monaco, Courier, monospace',
  headings: { fontFamily: 'Greycliff CF, sans-serif' },
});

export default function Page() {
  return (
    <MantineProvider theme={theme}>
		<div>Application</div>
    </MantineProvider>
  );
};
