/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			gridTemplateColumns: {
				projectOverview: 'repeat(auto-fit, minmax(20rem, 1fr))',
				collectionOverview: 'repeat(auto-fit, minmax(20rem, 1fr))',
			},
			colors: {
				background: 'var(--color-background)',
				text: 'var(--color-text)',
				http: 'var(--color-http)',
				'background-accent': 'var(--color-background-accent)',
				'text-accent': 'var(--color-text-accent)',
				'text-disabled': 'var(--color-text-disabled)',
				'text-highlight': 'var(--color-text-highlight)',
			},
		},
	},
	plugins: [],
};
