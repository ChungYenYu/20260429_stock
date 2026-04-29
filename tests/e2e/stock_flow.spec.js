const { test, expect } = require('@playwright/test');

test('should display the landing page with stocks', async ({ page }) => {
  await page.goto('http://localhost:3000');
  
  // Verify "Invest in what matters" exists
  await expect(page.getByText('Invest in what matters')).toBeVisible();
  
  // Verify "TSMC" card exists
  // Since it might take a moment for the data to load from the backend
  await expect(page.getByText('TSMC')).toBeVisible({ timeout: 10000 });
});
