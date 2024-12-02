import { expect, test } from '@playwright/test';

test('application is navigable', async ({ page }) => {
	await page.goto('/');
	const loader = page.locator('#loader');
	await loader.waitFor({ state: 'hidden', timeout: 10000 });
	await expect(page).toHaveTitle('Api-Client :: Startseite');
	await page.waitForSelector('footer');
	await page.getByRole('link', { name: 'Project Overview' }).click();

	await expect(page).toHaveTitle('Api-Client :: Project Overview');
});
