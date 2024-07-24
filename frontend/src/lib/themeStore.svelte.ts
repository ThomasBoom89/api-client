import { getContext, setContext } from 'svelte';
import { ConfigurationStore } from './configurationStore.svelte.js';

export class ThemeStore {
	private _currentTheme: string = $state<string>('');

	get currentTheme(): string {
		return this._currentTheme;
	}

	set currentTheme(currentTheme: string) {
		this._currentTheme = currentTheme;
		this.configurationStore.setProperty('theme', currentTheme);
		document.documentElement.setAttribute('data-theme', currentTheme);
	}

	public constructor(private readonly configurationStore: ConfigurationStore) {
	}

	public switchTheme(): string {
		if (this.currentTheme === 'dark') {
			this.currentTheme = 'light';
		} else {
			this.currentTheme = 'dark';
		}

		return this.currentTheme;
	}

	public initialize(): void {
		if (this.configurationStore.configuration.theme === '') {
			this.currentTheme = this.getPreferredTheme();
		} else {
			this.currentTheme = this.configurationStore.configuration.theme;
		}
	}

	private getPreferredTheme(): string {
		if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
			return 'dark';
		}

		return 'light';
	}
}

const themeStoreContextKey = 'themeStore';

export function initializeThemeStore(configurationStore: ConfigurationStore): void {
	const themeStore = new ThemeStore(configurationStore);
	themeStore.initialize();
	setContext<ThemeStore>(themeStoreContextKey, themeStore);
}

export function getThemeStore(): ThemeStore {
	return getContext<ThemeStore>(themeStoreContextKey);
}
