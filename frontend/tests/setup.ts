import { expect, type Page } from '@playwright/test';

export async function navigateToProject(page: Page): Promise<void> {
	await page.goto('/');
	const loader = page.locator('#loader');
	await loader.waitFor({ state: 'hidden', timeout: 10000 });
	await expect(page).toHaveTitle('Api-Client :: Startseite');
	await page.waitForSelector('footer');
	await page.getByRole('link', { name: 'Project Overview' }).click();
	await expect(page).toHaveTitle('Api-Client :: Project Overview');
}

export async function setupProjects(page: Page, projectSetupUUID: string): Promise<void> {
	await navigateToProject(page);

	await page.getByRole('textbox', { name: 'insert new project name' }).fill(projectSetupUUID);
	await page.getByRole('button', { name: 'create' }).click();
	await page.getByTestId('projects').getByRole('link', { name: projectSetupUUID }).click();

	await expect(page).toHaveTitle('Api-Client :: Project Details');
	await expect(page.getByRole('heading', { name: `Project ${projectSetupUUID}` })).toBeVisible();
}

export async function cleanupProjects(page: Page, projectSetupUUID: string) {
	await navigateToProject(page);

	await page
		.getByRole('listitem')
		.filter({ hasText: projectSetupUUID })
		.getByRole('button', { name: 'delete' })
		.click();
	await expect(page.getByTestId('projects').getByRole('link', { name: projectSetupUUID })).not.toBeVisible();
}

export async function setupCollections(page: Page, projectSetupUUID: string, collectionSetupUUID: string) {
	await setupProjects(page, projectSetupUUID);
	await page.getByRole('textbox', { name: 'insert new collection name' }).fill(collectionSetupUUID);
	await page.getByRole('button', { name: 'create' }).click();
	await expect(page.getByRole('textbox', { name: 'insert new collection name' })).toBeEmpty();
	await expect(page.getByTestId('collections').getByRole('link', { name: collectionSetupUUID })).toBeVisible();
	await page.getByTestId('collections').getByRole('link', { name: collectionSetupUUID }).click();
	await expect(page.getByRole('heading', { name: `Collection ${collectionSetupUUID}` })).toBeVisible();
}

export async function cleanupCollections(page: Page, projectSetupUUID: string, collectionsSetupUUID: string) {
	await navigateToProject(page);
	await page.getByTestId('projects').getByRole('link', { name: projectSetupUUID }).click();
	await expect(page.getByRole('heading', { name: `Project ${projectSetupUUID}` })).toBeVisible();
	await page
		.getByRole('listitem')
		.filter({ hasText: collectionsSetupUUID })
		.getByRole('button', { name: 'delete' })
		.click();

	await expect(page.getByRole('listitem').getByRole('link', { name: collectionsSetupUUID })).not.toBeVisible();
	await cleanupProjects(page, projectSetupUUID);
}
