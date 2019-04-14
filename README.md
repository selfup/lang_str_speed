```bash
git clone https://gist.github.com/selfup/c25874d25c98651b6e6224b608f0c1ff

time elixirc strstr.ex
time go run strstr.go
time ruby strstr.rb
time node strstr.js

rustc strstr.rs && time ./strstr
rustc -C debuginfo=0 -C opt-level=3 strstr.rs && time ./strstr
```
