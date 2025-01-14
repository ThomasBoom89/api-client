import { expect, test, vi } from 'vitest';
import * as models from '$app/navigation';
import { NavigationSystem } from '$lib/navigationSystem.svelte.ts';
import { NavigationState } from '$lib/enums/NavigationState.ts';

test('navigate with the navigation system', async () => {
	// Set needs to be mocked
	const godo = vi.spyOn(models, 'goto').mockImplementation((): Promise<void> => {
		return new Promise<void>((resolve) => {
			resolve();
		});
	});

	const navigationSystem: NavigationSystem = new NavigationSystem();
	expect(navigationSystem.currentState).toBe(NavigationState.Home);

	await vi.waitFor(() => {
		navigationSystem.navigateToOverview();
	});
	expect(navigationSystem.currentState).toBe(NavigationState.Overview);

	await vi.waitFor(() => {
		navigationSystem.navigateToProject(214);
	});
	expect(navigationSystem.currentState).toBe(NavigationState.Project);
	expect(navigationSystem.currentProjectId).toBe(214);

	await vi.waitFor(() => {
		navigationSystem.navigateToCollection(897);
	});
	expect(navigationSystem.currentState).toBe(NavigationState.Collection);
	expect(navigationSystem.currentCollectionId).toBe(897);

	await vi.waitFor(() => {
		navigationSystem.navigateToHome();
	});
	expect(navigationSystem.currentState).toBe(NavigationState.Home);
});
