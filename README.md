# go-mcp-universal

Универсальный MCP сервер для Go-приложений с поддержкой Cursor.ai и AI-инструментов.

## Возможности
- Доступ к информации о горутинах и их стеках
- Получение переменных окружения
- Запуск тестов
- Диагностика pprof (heap, goroutine)
- AI-инструмент (заглушка, можно расширить)

## Установка

```bash
go get github.com/anddm2001/go-mcp-universal
```

## Использование в своём приложении

```go
import (
    "github.com/mark3labs/mcp-go/server"
    "github.com/anddm2001/go-mcp-universal/pkg/tools"
)

func main() {
    // ... инициализация вашего приложения

    // Запуск MCP сервера в отдельной горутине
    go func() {
        s := server.NewMCPServer("my-app", "0.1.0")
        tools.RegisterRuntimeTools(s)
        tools.RegisterEnvTools(s)
        tools.RegisterTestTools(s)
        tools.RegisterPprofTools(s)
        tools.RegisterAITools(s)
        server.ServeStdio(s)
    }()

    // ... основной цикл приложения
}
```

## Структура проекта

```
go-mcp-universal/
├── cmd/
│   └── universal/        # Пример использования как отдельного приложения
├── pkg/
│   └── tools/           # Публичные инструменты для использования в других приложениях
│       ├── runtime.go   # Инструменты для работы с горутинами
│       ├── env.go       # Инструменты для работы с переменными окружения
│       ├── tests.go     # Инструменты для запуска тестов
│       ├── pprof.go     # Инструменты для pprof
│       └── ai.go        # AI-инструменты
└── examples/            # Примеры использования
```

## Способы запуска

### 1. Отдельный процесс (standalone)

**Описание:**
MCP сервер запускается как отдельное приложение. Все инструменты работают только с этим процессом.

**Команда запуска:**
```sh
go run ./cmd/universal
```
или
```sh
go build -o mcp-server ./cmd/universal
./mcp-server
```

**Плюсы:**
- Просто запускать и тестировать
- Не влияет на основное приложение

**Минусы:**
- Доступ только к рантайму самого MCP сервера, а не вашего основного приложения

---

### 2. Встраивание в основное приложение (рекомендуется)

**Описание:**
MCP сервер запускается внутри основного приложения. Все инструменты работают с этим процессом — вы получаете доступ к реальному состоянию вашего приложения.

**Пример интеграции:**
```go
package main

import (
    "github.com/mark3labs/mcp-go/server"
    "github.com/anddm2001/go-mcp-universal/pkg/tools"
    // ... ваш код
)

func main() {
    // ... инициализация вашего приложения

    // Запуск MCP сервера в отдельной горутине
    go func() {
        s := server.NewMCPServer("my-app", "0.1.0")
        tools.RegisterRuntimeTools(s)
        tools.RegisterEnvTools(s)
        tools.RegisterTestTools(s)
        tools.RegisterPprofTools(s)
        tools.RegisterAITools(s)
        server.ServeStdio(s)
    }()

    // ... основной цикл приложения
}
```

**Плюсы:**
- Доступ к реальному состоянию приложения (горутины, pprof, env)
- Можно использовать в production (с ограничением доступа)

**Минусы:**
- Требует изменения кода приложения

---

### 3. Запуск дочернего процесса (расширенный вариант)

**Описание:**
MCP сервер запускает ваше приложение как дочерний процесс и может управлять его жизненным циклом, но не имеет доступа к его рантайму напрямую (только к stdout/stderr, переменным окружения на старте и т.д.).

**Пример (упрощённо):**
```go
// cmd/universal/main.go
package main

import (
    "os/exec"
    "log"
    "github.com/mark3labs/mcp-go/server"
    "github.com/anddm2001/go-mcp-universal/pkg/tools"
)

func main() {
    // Запуск дочернего процесса
    cmd := exec.Command("./your-app-binary")
    if err := cmd.Start(); err != nil {
        log.Fatalf("failed to start child: %v", err)
    }

    s := server.NewMCPServer("go-mcp-universal", "0.1.0")
    tools.RegisterRuntimeTools(s)
    tools.RegisterEnvTools(s)
    tools.RegisterTestTools(s)
    tools.RegisterPprofTools(s)
    tools.RegisterAITools(s)
    server.ServeStdio(s)

    // Ожидание завершения дочернего процесса
    cmd.Wait()
}
```

**Плюсы:**
- Можно запускать и контролировать отдельные приложения

**Минусы:**
- Нет доступа к рантайму дочернего процесса (горутины, pprof и т.д. будут показывать только MCP сервер)

---

## Интеграция с Cursor.ai

1. Запустите MCP сервер любым из способов выше.
2. В Cursor.ai выберите "Connect MCP server" и укажите путь к исполняемому файлу или настройте stdio-подключение.
3. Используйте инструменты MCP прямо из IDE.

---

## Расширение AI-инструмента

В файле `pkg/tools/ai.go` реализован заглушечный инструмент. Вы можете заменить его на интеграцию с любым LLM (OpenAI, локальный LLM и т.д.).

---

## Вопросы и поддержка

Если возникли вопросы по интеграции или расширению — создайте issue или напишите в обсуждение! 