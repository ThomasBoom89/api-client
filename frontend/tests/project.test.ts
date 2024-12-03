import { expect, test } from '@playwright/test';

test('project workflow', async ({ page }) => {
	await page.goto('/');
	const loader = page.locator('#loader');
	await loader.waitFor({ state: 'hidden', timeout: 10000 });
	await expect(page).toHaveTitle('Api-Client :: Startseite');
	await page.waitForSelector('footer');
	await page.getByRole('link', { name: 'Project Overview' }).click();

	await expect(page).toHaveTitle('Api-Client :: Project Overview');

	const projectUUID = crypto.randomUUID();
	await page.getByRole('textbox', { name: 'insert new project name' }).fill(projectUUID);
	await page.getByRole('button', { name: 'create' }).click();
	await expect(page.getByRole('textbox', { name: 'insert new project name' })).toBeEmpty();
	await expect(page.getByTestId('projects').getByRole('link', { name: projectUUID })).toBeVisible();
	// update
	await page.getByRole('listitem').filter({ hasText: projectUUID }).getByRole('button', { name: 'edit' }).click();

	const newProjectUUID = crypto.randomUUID();
	await page.getByTestId('projects').getByRole('textbox').fill(newProjectUUID);
	await page.getByTestId('projects').getByRole('button', { name: 'save' }).click();

	// delete
	await page.getByRole('listitem').filter({ hasText: newProjectUUID }).getByRole('button', { name: 'delete' }).click();

	await expect(page.getByRole('listitem').getByRole('link', { name: newProjectUUID })).not.toBeVisible();
});
