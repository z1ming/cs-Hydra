# 浅谈 LISP

## 前言

今天学习了很有意思的一个编程语言，叫做 LISP[1]，它的历史悠久，由约翰·麦卡锡在 1958 年发明，到现在仍然被广泛使用。LISP 有以下特点：

- 只用圆括号，没有大括号，分号等
- 浓厚的递归思想，对学习递归有好处

LISP 目前有两个实现的分支，一个叫 Common LISP，一个叫 Scheme，下文我们会谈到 Scheme。

## Java 实现 LISP

### 定义

LISP 的数据结构是一个列表，类似链表，`ar` 是头部，`dr` 是后续部分。每一个元素我们定义为`Cell`，如下：

```java
import java.util.Objects;

public class Cell {
    public static final Cell nil = new Cell(null, null);
    protected Object ar;
    protected Object dr;

    protected Cell(Object ar, Object dr) {
        this.ar = ar;
        this.dr = dr;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Cell cell = (Cell) o;
        return Objects.equals(ar,cell.ar) && Objects.equals(dr, cell.dr);
    }

    @Override
    public int hashCode() {
        return Objects.hash(ar, dr);
    }

    private String tailString() {
        if (this == nil) {
            return "";
        } else if (dr instanceof Cell) {
            String tail = ((Cell) dr).tailString();
            return ar.toString() + (tail.isEmpty() ? "" : " ") + tail;
        } else {
            return ar.toString() + " . " + dr.toString();
        }
    }

    @Override
    public String toString() {
        return "(" + tailString() + ")";
    }
}
```

### 常用内置函数实现

LISP 有一些常用的内置函数，如

- `cons`: 组装一个新的 Cell
- `list`: 返回 Cell
- `car`: 返回头节点
- `cdr`: 返回后续节点


```java
public class Lisp {
    public static Cell cons(Object ar, Object dr) {
        return new Cell(ar, dr);
    }

    public static Cell _list(int index, Object... elements) {
        return index == elements.length ? nil : cons(elements[index], _list(index + 1, elements));
    }

    public static Cell list(Object... elements) {
        return _list(0, elements);
    }

    public static Object car(Object cell) {
        return ((Cell) cell).ar;
    }

    public static Object cdr(Object cell) {
        return ((Cell) cell).dr;
    }
}
```

除此之外，还有一些其他内置函数，都可以用递归的思想实现。以下代码实现了 `member`，`union`, `intersection`方法：

- `member`：判断元素是否存在列表中，如果存在，返回该元素引用
- `union`：求两个列表的并集
- `intersection`：求两个列表的交集

```java
public class Main {
    public static Object member(Object key, Object list) {
        if (list == nil) {
            return nil;
        }
        if (car(list).equals(key)) {
            return list;
        }
        return member(key, cdr(list));
    }


    public static Object union(Object x, Object y) {
        if (x == nil) {
            return y;
        }
        if (member(car(x), y) == nil) {
            return union(cdr(x), cons(car(x), y));
        }
        return union(cdr(x), y);
    }

    public static Object intersection(Object x, Object y) {
        return helper(nil, x, y);
    }

    private static Object helper(Cell ans, Object x, Object y) {
        if (x == nil) {
            return ans;
        }
        if (member(car(x), y) != nil) {
            return helper(cons(car(x), ans), cdr(x), y);
        }
        return helper(ans, cdr(x), y);
    }
}
```

## Scheme 

Scheme 是 LISP 具体实现的一个分支，官方文档[2]有其具体的描述。它有自己的 IDE 下载地址[3]，叫做 racket，界面也很简洁。

一个很简单的例子，在新建 racket 新建一个 `solution.rkt`，我们用 Scheme 实现一下 union 和 intersection 方法，记得在文件开头标注 Module Language：`#lang scheme`

```
#lang scheme
(define (union x y)
  (if (null? x)
      y
      (if (not (member (car x) y))
          (cons (car x) (union (cdr x) y))
          (union (cdr x) y))))

(define (intersection x y)
  (helper '() x y))

(define (helper ans x y)
  (if (null? x)
      ans
      (if (member (car x) y)
          (helper (cons (car x) ans) (cdr x) y)
          (helper ans (cdr x) y))))

(union '(3 4 5) '(1 2 3))        // 输出 (4 5 1 2 3)
(union '() '(1 2 3))             // 输出 (1 2 3)
(union '(3 4 5) '(4 3))          // 输出 (5 4 3)
(intersection '(3 4 5) '(1 2 3)) // 输出 (3)
(intersection '() '(1 2 3))      // 输出 ()
(intersection '(3 4 5) '(4 3))   // 输出 (4 3)
(car '(1 2 3 4 5))               // 输出 1
(cdr '(1 2 3 4 5))               // 输出 (2 3 4 5)
(append '(1) '(2 3 4 5))         // 输出 (1 2 3 4 5)
```

### 使用 scheme 实现 padovan 函数

padovan 函数[4]类似 Fibonacci 数列，两者公式有一些不同。Fibonacci 数列为 ：

- P(0)=P(1)=1
- P(n)=P(n-1)+P(n-2)

padovan 数列为：

- P(0)=P(1)=P(2)=1
- P(n)=P(n-2)+P(n-3)

```
#lang scheme
(define (padovan x)  
  (cond  
    [(or (= x 0) (= x 1) (= x 2))  
     1]  
    [else  
     (+ (padovan (- x 2)) (padovan (- x 3)))]))

(padovan 0)
(padovan 1)
(padovan 2)
(padovan 3)
(padovan 4)
(padovan 5)
(padovan 6)
(padovan 10)
```

输出为：

```
1
1
1
2
2
3
4
12
```

## 总结

自己尝试实现 LISP 的一些方法，有助于锻炼递归的思想，还能掌握一门编程语言，这个感觉简直太棒了。

## 参考

1. 中文维基百科: https://zh.wikipedia.org/zh-hans/LISP
2. Scheme 官方文档：https://docs.racket-lang.org/r5rs/r5rs-std/r5rs-Z-H-9.html
3. racket 下载地址：https://download.racket-lang.org/
4. padovan sequence 维基百科: https://en.wikipedia.org/wiki/Padovan_sequence
