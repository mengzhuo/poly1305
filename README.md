### Slow Poly1305 in pure Go
--------------------------

This naive poly1305 in Go for understand how poly1305 works instead of speed

Performance
old: golang.org/x/crypto/poly1305
new: github.com/mengzhuo/poly1305
```
name         old time/op    new time/op    delta
64              170ns ± 0%    1495ns ± 0%  +776.83%  (p=0.000 n=10+8)
1K             2.16µs ± 0%   12.52µs ± 0%  +478.67%  (p=0.000 n=10+9)
64Unaligned     169ns ± 0%    1497ns ± 0%  +785.74%  (p=0.000 n=10+10)
1KUnaligned    2.15µs ± 0%   12.52µs ± 0%  +482.38%  (p=0.000 n=10+8)
2M             4.31ms ± 0%   24.11ms ± 0%  +458.95%  (p=0.000 n=9+10)

name         old speed      new speed      delta
64            374MB/s ± 0%    43MB/s ± 0%   -88.56%  (p=0.000 n=10+8)
1K            473MB/s ± 0%    82MB/s ± 0%   -82.71%  (p=0.000 n=10+9)
64Unaligned   377MB/s ± 0%    43MB/s ± 0%   -88.66%  (p=0.000 n=10+10)
1KUnaligned   476MB/s ± 0%    82MB/s ± 0%   -82.82%  (p=0.000 n=10+8)
2M            486MB/s ± 0%    87MB/s ± 0%   -82.11%  (p=0.000 n=9+8)

name         old alloc/op   new alloc/op   delta
64              0.00B        224.00B ± 0%     +Inf%  (p=0.000 n=10+10)
1K              0.00B        224.00B ± 0%     +Inf%  (p=0.000 n=10+10)
64Unaligned     0.00B        224.00B ± 0%     +Inf%  (p=0.000 n=10+10)
1KUnaligned     0.00B        224.00B ± 0%     +Inf%  (p=0.000 n=10+10)
2M              0.00B        224.00B ± 0%     +Inf%  (p=0.000 n=10+10)

name         old allocs/op  new allocs/op  delta
64               0.00           4.00 ± 0%     +Inf%  (p=0.000 n=10+10)
1K               0.00           4.00 ± 0%     +Inf%  (p=0.000 n=10+10)
64Unaligned      0.00           4.00 ± 0%     +Inf%  (p=0.000 n=10+10)
1KUnaligned      0.00           4.00 ± 0%     +Inf%  (p=0.000 n=10+10)
2M               0.00           4.00 ± 0%     +Inf%  (p=0.000 n=10+10)
```
