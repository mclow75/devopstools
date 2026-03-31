/*
* Напишите программу, которая считывает пять строк, сохраняет их в той же последовательности в кортеж и затем выводит получившуюся коллекцию с помощью спецификатора формата :?.
Sample Input:
Упаковка
строк
в
кортеж
!
Sample Output:
("Упаковка\n", "строк\n", "в\n", "кортеж\n", "!\n")
*/
use std::io;

const ERR_INPUT_READ: &str = "Ошибка чтения ввода";

fn main() {
    let mut str1 = String::new();
    let mut str2 = String::new();
    let mut str3 = String::new();
    let mut str4 = String::new();
    let mut str5 = String::new();

    io::stdin().read_line(&mut str1).expect(ERR_INPUT_READ);
    io::stdin().read_line(&mut str2).expect(ERR_INPUT_READ);
    io::stdin().read_line(&mut str3).expect(ERR_INPUT_READ);
    io::stdin().read_line(&mut str4).expect(ERR_INPUT_READ);
    io::stdin().read_line(&mut str5).expect(ERR_INPUT_READ);

    let tup = (str1, str2, str3, str4, str5);
    println!("{:?}", tup);
}
