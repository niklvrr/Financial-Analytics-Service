# Financial Analytics Service

Система учета финансов с поддержкой банковских счетов, категорий и операций. Реализована на Go с использованием многослойной архитектуры и паттернов проектирования.

## Содержание

1. [Общая идея решения](#общая-идея-решения)
2. [Принципы SOLID](#принципы-solid)
3. [Принципы GRASP](#принципы-grasp)
4. [Паттерны GoF](#паттерны-gof)
5. [Инструкция по запуску](#инструкция-по-запуску)

---

## Общая идея решения

### Реализованный функционал

Проект представляет собой CLI-приложение для управления финансовыми данными с поддержкой следующих операций:

#### Управление банковскими счетами
- Создание, чтение, обновление, удаление счетов
- Получение списка всех счетов
- Импорт счетов из файлов (CSV, JSON, YAML)
- **Экспорт счетов в файлы (CSV, JSON, YAML)** — новая функциональность

#### Управление категориями
- CRUD операции для категорий
- Импорт категорий из файлов
- **Экспорт категорий в файлы** — новая функциональность

#### Управление операциями
- CRUD операции для финансовых операций
- Импорт операций из файлов
- **Экспорт операций в файлы с различными стратегиями фильтрации** — новая функциональность:
  - Полный экспорт всех операций
  - Экспорт по ID банковского счета
  - Экспорт по ID категории
  - Экспорт за указанный период (date range)

### Архитектура решения

Проект построен на основе **многослойной архитектуры (Layered Architecture)**:

```
┌─────────────────────────────────────┐
│   Transport Layer (CLI)             │
│   - Команды (Command pattern)      │
│   - Меню (Menu system)              │
└─────────────────────────────────────┘
              ↓
┌─────────────────────────────────────┐
│   Use Case Layer                    │
│   - Сервисы (Business Logic)         │
│   - Фасады (Facade pattern)         │
│   - Импортеры (Template Method)     │
│   - Экспортеры (Builder + Strategy) │
└─────────────────────────────────────┘
              ↓
┌─────────────────────────────────────┐
│   Domain Layer                      │
│   - Модели (Domain Models)          │
│   - Request/Response DTOs           │
│   - Фабрика (Factory pattern)       │
└─────────────────────────────────────┘
              ↓
┌─────────────────────────────────────┐
│   Infrastructure Layer              │
│   - Репозитории (Repository)        │
│   - Прокси (Proxy pattern)          │
│   - База данных (PostgreSQL)        │
└─────────────────────────────────────┘
```

### Ключевые изменения и дополнения

1. **Система экспорта данных** — реализована с использованием паттернов Builder и Strategy:
   - Builder для построения отчетов в различных форматах (CSV, JSON, YAML)
   - Strategy для выбора стратегии фильтрации данных при экспорте
   - Названия полей в экспортируемых файлах соответствуют полям из `internal/domain/request`

2. **DI-контейнер** — централизованная инициализация зависимостей в `internal/app/di_container.go`

3. **Система декораторов** — добавление кросс-срезочных аспектов (логирование) к командам

---

## Принципы SOLID

### 1. Single Responsibility Principle (SRP)

**Каждый класс имеет одну причину для изменения.**

**Реализация:**
- **`internal/usecase/service/*_service.go`** — сервисы отвечают только за бизнес-логику конкретной сущности
- **`internal/infrastructure/repository/*_repo.go`** — репозитории отвечают только за работу с БД
- **`internal/transport/cli/command/*_commands/*.go`** — каждая команда отвечает за один тип операции
- **`internal/usecase/exporter/*_builder.go`** — каждый builder отвечает за форматирование в одном формате
- **`internal/usecase/exporter/*_strategy.go`** — каждая стратегия отвечает за один способ фильтрации данных

**Пример:**
```go
// internal/usecase/exporter/csv_builder.go
type CSVBuilder struct {
    rows   [][]string
    header []string
}
// Отвечает только за построение CSV формата
```

### 2. Open/Closed Principle (OCP)

**Классы открыты для расширения, но закрыты для модификации.**

**Реализация:**
- **`internal/usecase/exporter/builder.go`** — интерфейс `Builder` позволяет добавлять новые форматы (XML, Excel) без изменения существующего кода
- **`internal/usecase/exporter/strategy.go`** — интерфейс `ExportStrategy` позволяет добавлять новые стратегии фильтрации без изменения фасадов
- **`internal/transport/cli/menu/menu.go`** — система меню расширяется добавлением новых команд без изменения базовой логики

**Пример:**
```go
// Можно добавить новый формат без изменения существующих
func NewBuilder(format string) (Builder, error) {
    switch format {
    case ".csv":
        return NewCSVBuilder(), nil
    case ".json":
        return NewJSONBuilder(), nil
    case ".yaml":
        return NewYAMLBuilder(), nil
    // Легко добавить новый формат здесь
    }
}
```

### 3. Liskov Substitution Principle (LSP)

**Объекты должны быть заменяемы экземплярами их подтипов без изменения корректности программы.**

**Реализация:**
- **`internal/usecase/service/*_service.go`** — все сервисы реализуют интерфейсы репозиториев, определенные в сервисах
- **`internal/infrastructure/proxy/*_proxy.go`** — прокси полностью заменяют репозитории, сохраняя тот же интерфейс
- **`internal/usecase/exporter/*_builder.go`** — все builders взаимозаменяемы через интерфейс `Builder`

**Пример:**
```go
// internal/usecase/service/bank_account_service.go
type BankAccountRepo interface {
    CreateBankAccount(ctx context.Context, account *model.BankAccount) error
    // ...
}

// internal/infrastructure/proxy/bank_account_proxy.go
// Proxy реализует тот же интерфейс и может заменить репозиторий
```

### 4. Interface Segregation Principle (ISP)

**Клиенты не должны зависеть от интерфейсов, которые они не используют.**

**Реализация:**
- **`internal/usecase/exporter/strategy.go`** — интерфейс `ExportStrategy` содержит только необходимые методы (`Collect`, `GetHeaders`)
- **`internal/transport/cli/menu/menu.go`** — интерфейс `Command` содержит только `Execute` и `Title`
- Разделение интерфейсов репозиториев по типам сущностей (BankAccountRepo, CategoryRepo, OperationRepo)

**Пример:**
```go
// Минимальный интерфейс для стратегии
type ExportStrategy interface {
    Collect(ctx context.Context, params ExportParams) ([]map[string]string, error)
    GetHeaders() []string
}
```

### 5. Dependency Inversion Principle (DIP)

**Зависимости должны быть направлены на абстракции, а не на конкретные реализации.**

**Реализация:**
- **`internal/usecase/service/*_service.go`** — сервисы зависят от интерфейсов репозиториев, а не от конкретных реализаций
- **`internal/app/di_container.go`** — DI-контейнер инвертирует зависимости, создавая конкретные реализации и передавая их через интерфейсы
- **`internal/usecase/facade/*_facade.go`** — фасады зависят от интерфейсов сервисов

**Пример:**
```go
// internal/usecase/service/bank_account_service.go
type BankAccountService struct {
    repo   BankAccountRepo  // Зависимость от интерфейса, а не от конкретного типа
    fabric *model.DomainFabric
}
```

---

## Принципы GRASP

### 1. High Cohesion (Высокая связность)

**Элементы внутри модуля должны быть тесно связаны и работать вместе для достижения одной цели.**

**Реализация:**
- **`internal/usecase/exporter/`** — все классы в пакете exporter работают вместе для экспорта данных
- **`internal/transport/cli/command/bank_account_commands/`** — все команды для банковских счетов сгруппированы вместе
- **`internal/infrastructure/repository/`** — все репозитории сгруппированы по функциональности работы с БД

**Пример:**
```go
// internal/usecase/exporter/
// Все файлы работают вместе для экспорта:
// - builder.go - интерфейс Builder
// - csv_builder.go, json_builder.go, yaml_builder.go - реализации
// - strategy.go - интерфейс Strategy
// - full_export_strategy.go, by_account_strategy.go и т.д. - реализации
```

### 2. Low Coupling (Низкая связанность)

**Модули должны быть слабо связаны друг с другом, изменения в одном модуле не должны влиять на другие.**

**Реализация:**
- **`internal/usecase/service/`** — сервисы не знают о деталях реализации репозиториев (работают через интерфейсы)
- **`internal/transport/cli/command/`** — команды не знают о деталях работы сервисов (работают через фасады)
- **`internal/usecase/facade/`** — фасады изолируют транспортный слой от деталей бизнес-логики

**Пример:**
```go
// Команда не знает о деталях экспорта, работает через фасад
type ExportBankAccountCommand struct {
    f *facade.BankAccountFacade  // Зависимость от абстракции
}
```

### 3. Information Expert (Информационный эксперт)

**Ответственность должна быть назначена классу, который имеет информацию, необходимую для выполнения задачи.**

**Реализация:**
- **`internal/infrastructure/repository/*_repo.go`** — репозитории знают, как работать с БД
- **`internal/usecase/exporter/*_strategy.go`** — стратегии знают, как фильтровать данные для экспорта
- **`internal/domain/model/fabric.go`** — фабрика знает, как создавать валидные доменные объекты

**Пример:**
```go
// internal/usecase/exporter/by_account_strategy.go
// Стратегия знает, как фильтровать операции по account ID
func (s *ByAccountStrategy) Collect(ctx context.Context, params ExportParams) ([]map[string]string, error) {
    // Логика фильтрации находится здесь, где есть доступ к данным
}
```

### 4. Creator (Создатель)

**Класс, который создает экземпляры другого класса, должен иметь для этого веские причины.**

**Реализация:**
- **`internal/domain/model/fabric.go`** — `DomainFabric` создает доменные объекты с валидацией
- **`internal/app/di_container.go`** — DI-контейнер создает все зависимости приложения
- **`internal/usecase/exporter/builder.go`** — `NewBuilder` создает конкретные реализации builders

**Пример:**
```go
// internal/domain/model/fabric.go
// Фабрика создает доменные объекты, т.к. знает правила валидации
func (f *DomainFabric) BuildBankAccount(id int64, name string, balance float64) (*BankAccount, error) {
    // Валидация и создание
    return NewBankAccount(id, name, balance), nil
}
```

### 5. Controller (Контроллер)

**Класс, который координирует работу системы, но не выполняет бизнес-логику.**

**Реализация:**
- **`internal/transport/cli/menu/menu.go`** — `Menu` координирует выполнение команд
- **`internal/app/di_container.go`** — `Container` координирует инициализацию всех компонентов
- **`internal/usecase/facade/*_facade.go`** — фасады координируют работу сервисов

**Пример:**
```go
// internal/transport/cli/menu/menu.go
// Menu координирует выполнение команд, но не содержит бизнес-логику
func (m *Menu) Run(ctx context.Context) {
    // Координация выполнения команд
}
```

---

## Паттерны GoF

### 1. Builder (Строитель)

**Назначение:** Пошаговое построение сложных объектов.

**Реализация:** `internal/usecase/exporter/`

**Важность:** Позволяет создавать отчеты в различных форматах (CSV, JSON, YAML) с единым интерфейсом. Упрощает добавление новых форматов без изменения существующего кода.

**Классы:**
- `internal/usecase/exporter/builder.go` — интерфейс `Builder`
- `internal/usecase/exporter/csv_builder.go` — `CSVBuilder`
- `internal/usecase/exporter/json_builder.go` — `JSONBuilder`
- `internal/usecase/exporter/yaml_builder.go` — `YAMLBuilder`

**Пример использования:**
```go
builder, _ := exporter.NewBuilder(".csv")
builder.Begin(ctx, "Отчет")
builder.AddHeader(ctx, "id", "name", "balance")
builder.AddRow(ctx, "1", "Account1", "1000.00")
report, _ := builder.End(ctx)
```

### 2. Strategy (Стратегия)

**Назначение:** Определяет семейство алгоритмов, инкапсулирует каждый из них и делает их взаимозаменяемыми.

**Реализация:** `internal/usecase/exporter/`

**Важность:** Позволяет выбирать различные стратегии фильтрации данных при экспорте (полный экспорт, по счету, по категории, по датам) без изменения кода фасадов и команд.

**Классы:**
- `internal/usecase/exporter/strategy.go` — интерфейс `ExportStrategy`
- `internal/usecase/exporter/full_export_strategy.go` — `FullExportStrategy`
- `internal/usecase/exporter/by_account_strategy.go` — `ByAccountStrategy`
- `internal/usecase/exporter/by_category_strategy.go` — `ByCategoryStrategy`
- `internal/usecase/exporter/date_range_strategy.go` — `DateRangeStrategy`

**Пример использования:**
```go
strategy, _ := exporter.NewStrategy("by_account", "operation")
strategy.SetService(operationService)
data, _ := strategy.Collect(ctx, params)
```

### 3. Factory (Фабрика)

**Назначение:** Создание объектов без указания конкретных классов.

**Реализация:** `internal/domain/model/fabric.go`, `internal/usecase/exporter/builder.go`, `internal/usecase/exporter/strategy.go`

**Важность:** Централизует создание объектов с валидацией и обеспечивает единообразие создания экземпляров.

**Классы:**
- `internal/domain/model/fabric.go` — `DomainFabric` (создание доменных объектов)
- `internal/usecase/exporter/builder.go` — `NewBuilder` (создание builders)
- `internal/usecase/exporter/strategy.go` — `NewStrategy` (создание стратегий)

**Пример:**
```go
// internal/domain/model/fabric.go
func (f *DomainFabric) BuildBankAccount(id int64, name string, balance float64) (*BankAccount, error) {
    // Валидация и создание
}
```

### 4. Repository (Репозиторий)

**Назначение:** Абстракция доступа к данным, скрывает детали работы с БД.

**Реализация:** `internal/infrastructure/repository/`

**Важность:** Изолирует бизнес-логику от деталей персистентности, упрощает тестирование и позволяет менять источник данных.

**Классы:**
- `internal/infrastructure/repository/bank_account_repo.go` — `BankAccountRepo`
- `internal/infrastructure/repository/category_repo.go` — `CategoryRepo`
- `internal/infrastructure/repository/operation_repo.go` — `OperationRepo`

**Интерфейсы определены в:**
- `internal/usecase/service/bank_account_service.go` — `BankAccountRepo interface`
- `internal/usecase/service/category_service.go` — `CategoryRepo interface`
- `internal/usecase/service/operation_service.go` — `OperationRepo interface`

### 5. Proxy (Прокси)

**Назначение:** Предоставляет суррогат или заместитель для другого объекта для контроля доступа к нему.

**Реализация:** `internal/infrastructure/proxy/`

**Важность:** Добавляет кэширование для оптимизации доступа к данным без изменения интерфейса репозитория.

**Классы:**
- `internal/infrastructure/proxy/bank_account_proxy.go` — `BankAccountProxy`
- `internal/infrastructure/proxy/category_proxy.go` — `CategoryProxy`
- `internal/infrastructure/proxy/operation_proxy.go` — `OperationProxy`

**Пример:**
```go
// Прокси добавляет кэширование поверх репозитория
func (p *BankAccountProxy) GetBankAccount(ctx context.Context, accountId int64) (*model.BankAccount, error) {
    if account, ok := p.cache[accountId]; ok {
        return account, nil  // Возврат из кэша
    }
    // Иначе запрос к репозиторию
}
```

### 6. Template Method (Шаблонный метод)

**Назначение:** Определяет скелет алгоритма, делегируя некоторые шаги подклассам.

**Реализация:** `internal/usecase/importer/*/`

**Важность:** Унифицирует процесс импорта (Load → Validate → Save) для всех форматов, позволяя изменять только детали парсинга.

**Классы:**
- `internal/usecase/importer/bank_account_importer/bank_account_importer.go` — `BankAccountTemplate`
- `internal/usecase/importer/category_importer/category_importer.go` — `CategoryTemplate`
- `internal/usecase/importer/operation_importer/operation_importer.go` — `OperationTemplate`

**Пример:**
```go
// Шаблонный метод определяет алгоритм
func (t *BankAccountTemplate) Run(ctx context.Context, path string) error {
    data, err := t.Impl.Load(path)      // Шаг 1
    err = t.Impl.Validate(data)        // Шаг 2
    err = t.Impl.Save(ctx, data)       // Шаг 3
    return nil
}
```

### 7. Command (Команда)

**Назначение:** Инкапсулирует запрос как объект, позволяя параметризовать клиентов с различными запросами.

**Реализация:** `internal/transport/cli/command/`

**Важность:** Позволяет унифицировать выполнение операций, добавлять декораторы (логирование, таймауты) и поддерживать отмену операций.

**Классы:**
- `internal/transport/cli/menu/menu.go` — интерфейс `Command`
- `internal/transport/cli/command/bank_account_commands/*.go` — команды для счетов
- `internal/transport/cli/command/category_commands/*.go` — команды для категорий
- `internal/transport/cli/command/operation_commands/*.go` — команды для операций

**Пример:**
```go
type Command interface {
    Execute(ctx context.Context) error
    Title() string
}
```

### 8. Decorator (Декоратор)

**Назначение:** Динамически добавляет объектам новую функциональность.

**Реализация:** `internal/transport/cli/command/decorator/`

**Важность:** Позволяет добавлять кросс-срезочные аспекты (логирование, измерение времени, таймауты) к командам без изменения их кода.

**Классы:**
- `internal/transport/cli/command/decorator/command_decorator.go` — `LoggingDecorator`

**Пример:**
```go
// Декоратор добавляет логирование к команде
decoratedCommand := decorator.WithLogging(originalCommand, logger)
```

### 9. Facade (Фасад)

**Назначение:** Предоставляет унифицированный интерфейс к набору интерфейсов в подсистеме.

**Реализация:** `internal/usecase/facade/`

**Важность:** Упрощает работу транспортного слоя, скрывая сложность взаимодействия с несколькими сервисами и координируя операции экспорта/импорта.

**Классы:**
- `internal/usecase/facade/bank_account_facade.go` — `BankAccountFacade`
- `internal/usecase/facade/category_facade.go` — `CategoryFacade`
- `internal/usecase/facade/operation_facade.go` — `OperationFacade`

**Пример:**
```go
// Фасад упрощает экспорт, координируя стратегию и builder
func (f *BankAccountFacade) Export(ctx context.Context, params exporter.ExportParams) (*exporter.Report, error) {
    strategy, _ := exporter.NewStrategy(params.Strategy, "bank_account")
    data, _ := strategy.Collect(ctx, params)
    builder, _ := exporter.NewBuilder(params.Format)
    // ... построение отчета
}
```

---

## Инструкция по запуску

### Требования

- Go 1.25.1 или выше
- PostgreSQL 12+
- Установленный `golang-migrate` (опционально, миграции запускаются автоматически)

### Шаг 1: Клонирование репозитория

```bash
git clone <repository-url>
cd Financial-Analytics-Service
```

### Шаг 2: Настройка базы данных

1. Создайте базу данных PostgreSQL:
```sql
CREATE DATABASE financial_analytics;
```

2. Создайте файл `.env` в корне проекта:
```env
APP_ENV=local
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=financial_analytics
```

### Шаг 3: Установка зависимостей

```bash
go mod download
```

### Шаг 4: Запуск приложения

```bash
go run cmd/main.go
```

Или скомпилируйте и запустите:

```bash
go build -o financial-analytics cmd/main.go
./financial-analytics
```

### Шаг 5: Использование

После запуска откроется интерактивное меню:

```
=== Главное меню ===
1) Управлять банковскими счетами
2) Управлять категориями
3) Управлять операциями
4) Выход
```

В каждом подменю доступны операции:
- Создание, чтение, обновление, удаление
- Получение всех записей
- Импорт из файла (CSV/JSON/YAML)
- **Экспорт в файл (CSV/JSON/YAML)** — новый функционал

### Примеры использования экспорта

#### Экспорт банковских счетов:
1. Выберите "Управлять банковскими счетами"
2. Выберите "Экспорт банковских счетов в файл"
3. Введите путь к файлу (например: `/path/to/accounts.csv`)
4. Формат определяется автоматически по расширению файла

#### Экспорт операций с фильтрацией:
1. Выберите "Управлять операциями"
2. Выберите "Экспорт операций в файл"
3. Выберите стратегию:
   - `full` — все операции
   - `by_account` — по ID счета (потребуется ввести ID)
   - `by_category` — по ID категории (потребуется ввести ID)
   - `date_range` — за период (потребуется ввести даты)
4. Введите путь к файлу

### Форматы экспорта

Экспорт поддерживает три формата:
- **CSV** — табличный формат с заголовками
- **JSON** — массив объектов с отступами
- **YAML** — структурированный формат

Названия полей в экспортируемых файлах соответствуют полям из `internal/domain/request`:
- Банковские счета: `id`, `name`, `balance`
- Категории: `id`, `kind`, `name`
- Операции: `id`, `kind`, `bank_account_id`, `amount`, `date`, `description`, `category_id`

### Структура проекта

```
.
├── cmd/
│   └── main.go                    # Точка входа приложения
├── internal/
│   ├── app/
│   │   ├── app.go                 # Основной класс приложения
│   │   └── di_container.go       # DI-контейнер
│   ├── config/
│   │   └── config.go              # Конфигурация
│   ├── domain/
│   │   ├── model/                 # Доменные модели
│   │   ├── request/               # DTO для запросов
│   │   └── response/              # DTO для ответов
│   ├── infrastructure/
│   │   ├── repository/            # Репозитории (Repository pattern)
│   │   ├── proxy/                 # Прокси с кэшированием (Proxy pattern)
│   │   └── db.go                  # Подключение к БД
│   ├── transport/
│   │   └── cli/
│   │       ├── command/           # Команды (Command pattern)
│   │       │   ├── decorator/     # Декораторы (Decorator pattern)
│   │       │   └── *_commands/    # Команды по типам сущностей
│   │       └── menu/              # Система меню
│   └── usecase/
│       ├── service/                # Бизнес-логика
│       ├── facade/                 # Фасады (Facade pattern)
│       ├── importer/               # Импортеры (Template Method)
│       └── exporter/               # Экспортеры (Builder + Strategy)
├── migrations/                     # SQL миграции
├── pkg/
│   ├── logger/                     # Логирование
│   └── utils/                      # Утилиты
└── go.mod                          # Зависимости
```

### Остановка приложения

Для корректного завершения нажмите `Ctrl+C` или выберите "Выход" в меню. Приложение выполнит graceful shutdown и закроет соединения с БД.

---

## Заключение

Проект демонстрирует применение современных принципов проектирования и паттернов GoF для создания масштабируемого и поддерживаемого приложения. Все компоненты слабо связаны, что позволяет легко расширять функциональность и тестировать отдельные части системы.



