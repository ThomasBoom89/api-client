import { expect, test } from '@playwright/test';
import { cleanupCollections, setupCollections } from './setup';

const projectSetupUUID = crypto.randomUUID();
const collectionsSetupUUID = crypto.randomUUID();

test.beforeEach('collection and project setup', async ({ page }) => {
	await setupCollections(page, projectSetupUUID, collectionsSetupUUID);
});

test.afterEach('collection and project cleanup', async ({ page }) => {
	await cleanupCollections(page, projectSetupUUID, collectionsSetupUUID);
});

test('request workflow', async ({ page }) => {
	const requestUUID = crypto.randomUUID();

	await page.locator('#new-request').fill(requestUUID);
	await page.locator('#create-new-request').click();

	await expect(page.locator('#new-request')).toBeEmpty();
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
