fn main() {
    let mut n = input_int() as u16;
    for s in ["Последняя", "Третья", "Вторая", "Первая"] {
        println!("{} цифра: {}", s, n % 10);
        n /= 10;
    }
}

const ERR_INPUT_READ: &str = "Ошибка чтения ввода";
const ERR_INTEGER_READ: &str = "Не удалось преобразовать стоку в число";

fn input_int() -> i32 {
    let mut buf = String::new();
    std::io::stdin()
        .read_line(&mut buf)
        .expect(ERR_INPUT_READ);
    return buf.trim()
        .parse::<i32>()
        .expect(ERR_INTEGER_READ);
}

fn input_double() -> f64 {
    let mut buf = String::new();
    std::io::stdin()
        .read_line(&mut buf)
        .expect(ERR_INPUT_READ);
    return buf.trim()
        .parse::<f64>()
        .expect(ERR_INTEGER_READ);
}

// Множественный ввод с консоли
let mut input = String::new();
for i in 0..3 {
    std::io::stdin()
        .read_line(&mut input)
        .expect("Ошибка при чтении ввода.");
    println!("Содержимое input: {:#?}", input);

    let num: i32 = input.trim().parse().expect("Введите целое число.");
    println!("Вы ввели число: {num}");

    input.clear();
}
