import { expect, test } from '@playwright/test';

test('index page has expected h1', async ({ page }) => {
	await page.goto('/');
	await expect(page.getByTestId('title')).toBeVisible();
	await expect(page.getByTestId('title')).toHaveText('todo');
	await expect(page.getByTestId('title')).toHaveCSS('text-transform', 'uppercase');
});

test('index page has background image', async ({ page }) => {
	await page.goto('/');
	await expect(page.getByTestId('background')).toBeVisible();
	await expect(page.getByTestId('background')).toBeInViewport();
});
