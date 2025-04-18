# Algorithms
Мои алгоритмы

### Аггрегация массивов - группировка списков на примере ip адресов с сопоставимыми доменами  
```
- 0 &{0 13 1 Обнаружен связанный ip адрес 176.0.1.3 для доменов doc.dom.ru, factor.dom.ru 2025-04-18 00:52:12}
- 1 &{0 13 1 Обнаружены связанные ip адреса 17.0.3.5, 176.0.1.3, 17.0.3.2 для домена fox.dom.ru 2025-04-18 00:52:12}
- 2 &{0 13 1 Обнаружены связанные ip адреса 192.168.3.2, 14.168.3.2 для доменов calc.domain.ru, fox.mozilla.ru 2025-04-18 00:52:12}
- 3 &{0 13 2 Обнаружены новые связанные поддомены fox.dom.ru, factor.dom.ru, doc.dom.ru для домена dom.ru 0001-01-01 00:00:00}
- 4 &{0 13 2 Обнаружены новые связанные поддомены fox.mozilla.ru, factor.mozilla.ru для домена mozilla.ru 0001-01-01 00:00:00}
- 5 &{0 13 2 Обнаружены новые связанные поддомены calc.domain.ru, factor.domain.ru для домена domain.ru 0001-01-01 00:00:00}
```

### Простая сортировка чисел
```
cpu: 11th Gen Intel(R) Core(TM) i7-11800H @ 2.30GHz
BenchmarkBubles     	1000000000	         0.02021 ns/op
BenchmarkBucket     	1000000000	         0.01458 ns/op
BenchmarkMy15minutesSort     	1000000000	         0.03899 ns/op
BenchmarkBublesBig        1	1421773200 ns/op
BenchmarkBucketBig   	1000000000	         0.07106 ns/op
BenchmarkMy15minutesSortBig    1	2687106700 ns/op
```
