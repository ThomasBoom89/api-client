import { expect, test } from '@playwright/test';
import { cleanupCollections, setupCollections } from './setup';

const projectSetupUUID = crypto.randomUUID();
const collectionsSetupUUID = crypto.randomUUID();

test.beforeEach('project setup', async ({ page }) => {
	await setupCollections(page, projectSetupUUID, collectionsSetupUUID);
});

test.afterEach('project cleanup', async ({ page }) => {
	await cleanupCollections(page, projectSetupUUID, collectionsSetupUUID);
});

test('request workflow', async ({ page }) => {
	const requestUUID = crypto.randomUUID();
	await page.getByRole('textbox', { name: 'insert new request name' }).fill(requestUUID);
	await page.getByRole('button', { name: 'create' }).click();
	await expect(page.getByRole('textbox', { name: 'insert new request name' })).toBeEmpty();
	await expect(page.getByTestId('requests').getByRole('listitem').filter({ hasText: requestUUID })).toBeVisible();

	//update
	await page.getByRole('listitem').filter({ hasText: requestUUID }).getByRole('button', { name: 'edit' }).click();

	const otherRequestUUID = crypto.randomUUID();
	await page.getByTestId('requests').getByRole('textbox').fill(otherRequestUUID);
	await page.getByTestId('requests').getByRole('button', { name: 'save' }).click();

	// delete
	await page
		.getByRole('listitem')
		.filter({ hasText: otherRequestUUID })
		.getByRole('button', { name: 'delete' })
		.click();

	await expect(page.getByRole('listitem').filter({ hasText: otherRequestUUID })).not.toBeVisible();
});
