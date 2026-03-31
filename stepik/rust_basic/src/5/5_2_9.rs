/*
* В редакторе кода представлен массив целых чисел. Дополните код и считайте целое число, а затем выведите сумму, разность, произведение предыдущего и следующего элемента по этому индексу, как показано в примере.

Гарантируется, что считанное число всегда больше нуля и не превышает индекса предпоследнего элемента!
Sample Input:
1
Sample Output:
3
-13
-40
*/

use std::io;

const ERR_INPUT_READ: &str = "Ошибка чтения ввода";
const ERR_INTEGER_READ: &str = "Не удалось прочитать число";

fn main() {
    let arr = [-5, 1, 8, 2, 30, 4000, 500000];

    let mut str = String::new();
    io::stdin().read_line(&mut str).expect(ERR_INPUT_READ);
    let i: usize = str.trim().parse().expect(ERR_INTEGER_READ);
    let prev = arr[i - 1];
    let next = arr[i + 1];
    println!("{}\n{}\n{}\n", prev + next, prev - next, prev * next);
}
