/*
* Напишите программу, которая считывает четырехзначное число (u16) и выводит его цифры с конца, как показано в примере.

Тестовые данные ✅
Sample Input:
1234
Sample Output:
Последняя цифра: 4
Третья цифра: 3
Вторая цифра: 2
Первая цифра: 1
*/
use std::io;

const ERR_INPUT_READ: &str = "Ошибка чтения ввода";
const ERR_INTEGER_READ: &str = "Не удалось прочитать число";

fn main() {
    let mut str = String::new();
    io::stdin().read_line(&mut str).expect(ERR_INPUT_READ);
    let mut num: u16 = str.trim().parse().expect(ERR_INTEGER_READ);
    let pos = vec!["Последняя", "Третья", "Вторая", "Первая"];
    for (_, v) in pos.iter().enumerate() {
        println!("{} цифра: {}", v, num % 10);
        num = num / 10;
    }
}
