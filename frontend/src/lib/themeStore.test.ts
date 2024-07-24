import { expect, test, vi } from 'vitest';
import { configuration } from './wailsjs/go/models';
import { ConfigurationStore } from './configurationStore.svelte';
import { ThemeStore } from './themeStore.svelte';
import * as models from '$lib/wailsjs/go/frontend/Configuration';
import Configuration = configuration.Configuration;

test('switch from dark to light mode', async () => {
	// Set needs to be mocked
	vi.spyOn(models, 'Set').mockImplementation((): Promise<void> => {
		return new Promise<void>((resolve) => {
			resolve();
		});
	});
	vi.spyOn(window, 'matchMedia').mockImplementationOnce(() => {
		return new MediaQueryList();
	});
	const configuration = new Configuration('{"theme":"dark"}');
	expect(configuration.theme).toBe('dark');
	const configurationStore = new ConfigurationStore(configuration);
	const themeStore = new ThemeStore(configurationStore);
	themeStore.initialize();
	expect(themeStore.currentTheme).toBe('dark');
	themeStore.switchTheme();
	expect(themeStore.currentTheme).toBe('light');
});
