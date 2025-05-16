# 奇門遁甲問答系統

這是一個基於奇門遁甲的問答系統，提供用戶根據生辰八字和時間進行命理諮詢的服務。

## 專案結構

```
.
├── mobile/           # Flutter 移動端應用
├── backend/          # Golang 後端服務
└── docs/            # 專案文檔
```

## 技術棧

- 前端：Flutter (iOS & Android)
- 後端：Golang
- 資料庫：PostgreSQL (Railway)
- 部署：Railway
- 支付系統：待定

## 功能特點

- 用戶可以消耗虛擬貨幣提出問題
- 系統根據用戶生辰和當前時間生成奇門盤
- 提供專業的命理分析和建議
- 支援 iOS 和 Android 雙平台
- 整合支付系統

## 開發環境設置

### 移動端 (Flutter)

1. 安裝 Flutter SDK
2. 安裝必要的開發工具（Android Studio / Xcode）
3. 進入 mobile 目錄
4. 運行 `flutter pub get`
5. 運行 `flutter run`

### 後端 (Golang)

1. 安裝 Go 1.21 或更高版本
2. 進入 backend 目錄
3. 運行 `go mod download`
4. 運行 `go run main.go`

## 部署

### Railway 部署

1. 在 Railway 創建新專案
2. 連接 GitHub 倉庫
3. 設置環境變數
4. 部署後端服務
5. 設置資料庫

## 上架流程

### iOS App Store

1. 註冊 Apple Developer 帳號
2. 準備 App Store Connect 所需資料
3. 提交審核

### Google Play Store

1. 註冊 Google Play Developer 帳號
2. 準備 Play Console 所需資料
3. 提交審核

## 收費功能

- 虛擬貨幣購買系統
- 問題提問計費
- 支付系統整合

## 開發團隊

- 前端開發
- 後端開發
- UI/UX 設計
- 命理顧問

## 授權

版權所有 © 2024 