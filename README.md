```bash
# clone this gist
git clone https://gist.github.com/selfup/c25874d25c98651b6e6224b608f0c1ff

# setup (this will make a 206MB file be careful)
./setup.sh

# pick a runtime that you have and run the file
time elixirc strstr.ex
time go run strstr.go
time ruby strstr.rb
time node strstr.js
rustc -C debuginfo=0 -C opt-level=3 strstr.rs && time ./strstr
```
