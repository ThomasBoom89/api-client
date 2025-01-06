import { getContext, setContext } from 'svelte';
import { NavigationState } from '$lib/enums/NavigationState.ts';
import { goto } from '$app/navigation';

export class NavigationSystem {
	private _currentState: NavigationState = $state(NavigationState.Home);
	private _currentCollectionId: number = $state(0);
	private _currentProjectId: number = $state(0);

	get currentState(): NavigationState {
		return this._currentState;
	}

	get currentCollectionId(): number {
		return this._currentCollectionId;
	}

	get currentProjectId(): number {
		return this._currentProjectId;
	}

	public navigateToHome(): void {
		goto('/').then(() => {
			this._currentState = NavigationState.Home;
		});
	}

	public navigateToOverview(): void {
		goto('/overview').then(() => {
			this._currentState = NavigationState.Overview;
		});
	}

	public navigateToProject(projectId: number): void {
		goto(`/project/${projectId}`).then(() => {
			this._currentProjectId = projectId;
			this._currentState = NavigationState.Project;
		});
	}

	public navigateToCollection(collectionId: number): void {
		goto(`/collection/${collectionId}`).then(() => {
			this._currentCollectionId = collectionId;
			this._currentState = NavigationState.Collection;
		});
	}
}

const navigationSystemContextKey = 'navigationSystem';

export function initializeNavigationSystem(): void {
	const navigationSystem = new NavigationSystem();
	setContext<NavigationSystem>(navigationSystemContextKey, navigationSystem);
}

export function getNavigationSystem(): NavigationSystem {
	return getContext<NavigationSystem>(navigationSystemContextKey);
}
