# Column vs Row format

| Time | Value |
| :--- | :---- |
| 123  |  2.0  |
| 124  |  1.2  |
| 125  |  2.1  |

````js
// row format
[
  {t: 123, v: 2.0},
  {t: 124, v: 1.2},
  {t: 125, v: 2.1}
]
// columnar format
{
  t: [123, 124, 125],
  v: [2.0, 1.2, 2.1]
}
````
