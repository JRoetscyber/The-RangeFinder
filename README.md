# The RangeFinder

> *"I consider that our present sufferings are not worth comparing with the glory that will be revealed in us."* — Romans 8:18

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Category](https://img.shields.io/badge/Category-Broken%20Access%20Control%20%2F%20IDOR-red?style=flat)](#)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

**The RangeFinder** is a high-velocity parameter fuzzer and anomaly detector written in Go. It is specifically designed to uncover Insecure Direct Object References (IDORs) and horizontal/vertical access control flaws across web applications and REST APIs.

---

## How It Works

1. **Baseline Fingerprinting**: The tool first requests a known-invalid ID (`-baseline`) to establish a reference point of the target's "not found" or "forbidden" behavior (capturing status code, content length, and raw response bytes).
2. **Range Traversal**: It iterates through your target range (`-start` to `-end`), swapping `{ID}` in the URL template across a pool of concurrent workers.
3. **Deep Differential Evaluation**: Unlike naive tools that only check HTTP status codes, The RangeFinder performs content-level differential analysis (`bytes.Equal`). This catches soft-404s, 200 OK error responses, and template changes where the status code remains identical.
4. **Session Authentication**: Passes active session credentials (`-cookie`) to audit authenticated endpoints and private user resources.
5. **Memory Guard**: Uses Go's `io.LimitReader` (capped at 4MB) to safely handle large downloads or media endpoints without memory exhaustion.

---

## Installation & Quick Start

### Build from source

```bash
git clone https://github.com/JRoetscyber/The-RangeFinder.git
cd The-RangeFinder
go build -o rangefinder.exe main.go
```

### Run directly with `go run`

```bash
# Basic scan
go run main.go -url "http://127.0.0.1:8080/download-transcript/{ID}.txt" -start 1 -end 50 -baseline 999999

# Authenticated concurrent scan (10 workers)
go run main.go -url "https://target.com/api/users/{ID}" -cookie "session=abc123xyz" -start 1 -end 1000 -workers 10 -baseline 999999
```

---

## CLI Flags

| Flag | Type | Default | Description |
| :--- | :--- | :--- | :--- |
| `-url` | `string` | *(required)* | Target URL template containing the `{ID}` placeholder |
| `-cookie` | `string` | `""` | Session Cookie string for authenticated requests |
| `-workers` | `int` | `2` | Number of concurrent workers (goroutines) |
| `-start` | `int` | `1` | Start of the numeric ID range |
| `-end` | `int` | `100` | End of the numeric ID range (inclusive) |
| `-baseline` | `string` | `"999999"` | Known-invalid ID used to calibrate baseline response |
| `-v` | `bool` | `false` | Enable verbose output |

---

## Example Output

```text
{Hit} Found ID: 1 Status: 200 Length: 60
{Hit} Found ID: 2 Status: 200 Length: 61
{Hit} Found ID: 3 Status: 200 Length: 56
{Hit} Found ID: 42 Status: 200 Length: 59
```

---

## Included Mock Server

A built-in test laboratory replicating real-world IDOR scenarios (e.g. PortSwigger Web Security Academy chat transcript labs) is included under `mock_server/`.

To launch the test lab:

```bash
cd mock_server
go run main.go
```

The server listens on `http://127.0.0.1:8080` with valid transcripts at IDs `1`, `2`, `3`, and `42`, returning 404 for all other IDs.

---

## Roadmap

- [x] **v1.0**: Core CLI, `{ID}` templating, baseline health-check, dual-criteria evaluation (status + body bytes), ANSI color output.
- [ ] **v2.0**:
  - [x] Custom session authentication (`-cookie` flag).
  - [x] Concurrent worker pools (`-workers` flag with goroutines + channels).
  - [ ] Adaptive rate limiting with exponential backoff on 429/503 (`-rps`).
  - [ ] Wordlist / slug mode (`-wordlist` for usernames & UUIDs).
  - [ ] Structured JSONL export.
- [ ] **v3.0**: Burp Suite proxy routing, response similarity scoring (Jaccard index), multi-endpoint sweeps, session fingerprinting / User-Agent matching.

---

## Legal Disclaimer & Limitation of Liability

This tool is designed, developed, and released strictly for educational purposes, defensive security auditing, and authorized penetration testing on systems where the tester has received prior, explicit, written permission from the system owner.

1. **No Liability**: IN NO EVENT SHALL THE AUTHOR(S), CONTRIBUTORS, OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES, LOSSES, OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT, OR OTHERWISE, ARISING FROM, OUT OF, OR IN CONNECTION WITH THE SOFTWARE, ITS CODE, OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
2. **User Responsibility**: The user assumes all legal liability, civil accountability, and criminal responsibility for their actions. It is the sole responsibility of the end user to ensure compliance with all applicable local, national, and international laws, regulations, and program policies (including computer misuse and cybercrime legislation).
3. **Prohibited Use**: Any use of this software against targets without explicit written authorization is strictly prohibited. The author expressly disclaims any responsibility for misuse or damage caused by this program.

---

## License

This project is licensed under the terms of the [MIT License](LICENSE).
