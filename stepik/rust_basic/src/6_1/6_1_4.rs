/*
* Напишите программу, которая считывает трехзначное положительное число (u16) и выводит сумму его цифр.
Чтобы найти сумму цифр трехзначного числа, можно разложить число на отдельные цифры и затем сложить их.
Тестовые данные ✅
Sample Input:
123
Sample Output:
6
*/
use std::io;

const ERR_INPUT_READ: &str = "Ошибка чтения ввода";
const ERR_INTEGER_READ: &str = "Не удалось прочитать число";

fn main() {
    let mut str = String::new();
    io::stdin().read_line(&mut str).expect(ERR_INPUT_READ);
    let mut num: u16 = str.trim().parse().expect(ERR_INTEGER_READ);
    let mut sum: u16 = 0;
    for _ in 0..3 {
        sum += num % 10;
        num = num / 10;
    }
    println!("{}", sum);
}
