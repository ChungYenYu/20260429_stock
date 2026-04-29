# 20260429_stock 股市查詢網站 - 全自動化實作與測試計畫

> **For Hermes:** Use subagent-driven-development skill to implement this plan task-by-task.

**Goal:** 建立一個使用 Vue 3 (Apple Style)、Golang 和 PostgreSQL 的股市查詢網站。開發過程包含完整的單元測試與集成測試，並在交付前通過全自動化測試驗證。

**Architecture:** 
- **Style:** Apple Premium Minimalism (DESIGN.md - Apple Spec)。
- **Frontend:** Vue 3 + Tailwind CSS (Apple Style)。
- **Backend:** Golang (Gin) + GORM + PostgreSQL。
- **Testing:** Golang (Standard Test + Mocking), Vue (Vitest + Playwright for E2E)。
- **Infra:** Docker Compose。

---

### Task 1: 更新 DESIGN.md 並配置 Tailwind Apple Style
**Objective:** 在專案中建立 `DESIGN.md` 並根據 Apple 設計規範設定前端 Tailwind。

**Files:**
- Create: `DESIGN.md` (Apple Spec)
- Modify: `frontend/tailwind.config.js`
- Modify: `frontend/src/style.css`

**Step 1: 建立 DESIGN.md**
（包含 Apple Blue #0071e3, SF Pro 字體規範, Cinematic Rhythm 等）

**Step 2: 配置 Tailwind**
設定色彩 `apple-blue`, `apple-gray`, `near-black` 與字體家族。

---

### Task 2: Backend - 實作 Stock 模型與 API 測試 (TDD)
**Objective:** 建立資料庫模型並實作具有自動化測試的 API。

**Files:**
- Create: `backend/models/stock.go`
- Create: `backend/models/stock_test.go`
- Create: `backend/handlers/stock_handler.go`
- Create: `backend/handlers/stock_handler_test.go`

**Step 1: 撰寫 Stock 模型單元測試**
測試 GORM 的 AutoMigrate 與 CRUD 邏輯。

**Step 2: 撰寫 API Handlers 單元測試**
使用 `httptest` 模擬 API 請求並驗證回傳 JSON。

---

### Task 3: Backend - 串接 PostgreSQL 與 Mock 數據
**Objective:** 建立資料庫連線池，並在啟動時自動插入測試數據（如 0056, 00919, 2330）。

**Files:**
- Modify: `backend/main.go`
- Create: `backend/database/db.go`

**Step 1: 實作資料庫初始化邏輯**
**Step 2: 實作 Health Check 與 DB Ping 測試**

---

### Task 4: Frontend - 實作 Apple Style UI 組件與測試
**Objective:** 建立符合 Apple 規範的 Navbar、Hero Section 與 Stock Card。

**Files:**
- Create: `frontend/src/components/AppleNavbar.vue`
- Create: `frontend/src/components/StockCard.vue`
- Create: `frontend/src/components/__tests__/StockCard.spec.js`

**Step 1: 撰寫 Vitest 組件測試**
驗證 Stock Card 傳入不同數據（漲跌色）時的呈現效果。

**Step 2: 實作 Apple Navbar (Glass Effect)**
使用 `backdrop-filter: blur(20px)`。

---

### Task 5: 整合測試與 E2E 驗證
**Objective:** 使用 Playwright 進行全流程自動化測試。

**Files:**
- Create: `tests/e2e/stock_flow.spec.js`
- Create: `scripts/test-and-deliver.sh`

**Step 1: 撰寫 E2E 腳本**
模擬使用者打開網頁、查看 0056 股價、點擊詳細資料。

**Step 2: 建立交付腳本 `test-and-deliver.sh`**
執行順序：
1. `docker-compose up -d db`
2. `go test ./backend/...`
3. `cd frontend && npm run test:unit`
4. `npm run test:e2e`
5. 若全過，輸出 "Tests Passed - Ready for Delivery"。

---

### Task 6: 最終部署與 GitHub 同步
**Objective:** 確保所有代碼與測試通過後，推送到 GitHub。
