import { expect, test } from '@playwright/test';

test('toggle dark light mode', async ({ page }) => {
	await page.goto('/');
	await expect(page.locator('h1')).toBeVisible();
	const currentTheme = page.locator('html').getAttribute('data-theme');
	await page.getByTestId('dark-light-toggle').click();

	expect(currentTheme).not.toEqual(page.locator('html').getAttribute('data-theme'));
});
