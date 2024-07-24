import { configuration } from '$lib/wailsjs/go/models.ts';
import { getContext, setContext } from 'svelte';
import { Set } from '$lib/wailsjs/go/frontend/Configuration';

export class ConfigurationStore {
	private readonly _configuration: configuration.Configuration = $state(new configuration.Configuration());

	get configuration(): configuration.Configuration {
		return this._configuration;
	}

	constructor(configuration: configuration.Configuration) {
		this._configuration = configuration;
	}

	/* eslint-disable @typescript-eslint/no-explicit-any */
	setProperty(property: string, value: any): void {
		// @ts-expect-error this is fine!
		this._configuration[property] = value;
		Set(this._configuration);
	}
}

const configurationStoreContextKey = 'configurationStore';

export function initializeConfigurationStore(configuration: configuration.Configuration): void {
	setContext<ConfigurationStore>(configurationStoreContextKey, new ConfigurationStore(configuration));
}

export function getConfigurationStore(): ConfigurationStore {
	return getContext<ConfigurationStore>(configurationStoreContextKey);
}
