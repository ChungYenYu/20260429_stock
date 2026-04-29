const { chromium } = require('playwright');
const path = require('path');

(async () => {
  const browser = await chromium.launch({
    args: ['--no-sandbox', '--disable-setuid-sandbox']
  });
  const page = await browser.newPage();
  await page.setViewportSize({ width: 1280, height: 1000 });
  
  console.log('Navigating to http://localhost:3000...');
  await page.goto('http://localhost:3000', { waitUntil: 'networkidle' });
  
  // Wait for the app to be visible
  await page.waitForSelector('#app');
  
  // Wait a bit for animations/data to load
  await page.waitForTimeout(2000);
  
  const screenshotPath = path.join(__dirname, '..', 'assets', 'screenshot.png');
  console.log(`Saving screenshot to ${screenshotPath}...`);
  await page.screenshot({ path: screenshotPath, fullPage: true });
  
  await browser.close();
  console.log('Done.');
})();
