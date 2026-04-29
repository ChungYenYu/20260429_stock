# 20260429_stock 股市查詢網站實作計畫

> **For Hermes:** Use subagent-driven-development skill to implement this plan task-by-task.

**Goal:** 建立一個使用 Vue 3、Golang 和 PostgreSQL 的股市查詢網站，並支援 Docker Compose 一鍵部署。

**Architecture:** 前後端分離架構。Golang 後端負責 API 與資料抓取，Vue 3 前端負責 UI 呈現，PostgreSQL 儲存股價歷史與使用者清單。

**Tech Stack:** 
- Frontend: Vue 3 (Vite), Tailwind CSS, Axios
- Backend: Golang (Gin Gonic), GORM
- DB: PostgreSQL 16
- Infra: Docker, Docker Compose

---

### Task 1: 初始化 Docker Compose 環境
**Objective:** 建立基礎的 `docker-compose.yml`，包含 PostgreSQL 服務。

**Files:**
- Create: `docker-compose.yml`
- Create: `.env`

**Step 1: 建立 docker-compose.yml**
```yaml
services:
  db:
    image: postgres:16-alpine
    container_name: stock_db
    environment:
      POSTGRES_USER: ${DB_USER:-postgres}
      POSTGRES_PASSWORD: ${DB_PASSWORD:-password}
      POSTGRES_DB: ${DB_NAME:-stockdb}
    ports:
      - "5432:5432"
    volumes:
      - ./data/postgres:/var/lib/postgresql/data

  backend:
    build: ./backend
    container_name: stock_backend
    ports:
      - "8080:8080"
    environment:
      DB_HOST: db
      DB_PORT: 5432
      DB_USER: ${DB_USER:-postgres}
      DB_PASSWORD: ${DB_PASSWORD:-password}
      DB_NAME: ${DB_NAME:-stockdb}
    depends_on:
      - db

  frontend:
    build: ./frontend
    container_name: stock_frontend
    ports:
      - "3000:3000"
    depends_on:
      - backend
```

**Step 2: 驗證**
Run: `docker-compose config`
Expected: 輸出完整的配置資訊。

---

### Task 2: Backend - 初始化 Golang 專案
**Objective:** 設定 Golang 環境，安裝 Gin 框架並建立一個 Health Check API。

**Files:**
- Create: `backend/go.mod`
- Create: `backend/main.go`
- Create: `backend/Dockerfile`

**Step 1: 建立 main.go**
```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	r.Run(":8080")
}
```

**Step 2: 建立 Dockerfile (Multi-stage)**
```dockerfile
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
```

---

### Task 3: Frontend - 初始化 Vue 3 專案
**Objective:** 使用 Vite 建立 Vue 3 專案並安裝 Tailwind CSS。

**Files:**
- Create: `frontend/Dockerfile`
- Create: `frontend/package.json` (via Vite)

**Step 1: 執行初始化指令**
Run: `cd frontend && npm create vite@latest . -- --template vue`
Run: `npm install -D tailwindcss postcss autoprefixer && npx tailwindcss init -p`

**Step 2: 建立前端 Dockerfile**
```dockerfile
FROM node:20-alpine
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 3000
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0", "--port", "3000"]
```

---

### Task 4: Backend - 資料庫連線與模型定義
**Objective:** 使用 GORM 連線 PostgreSQL 並定義 Stock 模型。

**Files:**
- Modify: `backend/main.go`
- Create: `backend/models/stock.go`

**Step 1: 定義模型**
```go
type Stock struct {
    ID    uint   `gorm:"primaryKey" json:"id"`
    Code  string `gorm:"uniqueIndex" json:"code"`
    Name  string `json:"name"`
    Price float64 `json:"price"`
}
```

---

### Task 5: 前後端串接與功能驗證
**Objective:** 前端調用後端 API 顯示股價資訊。

**Files:**
- Modify: `frontend/src/App.vue`

**Step 1: 實作 API 調用**
```javascript
const getStock = async (code) => {
  const res = await axios.get(`http://localhost:8080/api/stock/${code}`)
  stockData.value = res.data
}
```
