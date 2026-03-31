/*
 * В редакторе кода представлен массив значений. Дополните код и считайте два целых числа. Далее поменяйте местами значения по этим индексам, а затем выведите (до 9 знаков) получившийся массив, как показано в примере. Гарантируется, что считанные числа (индексы) заданы корректно!

 Тестовые данные ✅
 Sample Input:
 0
 6
 Sample Output:
 [0.000051789, 11.100000000, 2.000000000, -7.123000000, 0.125000000, 0.000000000, -621.500000000]
*/
use std::io;

const ERR_INPUT_READ: &str = "Ошибка чтения ввода";
const ERR_INTEGER_READ: &str = "Не удалось прочитать число";

fn main() {
    let mut arr = [-621.5, 11.1, 2.0, -7.123, 0.125, 0.0, 0.000051789];
    let mut str1 = String::new();
    let mut str2 = String::new();

    io::stdin().read_line(&mut str1).expect(ERR_INPUT_READ);
    io::stdin().read_line(&mut str2).expect(ERR_INPUT_READ);
    let idx1: usize = str1.trim().parse().expect(ERR_INTEGER_READ);
    let idx2: usize = str2.trim().parse().expect(ERR_INTEGER_READ);
    let tmp = arr[idx1];
    arr[idx1] = arr[idx2];
    arr[idx2] = tmp;
    println!("{:.9?}", arr);
}
