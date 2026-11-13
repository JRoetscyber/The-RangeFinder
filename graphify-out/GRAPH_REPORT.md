# Graph Report - The RangeFinder  (2026-11-12)

## Corpus Check
- Corpus is ~4,113 words - fits in a single context window. You may not need a graph.

## Summary
- 50 nodes · 69 edges · 12 communities (8 shown, 4 thin omitted)
- Extraction: 86% EXTRACTED · 14% INFERRED · 0% AMBIGUOUS · INFERRED: 10 edges (avg confidence: 0.9)
- Token cost: 0 input · 0 output

## Community Hubs (Navigation)
- v2 Auth & Scaling Roadmap
- v3 Production Architecture
- Mock Target Environment
- v3 Production Architecture
- Core Engine & Scanner CLI
- Core Engine & Scanner CLI
- Core Engine & Scanner CLI
- Core Engine & Scanner CLI
- Core Engine & Scanner CLI
- Core Engine & Scanner CLI
- Package main (The RangeFinder)
- Legal Disclaimer

## God Nodes (most connected - your core abstractions)
1. `v2 Specification` - 10 edges
2. `worker() (Goroutine)` - 8 edges
3. `v3 Specification` - 8 edges
4. `main.go (Entrypoint & CLI)` - 7 edges
5. `Target.fetch()` - 7 edges
6. `main() (Orchestrator)` - 6 edges
7. `The RangeFinder` - 6 edges
8. `Limiter (Struct)` - 5 edges
9. `v1 Specification` - 5 edges
10. `Session Fingerprinting Gotchas` - 5 edges

## Surprising Connections (you probably didn't know these)
- `Target.fetch()` --implements--> `Memory Guard (4MB Cap)`  [EXTRACTED]
  main.go → README.md
- `Limiter (Struct)` --implements--> `Token-Bucket Rate Limiter`  [EXTRACTED]
  main.go → v2.md
- `Baseline.Evaluate()` --implements--> `Body Comparison (bytes.Equal)`  [EXTRACTED]
  main.go → v1.md
- `Target.fetch()` --implements--> `Session Authentication`  [EXTRACTED]
  main.go → README.md
- `worker() (Goroutine)` --implements--> `Concurrent Worker Pool`  [EXTRACTED]
  main.go → v2.md

## Import Cycles
- None detected.

## Communities (12 total, 4 thin omitted)

### Community 0 - "v2 Auth & Scaling Roadmap"
Cohesion: 0.28
Nodes (9): Deep Differential Evaluation, Body Comparison (bytes.Equal), v1 Specification, URL Template ({ID} Placeholder), Adaptive Backoff (429/503), Fetcher (http.NewRequest + Headers), Token-Bucket Rate Limiter, Jaccard Similarity Scoring (+1 more)

### Community 1 - "v3 Production Architecture"
Cohesion: 0.22
Nodes (9): Structured JSONL Output, Slug/Wordlist Mode, YAML/JSON Config File, CSV/JSON Reporting, Pluggable ID Generators, Clean Package Architecture, Multi-Endpoint Sweeps, Burp/ZAP Proxy Routing (+1 more)

### Community 2 - "Mock Target Environment"
Cohesion: 0.29
Nodes (5): main() (Orchestrator), Memory Guard (4MB Cap), Range Traversal, The RangeFinder, CLI Flag Parsing

### Community 3 - "v3 Production Architecture"
Cohesion: 0.33
Nodes (6): Session Authentication, Auth: Cookies & Bearer Tokens, IP Address Binding, Session Fingerprinting Gotchas, TLS / Device Bound Credentials, User-Agent Binding

### Community 4 - "Core Engine & Scanner CLI"
Cohesion: 0.50
Nodes (4): sync.Mutex, time.Duration, time.Time, Limiter (Struct)

### Community 5 - "Core Engine & Scanner CLI"
Cohesion: 0.67
Nodes (3): net/http.Client, Target (Struct), Target.fetch()

### Community 6 - "Core Engine & Scanner CLI"
Cohesion: 0.67
Nodes (3): sync.WaitGroup, worker() (Goroutine), Concurrent Worker Pool

### Community 7 - "Core Engine & Scanner CLI"
Cohesion: 0.67
Nodes (3): Baseline Fingerprinting, Baseline (Struct), Baseline Health Check

## Knowledge Gaps
- **13 isolated node(s):** `Package main (The RangeFinder)`, `Evaluator (Interface)`, `Range Traversal`, `Legal Disclaimer`, `URL Template ({ID} Placeholder)` (+8 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 18 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **4 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `v2 Specification` connect `v2 Auth & Scaling Roadmap` to `v3 Production Architecture`, `v3 Production Architecture`, `Core Engine & Scanner CLI`?**
  _High betweenness centrality (0.415) - this node is a cross-community bridge._
- **Why does `v3 Specification` connect `v3 Production Architecture` to `v2 Auth & Scaling Roadmap`, `v3 Production Architecture`?**
  _High betweenness centrality (0.228) - this node is a cross-community bridge._
- **Why does `worker() (Goroutine)` connect `Core Engine & Scanner CLI` to `Mock Target Environment`, `Core Engine & Scanner CLI`, `Core Engine & Scanner CLI`, `Core Engine & Scanner CLI`, `Core Engine & Scanner CLI`?**
  _High betweenness centrality (0.160) - this node is a cross-community bridge._
- **Are the 2 inferred relationships involving `v2 Specification` (e.g. with `v1 Specification` and `v3 Specification`) actually correct?**
  _`v2 Specification` has 2 INFERRED edges - model-reasoned connections that need verification._
- **What connects `Package main (The RangeFinder)`, `Evaluator (Interface)`, `Range Traversal` to the rest of the system?**
  _13 weakly-connected nodes found - possible documentation gaps or missing edges._