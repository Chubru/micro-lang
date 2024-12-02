# MicroLang

MicroLang is a programming language designed compatible with three-address code.

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

Three-address code:
```
x = 2
t1 = 2 * 2
t2 = 1 + t1
t3 = x + t2
x = t3
print x
print 4
L1:
t4 = x < 10
if False t4 goto L2
t5 = x + 1
x = t5
goto L1
L2:
print x
```