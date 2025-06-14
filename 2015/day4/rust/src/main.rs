fn hash(text: std::string::String) -> std::string::String {
    format!("{:x}", md5::compute(text))
}

fn check(text: &str) -> bool {
    text.starts_with("00000")
}

fn check2(text: &str) -> bool {
    text.starts_with("000000")
}
fn main() {
    let key = "yzbqklnj";
    let mut index = 0;

    loop {
        let h = hash(format!("{key}{index}"));
        // println!("Testing hash '{}'", h);
        if check2(&h) {
            println!("[!!!] {} at index {}", h, index);
            break;
        } else {
            // println!("[*] {} at index {}", h, index);
        }

        index += 1;
    }
}
