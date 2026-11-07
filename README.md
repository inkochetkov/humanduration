# humanduration

Minimal, dependency-free Go package for formatting `time.Duration` into human-readable strings — for logs, CLIs, and user-facing output.

## Why

`time.Duration.String()` gives you `"24h0m0s"`. Fine for machines, noisy for humans.

```go
d := 25 * time.Hour

d.String()                                 // "25h0m0s"
humanduration.Format(d)                    // "1 day 1 hour"
humanduration.FormatShort(d)               // "1d 1h"
```

## Install

```bash
go get github.com/inkochetkov/humanduration
```

Requires Go 1.18+. No dependencies outside the standard library.

## Usage

```go
package main

import (
	"fmt"
	"time"

	"github.com/inkochetkov/humanduration"
)

func main() {
	fmt.Println(humanduration.Format(90 * time.Second))         // 1 minute 30 seconds
	fmt.Println(humanduration.Format(25 * time.Hour))           // 1 day 1 hour
	fmt.Println(humanduration.Format(9 * 24 * time.Hour))       // 1 week 2 days

	fmt.Println(humanduration.FormatN(90*time.Second, 1))       // 1 minute
	fmt.Println(humanduration.FormatShort(90 * time.Second))    // 1m 30s
	fmt.Println(humanduration.FormatShort(500 * time.Millisecond)) // 500ms
}
```

### API

| Function | Description |
|---|---|
| `Format(d time.Duration) string` | Long form, up to 2 units. |
| `FormatN(d time.Duration, n int) string` | Long form, up to `n` units (`n < 1` treated as `1`). |
| `FormatShort(d time.Duration) string` | Compact form (`1h 30m`, `2d`, `500ms`). |

### Units

`week` · `day` · `hour` · `minute` · `second` · `millisecond`

Longest unit wins (`1 week`, not `7 days`).

### Edge cases

| Input | Output |
|---|---|
| `0` | `"0 seconds"` / `"0s"` |
| negative | prefixed with `"-"` |
| `< 1ms` | `"0 seconds"` / `"0s"` |

Never returns an empty string — safe to embed in logs, JSON, templates.

## Design

- One job, three functions, zero deps.
- No reflection, no `init()`.
- Stable API: once `Format(d) string` is shipped, it stays.
- Formatting only. Parsing is intentionally out of scope (`time.ParseDuration` exists).

## License

MIT — see [LICENSE](LICENSE).