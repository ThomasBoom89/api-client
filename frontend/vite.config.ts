import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vitest/config';

export default defineConfig({
	plugins: [sveltekit()],
	test: {
		include: ['src/lib/*.test.ts'],
		coverage: {
			include: ['src/lib/*.ts'],
		},
		// we need this for mocking window
		environment: 'happy-dom',
	},
});
