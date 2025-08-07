import { expect, test } from '@playwright/test';
import { cleanupRequest, setupRequest } from './setup';
import { RequestTypes } from '$lib/enums/RequestTypes.ts';

const projectSetupUUID = crypto.randomUUID();
const collectionsSetupUUID = crypto.randomUUID();
const requestSetupUUID = crypto.randomUUID();

test.beforeEach('project setup', async ({ page }) => {
	await setupRequest(page, projectSetupUUID, collectionsSetupUUID, requestSetupUUID, RequestTypes.HTTP);
});

test.afterEach('project cleanup', async ({ page }) => {
	await cleanupRequest(page, projectSetupUUID, collectionsSetupUUID, requestSetupUUID);
});

test('request workflow', async ({ page }) => {
	const url = 'https://localhost:1234';
	await page.locator('#request-method').selectOption('POST');
	await expect(page.locator('#request-method')).toHaveValue('POST');
	await page.locator('#request-url').fill(url);
	await expect(page.locator('#request-url')).toHaveValue(url);
	await expect(page.locator('#run-request')).not.toBeDisabled();
	await page.locator('#run-request').click();
	await expect(page.getByTestId('response-error')).toHaveText(
		'Post "' + url + '": dial tcp [::1]:1234: connect: connection refused\n',
	);

	// body
	await page.getByTestId('request-tabs').getByText('Body').click();
	await page.locator('#body-type').selectOption('json');
	await expect(page.getByTestId('json-body-error')).toBeVisible();
	await page.getByPlaceholder('> your body here <').fill('{"foo": "bar"}');
	await expect(page.getByTestId('json-body-error')).not.toBeVisible();

	// parameter
	await page.getByTestId('request-tabs').getByText('Parameter').click();
	await expect(page.getByTestId('request-parameter-preview')).toHaveText(url);
	await page.locator('#new-parameter-name').fill('hello');
	await page.locator('#new-parameter-value').fill('world');
	await page.getByLabel('save').click();
	await expect(page.getByTestId('request-parameter-preview')).toHaveText(url + '?hello=world');
	await page.getByTestId('request-parameters').getByLabel('delete').click();
	await expect(page.getByTestId('request-parameters')).toBeEmpty();

	// header
	await page.getByTestId('request-tabs').getByText('Header').click();
	await page.locator('#new-header-key').fill('foo');
	await page.locator('#new-header-value').fill('bar');
	await page.getByLabel('save').click();
	await expect(page.getByTestId('request-headers')).toHaveCount(1);
	await page.getByTestId('request-headers').getByLabel('delete').click();
	await expect(page.getByTestId('request-headers')).toBeEmpty();
});
