# 题目 1：计算数组中元素的和

编写一个函数 `sumArray(arr []int) int`，接收一个整数数组作为参数，返回数组中所有元素的和。在主函数中定义一个数组并调用这个函数，输出结果。

示例：

```css
输入：[1, 2, 3, 4, 5]
输出：数组元素的和是 15
```

# 题目 2：查找切片中的最大值

编写一个函数 `findMax(slice []int) int`，接收一个整数切片，返回其中的最大值。处理空切片的情况，返回一个错误提示。

示例：

```css
输入：[7, 2, 9, 4]
输出：切片中的最大值是 9
```

# 题目 3：模拟银行账户

定义一个结构体 `Account`，具有 `balance` 属性。编写方法：

`Deposit(amount float64)`：存款。
`Withdraw(amount float64) error`：取款，如果取款金额大于余额，返回错误。
在主函数中模拟一些存款和取款操作，并处理可能的错误。

示例：

```rust
初始余额：1000
存款：500
取款：2000 -> 错误：余额不足
```