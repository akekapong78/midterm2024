## Getting Started

1. Run Database

```bash
docker compose up -d
```

2. Run Backend

```bash
cd backend
go mod tidy
go run cmd/app/main.go
```

3. Run Frontend

```bash
cd frontend
npm i
npm run dev 
```

# Project tree
.
├── README.md
├── compose.yml
├── .gitignore
├── .env
├── backend
│   ├── cmd
│   ├── go.mod
│   ├── go.sum
│   ├── internal
│   └── migrations
└── frontend
    ├── README.md
    ├── app
    ├── next-env.d.ts
    ├── next.config.mjs
    ├── node_modules
    ├── package-lock.json
    ├── package.json
    ├── postcss.config.mjs
    ├── tailwind.config.ts
    └── tsconfig.json