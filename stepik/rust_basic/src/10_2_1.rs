fn main() {
    let mut input = String::new();
    for i in 0..10 {
        std::io::stdin()
            .read_line(&mut input)
            .expect("Ошибка при чтении ввода.");
        input.trim();
        print!("Вы ввели: {input}");

        input.clear();
    }
}
