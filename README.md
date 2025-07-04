# Valuta

**Valuta** is a CLI tool for currency conversion and exchange rate listing. It provides a convenient way to get up-to-date exchange rates and perform conversions directly from your terminal.

---

## ✨ Features

- 📈 List exchange rates for the Georgian Lari (GEL)
- 💱 Convert an amount from one currency to another

---

### Prerequisites
- Go 1.24+ ([installation guide](https://go.dev/doc/install))

### Option 1: Install via Go:
```bash
go install github.com/georgedolee/valuta@latest
```

### Option 2: Build from source:

```bash
git clone git@github.com:georgedolee/valuta.git
cd amindi
go build
```

## Usage

```bash
amindi [location] [flags]
```

### List Exchange Rates

```bash
valuta rates
```

**Flags:**
| Flag                  | Description                            | Default            |
|-----------------------|----------------------------------------|--------------------|
| `-c`, `--currencies`  | Comma-separated list of currency codes | "USD,EUR,RUB,CNY"  |

### Convert Currency

```bash
valuta convert <amount> <from> <to>
```

## Dependencies

- [cobra](https://github.com/spf13/cobra) – CLI framework
- [go-pretty](https://github.com/jedib0t/go-pretty) – Table formatting
- [fatih/color](https://github.com/fatih/color) – Terminal color output

## License

This project is licensed under the [MIT License](./LICENSE).