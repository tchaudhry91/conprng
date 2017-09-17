## Concurrent Prime Number Generator in Go written on the pipe-filter pattern

This is a learning project

tanmay.chaudhry@gmail.com


Benchmarking via https://play.golang.org/p/Id4PMh94Ik on a 4 core machine:

```
tchaudhr:prime-generator/ $ ./prime-generator 10000 1                                 
[104677 104681 104683 104693 104701 104707 104711 104717 104723 104729]
Time: 15.68725599s
tchaudhr:prime-generator/ $ ./prime-generator 10000 2         
[104677 104681 104683 104693 104701 104707 104711 104717 104723 104729]
Time: 7.901663202s
tchaudhr:prime-generator/ $ ./prime-generator 10000 3                    
[104677 104681 104683 104693 104701 104707 104711 104717 104723 104729]
Time: 5.453621754s
tchaudhr:prime-generator/ $ ./prime-generator 10000 4                     
[104677 104681 104683 104693 104701 104707 104711 104717 104723 104729]
Time: 3.300964504s
tchaudhr:prime-generator/ $ ./prime-generator 10000 5
[104677 104681 104683 104693 104701 104707 104711 104717 104723 104729]
Time: 3.779781834s
tchaudhr:prime-generator/ $ ./prime-generator 10000 6
[104677 104681 104683 104693 104701 104707 104711 104717 104723 104729]
Time: 3.44113436s
```
