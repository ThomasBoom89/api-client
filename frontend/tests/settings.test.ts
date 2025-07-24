import { expect, test } from '@playwright/test';

test('switch skip verify flag', async ({ page }) => {
	await page.goto('/');
	const loader = page.locator('#loader');
	await loader.waitFor({ state: 'hidden', timeout: 10000 });
	await page.getByTestId('open-settings').click();
	const skipTlsVerify = page.getByTestId('skip-tls-verify');
	const checked = await skipTlsVerify.isChecked();
	if (checked) {
		await expect(skipTlsVerify).toBeChecked();
		await skipTlsVerify.uncheck();
		await expect(skipTlsVerify).not.toBeChecked();
	} else {
		await expect(skipTlsVerify).not.toBeChecked();
		await skipTlsVerify.check();
		await expect(skipTlsVerify).toBeChecked();
	}
});
