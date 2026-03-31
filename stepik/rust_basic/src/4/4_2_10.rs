/*
* Вам нужно создать программу, которая считывает текущий баланс (f64) и вносимую сумму (u32), а затем выводит (до 1 знака) итоговый баланс после проведения операции пополнения.
Тестовые данные ✅
Sample Input:
1000.0
500
Sample Output:
1500.0
*/
use std::io;

const ERR_INPUT_READ: &str = "Ошибка чтения";
const ERR_INTEGER_READ: &str = "Не удалось прочитать число";

fn main() {
    let mut str_a = String::new();
    let mut str_b = String::new();
    io::stdin().read_line(&mut str_a).expect(ERR_INPUT_READ);
    io::stdin().read_line(&mut str_b).expect(ERR_INPUT_READ);
    let a: f64 = str_a.trim().parse().expect(ERR_INTEGER_READ);
    let b: u32 = str_b.trim().parse().expect(ERR_INTEGER_READ);
    println!("{:.1}", a + b as f64);
}
