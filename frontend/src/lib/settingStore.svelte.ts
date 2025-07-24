import { getContext, setContext } from 'svelte';
import { ConfigurationStore } from './configurationStore.svelte.js';

export class SettingStore {
	private _currentTlsVerify: boolean = $state<boolean>(false);

	get currentTlsVerify(): boolean {
		return this._currentTlsVerify;
	}

	set currentTlsVerify(state: boolean) {
		this._currentTlsVerify = state;
		this.configurationStore.setProperty('skip_tls_verify', state);
	}

	public constructor(private readonly configurationStore: ConfigurationStore) {}

	public initialize() {
		this._currentTlsVerify = this.configurationStore.configuration.skip_tls_verify;
	}
}

const settingStoreContextKey = 'settingStore';

export function initializeSettingStore(configurationStore: ConfigurationStore): void {
	const settingStore = new SettingStore(configurationStore);
	settingStore.initialize();
	setContext<SettingStore>(settingStoreContextKey, settingStore);
}

export function getSettingStore(): SettingStore {
	return getContext<SettingStore>(settingStoreContextKey);
}
