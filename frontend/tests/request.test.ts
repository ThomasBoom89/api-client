import { expect, test } from '@playwright/test';
import { cleanupCollections, setupCollections } from './setup';
import { RequestTypes } from '$lib/enums/RequestTypes.ts';

const projectSetupUUID = crypto.randomUUID();
const collectionsSetupUUID = crypto.randomUUID();

test.beforeEach('collection and project setup', async ({ page }) => {
	await setupCollections(page, projectSetupUUID, collectionsSetupUUID);
});

test.afterEach('collection and project cleanup', async ({ page }) => {
	await cleanupCollections(page, projectSetupUUID, collectionsSetupUUID);
});

test('http request workflow', async ({ page }) => {
	const requestUUID = crypto.randomUUID();

	await page.locator('#new-request').fill(requestUUID);
	await page.locator('#request-type').selectOption(RequestTypes.HTTP);
	await page.locator('#create-new-request').click();

	await expect(page.locator('#new-request')).toBeEmpty();
	await expect(page.getByTestId('requests').getByRole('button', { name: requestUUID })).toBeVisible();

	//update
	await page.getByTestId('request').filter({ hasText: requestUUID }).getByLabel('edit').click();

	const otherRequestUUID = crypto.randomUUID();
	await page.getByTestId('requests').getByRole('textbox').fill(otherRequestUUID);
	await page.getByTestId('requests').getByLabel('save').click();

	// delete
	await page.getByTestId('request').filter({ hasText: otherRequestUUID }).getByLabel('delete').click();

	await expect(page.getByTestId('request').getByRole('button', { name: otherRequestUUID })).not.toBeVisible();
});

test('websocket request workflow', async ({ page }) => {
	const requestUUID = crypto.randomUUID();

	await page.locator('#new-request').fill(requestUUID);
	await page.locator('#request-type').selectOption(RequestTypes.WEBSOCKET);
	await page.locator('#create-new-request').click();

	await expect(page.locator('#new-request')).toBeEmpty();
	await expect(page.getByTestId('requests').getByRole('button', { name: requestUUID })).toBeVisible();

	//update
	await page.getByTestId('request').filter({ hasText: requestUUID }).getByLabel('edit').click();

	const otherRequestUUID = crypto.randomUUID();
	await page.getByTestId('requests').getByRole('textbox').fill(otherRequestUUID);
	await page.getByTestId('requests').getByLabel('save').click();

	// delete
	await page.getByTestId('request').filter({ hasText: otherRequestUUID }).getByLabel('delete').click();

	await expect(page.getByTestId('request').getByRole('button', { name: otherRequestUUID })).not.toBeVisible();
});
