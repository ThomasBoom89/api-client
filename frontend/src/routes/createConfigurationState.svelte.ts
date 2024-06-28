import { Write } from '$lib/wailsjs/go/configuration/ReadWriter';
import { configuration } from '$lib/wailsjs/go/models.ts';

class ConfigurationStore {
	set configuration(value: configuration.Configuration) {
		this._configuration = value;
	}

	get configuration(): configuration.Configuration {
		return this._configuration;
	}

	private _configuration: configuration.Configuration = $state(new configuration.Configuration());

	setProperty(property: string, value: any): void {
		// @ts-ignore
		this._configuration[property] = value;
		Write(this._configuration);
	}
}

export const configurationStore = new ConfigurationStore();
