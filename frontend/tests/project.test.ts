import { expect, test } from '@playwright/test';

test('project workflow', async ({ page }) => {
	await page.goto('/');
	const loader = page.locator('#loader');
	await loader.waitFor({ state: 'hidden', timeout: 10000 });
	await expect(page).toHaveTitle('Api-Client :: Startseite');
	await page.waitForSelector('footer');
	await page.getByRole('link', { name: 'Project Overview' }).click();

	await expect(page).toHaveTitle('Api-Client :: Project Overview');

	await page.getByRole('textbox', { name: 'insert new project name' }).fill('new project name');
	await page.getByRole('button', { name: 'create' }).click();
	await expect(page.getByRole('textbox', { name: 'insert new project name' })).toBeEmpty();
	expect(page.getByTestId('projects')).not.toBeNull();
	// update
	await page
		.getByTestId('projects')
		.filter({ hasText: 'new project name' })
		.getByRole('button', { name: 'edit' })
		.first()
		.click();

	await page.getByTestId('projects').getByRole('textbox').fill('other project name');
	await page.getByTestId('projects').getByRole('button', { name: 'save' }).click();

	// delete
	await page
		.getByTestId('projects')
		.filter({ hasText: 'other project name' })
		.getByRole('button', { name: 'delete' })
		.first()
		.click();

	await expect(page.getByTestId('projects')).toBeEmpty();
});
