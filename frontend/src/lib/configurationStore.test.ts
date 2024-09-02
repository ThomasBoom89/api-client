import { expect, test, vi } from 'vitest';
import { configuration } from './wailsjs/go/models';
import { ConfigurationStore } from './configurationStore.svelte';
import * as models from '$lib/wailsjs/go/frontend/Configuration';
import Configuration = configuration.Configuration;

test('toggle theme', () => {
	// Set needs to be mocked
	vi.spyOn(models, 'Set').mockImplementationOnce((): Promise<void> => {
		return new Promise<void>((resolve) => {
			resolve();
		});
	});
	const configuration = new Configuration('{"theme":"dark"}');
	expect(configuration.theme).toBe('dark');
	const configurationStore = new ConfigurationStore(configuration);
	expect(configurationStore.configuration.theme).toBe('dark');
	configurationStore.setProperty('theme', 'light');
	expect(configurationStore.configuration.theme).toBe('light');
});
