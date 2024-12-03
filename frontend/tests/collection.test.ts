import { expect, test } from '@playwright/test';

const projectSetupUUID = crypto.randomUUID();

test.beforeEach('project setup', async ({ page }) => {
	await page.goto('/');
	const loader = page.locator('#loader');
	await loader.waitFor({ state: 'hidden', timeout: 10000 });
	await expect(page).toHaveTitle('Api-Client :: Startseite');
	await page.waitForSelector('footer');
	await page.getByRole('link', { name: 'Project Overview' }).click();

	await expect(page).toHaveTitle('Api-Client :: Project Overview');
	await page.getByRole('textbox', { name: 'insert new project name' }).fill(projectSetupUUID);
	await page.getByRole('button', { name: 'create' }).click();
	await page.getByTestId('projects').getByRole('link', { name: projectSetupUUID }).click();

	await expect(page).toHaveTitle('Api-Client :: Project Details');
	await expect(page.getByRole('heading', { name: `Project ${projectSetupUUID}` })).toBeVisible();
});

test.afterEach('project cleanup', async ({ page }) => {
	await page.goto('/');
	const loader = page.locator('#loader');
	await loader.waitFor({ state: 'hidden', timeout: 10000 });
	await expect(page).toHaveTitle('Api-Client :: Startseite');
	await page.waitForSelector('footer');
	await page.getByRole('link', { name: 'Project Overview' }).click();

	await expect(page).toHaveTitle('Api-Client :: Project Overview');
	await page
		.getByRole('listitem')
		.filter({ hasText: projectSetupUUID })
		.getByRole('button', { name: 'delete' })
		.click();
	await expect(page.getByTestId('projects').getByRole('link', { name: projectSetupUUID })).not.toBeVisible();
});

test('collection workflow', async ({ page }) => {
	const collectionUUID = crypto.randomUUID();
	await page.getByRole('textbox', { name: 'insert new collection name' }).fill(collectionUUID);
	await page.getByRole('button', { name: 'create' }).click();

	await expect(page.getByRole('textbox', { name: 'insert new collection name' })).toBeEmpty();
	await expect(page.getByTestId('collections').getByRole('link', { name: collectionUUID })).toBeVisible();

	// update
	await page.getByRole('listitem').filter({ hasText: collectionUUID }).getByRole('button', { name: 'edit' }).click();

	const otherCollectionUUID = crypto.randomUUID();

	await page.getByTestId('collections').getByRole('textbox').fill(otherCollectionUUID);
	await page.getByTestId('collections').getByRole('button', { name: 'save' }).click();
});
