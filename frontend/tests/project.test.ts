import { expect, test } from '@playwright/test';
import { navigateToProject } from './setup';

test.beforeEach('project setup', async ({ page }) => {
	await navigateToProject(page);
});

test('project workflow', async ({ page }) => {
	const projectUUID = crypto.randomUUID();
	await page.locator('#new-project').fill(projectUUID);
	await page.locator('#create-new-project').click();
	await expect(page.locator('#new-project')).toBeEmpty();
	await expect(page.getByTestId('projects').getByRole('button', { name: projectUUID })).toBeVisible();
	// update
	await page.getByTestId('project').filter({ hasText: projectUUID }).getByLabel('edit').click();

	const newProjectUUID = crypto.randomUUID();
	await page.getByTestId('projects').getByRole('textbox').fill(newProjectUUID);
	await page.getByTestId('projects').getByLabel('save').click();

	// delete
	await page.getByTestId('project').filter({ hasText: newProjectUUID }).getByLabel('delete').click();

	await expect(page.getByTestId('projects').getByRole('button', { name: projectUUID })).not.toBeVisible();
});
