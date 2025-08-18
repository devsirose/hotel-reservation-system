# simple-bank Concurrency Learning

## 1. Goals
- Preserve money: total Σ(balance) unchanged after all scenarios
- No deadlock, no lost update
- Idempotent transactions with proper timeout/cancellation
- Metrics: p50/p95 latency, error rate, deadlock count, retry count

## 2. Mandatory Exercises

### A. TransferTx under concurrent load
**Problem**  
Implement `TransferTx(ctx, from, to, amount)` safely under concurrency. Avoid deadlock on account updates.

**Pass Criteria**
- Run N=1000 concurrent requests
- Σ(balance) equals initial total
- No panic, only retryable errors handled automatically
- No deadlock

### B. Reproduce and remove deadlock
**Problem**  
Design test with A→B and B→A transfers to trigger deadlock, then eliminate via correct locking/isolation.

**Pass Criteria**
- Deadlock eliminated
- If high isolation: only “serialization failure” remains, handled with retries

### C. Idempotency for CreateTransfer
**Problem**  
Enforce idempotency with request identifier.

**Pass Criteria**
- 1000 duplicate requests → one transfer record
- Duplicates return same result, no double effect

### D. Context timeout and cancel
**Problem**  
APIs must respect `context.Context` timeout.

**Pass Criteria**
- Short timeout + forced delay → `deadline exceeded`
- No hanging transaction or partial change

### E. Retry for transient errors
**Problem**  
Retry transient errors, not business logic errors.

**Pass Criteria**
- ≥99% success with 3 retries + backoff
- No retry on business violations

### F. Rate limiting per owner
**Problem**  
Limit requests per owner.

**Pass Criteria**
- Config 10 req/sec, send 20 req/sec
- Accept ~10, reject rest

### G. Outbox event “TransferCreated”
**Problem**  
Guarantee event publishing via outbox pattern.

**Pass Criteria**
- 100 transfers created offline
- On worker start: all events published once, no loss/duplication

### H. Observability and metrics
**Problem**  
Expose metrics for latency, deadlocks, retries.

**Pass Criteria**
- Histogram for latency
- Counters for deadlocks/retries
- For N=1000 run: deadlock=0, report p50/p95

### I. DB connection pool optimization
**Problem**  
Tune connection pool for throughput and stability.

**Pass Criteria**
- Higher throughput than default at N=1000
- No “too many connections” error

## 3. Concurrency Testing
**Problem**  
Test property: no money lost under random concurrent transfers.

**Pass Criteria**
- Σ(balance) unchanged
- `go test -race` reports no race
- Benchmark `TransferTx` runs ≥10s

## 4. Diagnostics
**Problem**  
Provide visibility into DB locks, conflicts, slow queries.

**Pass Criteria**
- Have hooks/procedure to log lock waits, conflicts, slow queries (no commands in README)

## 5. Constraints
- Do not commit secrets or sensitive data
- No solutions, code, or configs in README
- Tests must run via standard `go test`  
