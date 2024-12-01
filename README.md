# MicroLang

MicroLang is a programming language designed for [micromashinki/noVax](https://github.com/micromashinki/noVax) and partially compatible with `VAX`.

## Example

```antlrv4-tool
/* /* bober
*/
x = 2;
x = x + ( 1 + 2 * 2);
x;
4;
while (x<10){
    x = x + 1;
}

if (x>100){
    x = 0;
}
x;
```
Output:
```
7
4
10
```