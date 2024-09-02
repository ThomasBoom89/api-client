import { expect, test } from '@playwright/test';

test('application is navigable', async ({ page }) => {
	await page.goto('/');
	await expect(page).toHaveTitle('Api-Client :: Startseite');
	await page.waitForSelector('footer');
	await page.getByRole('link', { name: 'Project Overview' }).click();

	await expect(page).toHaveTitle('Api-Client :: Project Overview');
	await expect(page.locator('h2')).toBeVisible();
	// await page.getByRole('link', { name: 'first project' }).click();

	// await expect(page).toHaveTitle('Api-Client :: Project Details');
	// await expect(page.locator('h2')).toBeVisible();
	// await page.getByRole('link', { name: 'collection one' }).click();

	// await expect(page).toHaveTitle('Api-Client :: Collection Overview');
	// await expect(page.locator('h2')).toBeVisible();
});
