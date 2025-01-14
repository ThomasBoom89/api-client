import { expect, test } from '@playwright/test';
import { setupRequest, cleanupRequest } from './setup';

const projectSetupUUID = crypto.randomUUID();
const collectionsSetupUUID = crypto.randomUUID();
const requestSetupUUID = crypto.randomUUID();

test.beforeEach('project setup', async ({ page }) => {
	await setupRequest(page, projectSetupUUID, collectionsSetupUUID, requestSetupUUID);
});

test.afterEach('project cleanup', async ({ page }) => {
	await cleanupRequest(page, projectSetupUUID, collectionsSetupUUID, requestSetupUUID);
});

test('request workflow', async ({ page }) => {
	const url = 'https://localhost:1234';
	await page.locator('#request-method').selectOption('POST');
	await expect(page.locator('#request-method')).toHaveValue('POST');
	await page.locator('#request-url').fill(url);
	await page.locator('#run-request-icon').click();
	await expect(page.getByTestId('response-error')).toHaveText(
		'Post "' + url + '": dial tcp [::1]:1234: connect: connection refused\n',
	);

	// body
	await page.getByTestId('request-tabs').getByText('Body').click();
	await page.locator('#type-selector').selectOption('json');
	await expect(page.getByTestId('json-body-error')).toBeVisible();
	await page.getByPlaceholder('write your body').fill('{"foo": "bar"}');
	await expect(page.getByTestId('json-body-error')).not.toBeVisible();

	// parameter
	await page.getByTestId('request-tabs').getByText('Parameter').click();
	await expect(page.getByTestId('request-parameter-preview')).toHaveText(url);
	await page.getByPlaceholder('Enter parameter name').fill('hello');
	await page.getByPlaceholder('Enter parameter value').fill('world');
	await page.getByRole('button', { name: 'save' }).click();
	await expect(page.getByTestId('request-parameter-preview')).toHaveText(url + '?hello=world');
	await page.getByTestId('request-parameters').getByRole('button', { name: 'delete' }).click();
	await expect(page.getByTestId('request-parameters')).toBeEmpty();

	// header
	await page.getByTestId('request-tabs').getByText('Header').click();
	await page.getByPlaceholder('Enter header key').fill('foo');
	await page.getByPlaceholder('Enter header value').fill('bar');
	await page.getByRole('button', { name: 'save' }).click();
	await expect(page.getByTestId('request-headers')).toHaveCount(1);
	await page.getByTestId('request-headers').getByRole('button', { name: 'delete' }).click();
	await expect(page.getByTestId('request-headers')).toBeEmpty();
});
