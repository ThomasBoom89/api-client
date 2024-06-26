/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			colors: {
				background: 'var(--color-background)',
				text: 'var(--color-text)',
				primary: 'var(--color-primary)',
				secondary: 'var(--color-secondary)',
				buttons: 'var(--color-buttons)',
				typography: 'var(--color-typography)',
				'background-accent': 'var(--color-background-accent)'
			}
		}
	}, plugins: []
};
