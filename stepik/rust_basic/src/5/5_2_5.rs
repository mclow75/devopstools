/*
* Дополните код и считайте целое число (usize), а затем выведите элемент массива по этому индексу, как показано в примере. Гарантируется, что считанное число не превышает длины массива!
* Sample Input:
0
Sample Output:
-1
*/
use std::io;

fn main() {
    let arr = [-1, 0, 1, 2, 30, 4, 500];

    const ERR_INPUT_READ: &str = "Ошибка чтения";
    const ERR_INTEGER_READ: &str = "Не удалось прочитать число";
    let mut str_n = String::new();
    io::stdin().read_line(&mut str_n).expect(ERR_INPUT_READ);
    let n: usize = str_n.trim().parse().expect(ERR_INTEGER_READ);
    println!("{}", arr[n]);
}
