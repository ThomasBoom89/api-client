import { expect, test, vi } from 'vitest';
import { configuration } from './wailsjs/go/models';
import { ConfigurationStore } from './configurationStore.svelte';
import * as models from '$lib/wailsjs/go/frontend/Configuration';
import Configuration = configuration.Configuration;
import { SettingStore } from '$lib/settingStore.svelte.ts';

test('toggle skip tls verify', async () => {
	// Set needs to be mocked
	vi.spyOn(models, 'Set').mockImplementation((): Promise<void> => {
		return new Promise<void>((resolve) => {
			resolve();
		});
	});

	const configuration = new Configuration('{"skip_tls_verify":false}');
	expect(configuration.skip_tls_verify).toBe(false);
	const configurationStore = new ConfigurationStore(configuration);
	const settingStore = new SettingStore(configurationStore);
	settingStore.initialize();
	expect(settingStore.currentTlsVerify).toBe(false);
	settingStore.currentTlsVerify = true;
	expect(settingStore.currentTlsVerify).toBe(true);
});
