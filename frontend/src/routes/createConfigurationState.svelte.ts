import { Write } from '$lib/wailsjs/go/configuration/ReadWriter';
import { configuration } from '$lib/wailsjs/go/models.ts';
import { getContext, setContext } from 'svelte';

class ConfigurationStore {
	private readonly _configuration: configuration.Configuration = $state(new configuration.Configuration());

	get configuration(): configuration.Configuration {
		return this._configuration;
	}

	constructor(configuration: configuration.Configuration) {
		this._configuration = configuration;
	}

	setProperty(property: string, value: any): void {
		// @ts-ignore
		this._configuration[property] = value;
		Write(this._configuration);
	}
}

const configurationStoreContextKey = 'configurationStore';

export function initializeConfigurationStore(configuration: configuration.Configuration): void {
	setContext<ConfigurationStore>(configurationStoreContextKey, new ConfigurationStore(configuration));
}

export function getConfigurationStore(): ConfigurationStore {
	return getContext<ConfigurationStore>(configurationStoreContextKey);
}
