# minhash
b-bit minhash (b=1)

## Feature
- Use MurmurHash3 (128-bit)
- The number of hash functions is 128

## Require
- Go 1.9 or above
- spaolacci/murmur3 [https://github.com/spaolacci/murmur3]

## Install
```
go get github.com/dono/minhash
```

## Usage
```
d1 := []string{"foo", "bar", "baz", "qux", "quux"}
d2 := []string{"foo", "bar", "baz", "corge", "grault"}

s1 := minhash.Sketch(d1)
s2 := minhash.Sketch(d2)

jaccard := minhash.Jaccard(s1, s2)
fmt.Println(jaccard)
```

## License
MIT