```bash
time elixirc strstr.ex
time go run strstr.go
time ruby strstr.rb
time node strstr.js

rustc strstr.rs && time ./strstr
rustc -C debuginfo=0 -C opt-level=3 strstr.rs && time ./strstr
```
