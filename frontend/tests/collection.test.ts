import { expect, test } from '@playwright/test';
import { cleanupProjects, setupProjects } from './setup';

const projectSetupUUID = crypto.randomUUID();

test.beforeEach('project setup', async ({ page }) => {
	await setupProjects(page, projectSetupUUID);
});

test.afterEach('project cleanup', async ({ page }) => {
	await cleanupProjects(page, projectSetupUUID);
});

test('collection workflow', async ({ page }) => {
	const collectionUUID = crypto.randomUUID();
	await page.getByRole('textbox', { name: 'insert new collection name' }).fill(collectionUUID);
	await page.getByRole('button', { name: 'create' }).click();

	await expect(page.getByRole('textbox', { name: 'insert new collection name' })).toBeEmpty();
	await expect(page.getByTestId('collections').getByRole('button', { name: collectionUUID })).toBeVisible();

	// update
	await page.getByRole('listitem').filter({ hasText: collectionUUID }).getByRole('button', { name: 'edit' }).click();

	const otherCollectionUUID = crypto.randomUUID();

	await page.getByTestId('collections').getByRole('textbox').fill(otherCollectionUUID);
	await page.getByTestId('collections').getByRole('button', { name: 'save' }).click();

	// delete
	await page
		.getByRole('listitem')
		.filter({ hasText: otherCollectionUUID })
		.getByRole('button', { name: 'delete' })
		.click();

	await expect(page.getByRole('listitem').getByRole('button', { name: otherCollectionUUID })).not.toBeVisible();
});
