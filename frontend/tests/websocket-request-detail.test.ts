import { expect, test } from '@playwright/test';
import { cleanupRequest, setupRequest } from './setup';
import { RequestTypes } from '$lib/enums/RequestTypes.ts';

const projectSetupUUID = crypto.randomUUID();
const collectionsSetupUUID = crypto.randomUUID();
const requestSetupUUID = crypto.randomUUID();

test.beforeEach('project setup', async ({ page }) => {
	await setupRequest(page, projectSetupUUID, collectionsSetupUUID, requestSetupUUID, RequestTypes.WEBSOCKET);
});

test.afterEach('project cleanup', async ({ page }) => {
	await cleanupRequest(page, projectSetupUUID, collectionsSetupUUID, requestSetupUUID);
});

test('request workflow', async ({ page }) => {
	const url = 'ws://localhost:12345/echo';

	await page.locator('#request-url').fill(url);
	await page.locator('#connect').click();
	await expect(page.locator('#disconnect')).toBeVisible();
	await expect(page.locator('#connect')).not.toBeVisible();

	// send message
	const message = 'podman schmodman';
	await page.locator('#message').fill(message);
	await page.locator('#send-message').click();
	await expect(page.locator('#messages > div')).toHaveCount(2);
});
