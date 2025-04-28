

// // 01-min_and_max
// pub fn min_and_max(nb_1: i32, nb_2: i32, nb_3: i32) -> (i32, i32) {
//     let min = if nb_1 < nb_2 {
//         if nb_1 < nb_3 {
//             nb_1
//         } else {
//             nb_3
//         }
//     } else {
//         if nb_2 < nb_3 {
//             nb_2
//         } else {
//             nb_3
//         }
//     };
//     let max = if nb_1 > nb_2 {
//         if nb_1 > nb_3 {
//             nb_1
//         } else {
//             nb_3
//         }
//     } else {
//         if nb_2 > nb_3 {
//             nb_2
//         } else {
//             nb_3
//         }
//     };
//     (min, max)
    // use std::cmp::{min, max};
    // (min(nb_1, min(nb_2, nb_3)), max(nb_1, max(nb_2, nb_3)))
// }

// // 02-count_factorial_steps
// pub fn count_factorial_steps(factorial: u64) -> u64 {
//     if factorial == 0 || factorial == 1 {
//         return 1;
//     }
//     let mut count = 1;
//     let mut fact = 1;
//     for i in 2..=factorial {
//         fact *= i;
//         count += 1;
//         if fact == factorial {
//             return count;
//         } else if fact > factorial {
//             return 0;
//         }
//     }
//     println!("The value of fact is {}", fact);
//     0
//     if factorial <= 1 {
//         return 1;
//     }
//     let mut fact = 1;
//     let mut i = 1;
    
//     while fact < factorial {
//         i += 1;
//         fact *= i;
//     }
    
//     if fact == factorial {
//         i
//     } else {
//         0
//     }
// }

// // 3-matrix_multiplication
// #[derive(Debug, PartialEq, Eq)]
// pub struct Matrix( pub (i32, i32), pub (i32, i32));

// pub fn multiply(m: Matrix, multiplier: i32) -> Matrix {
//     Matrix(
//         (multiplier * m.0.0, multiplier * m.0.1),
//         (multiplier * m.1.0, multiplier * m.1.1),
//     )
// }

// // 4-reverse_it
// pub fn reverse_it(v: i32) -> String {
//     let mut result = String::new();
//     let mut value = v;
//     if v < 0 {
//         result.push_str("-");
//         value = value.abs();
//     }
//     let value_str = value.to_string();
//     let value_str_rev: String = value_str.chars().rev().collect();
//     result.push_str(&value_str_rev);
//     result.push_str(&value_str);
//     result

//     let value = v.abs().to_string();
//     let reversed: String = value.chars().rev().collect();

//     if v < 0 {
//         format!("-{}{}", reversed, value)
//     } else {
//         format!("{}{}", reversed, value)
//     }
// }

// // 5-counting_words
// use std::collections::HashMap;

// pub fn counting_words(words: &str) -> HashMap<String, u32> {
//     // let mut result = HashMap::new();
//     // let cleaned: String = words
//     //     .chars()
//     //     .filter(|c| c.is_alphanumeric() || c.is_whitespace() || *c == '\'')
//     //     .collect();

//     // for word in cleaned.split_whitespace() {
//     //     let normalized_word = word.to_lowercase();
//     //     *result.entry(normalized_word).or_insert(0) += 1;
//     // }
//     // result
//     // let mut result = HashMap::new();
//     // for word in words.split_whitespace() {
//     //     // Remove leading and trailing punctuation (except valid inner apostrophes)
//     //     let trimmed = word
//     //         .trim_matches(|c: char| !c.is_alphanumeric())
//     //         .to_lowercase();

//     //     // Check if the word is now empty after trimming
//     //     if trimmed.is_empty() {
//     //         continue;
//     //     }

//     //     // Ignore words with more than one apostrophe
//     //     let apostrophe_count = trimmed.chars().filter(|&c| c == '\'').count();
//     //     if apostrophe_count > 1 {
//     //         continue;
//     //     }

//     //     *result.entry(trimmed).or_insert(0) += 1;
//     // }
//     // result
// }

// // 6-smallest

// use std::{collections::HashMap, i32};

// pub fn smallest(h: HashMap<&str, i32>) -> i32 {
//     let mut smallest = i32::MAX ;
//     for (_, num) in h {
//         if num < smallest {
//             smallest = num;
//         }
//     }
//     smallest

//     h.into_values().min().unwrap_or(i32::MAX)
// }

// // 7-modify_letter

// pub fn remove_letter_sensitive(s: &str, letter: char) -> String {
//     s.chars().filter(|x| x != &letter).collect()
// }

// pub fn remove_letter_insensitive(s: &str, letter: char) -> String {
//     let letter = letter.to_ascii_lowercase();
//     s.chars()
//         .filter(|x| x.to_ascii_lowercase() != letter)
//         .collect()
// }

// pub fn swap_letter_case(s: &str, letter: char) -> String {
//     let lower = letter.to_ascii_lowercase();
//     let upper = letter.to_ascii_uppercase();
//     s.chars()
//         .map(|x| {
//             if x == lower {
//                 upper
//             } else if x == upper {
//                 lower
//             } else {
//                 x
//             }
//         })
//         .collect()
// }

// // 8-partial_sums
// pub fn parts_sums(arr: &[u64]) -> Vec<u64> {
//     let mut result = Vec::with_capacity(arr.len() + 1);
//     // let vector = Vec::from(arr);
//     let mut sum: u64 = arr.iter().sum();
//     result.push(sum);
//     for i in arr.iter().rev() {
//         sum -= i;
//         result.push(sum);
//     }
//     result
// }

// // 9-inv_pyramid
// pub fn inv_pyramid(v: String, i: u32) -> Vec<String> {
//     let mut result: Vec<String> = Vec::new();
//     for n in 1..=i {
//         let spaces = " ".repeat(n as usize);
//         let symbols = v.repeat(n as usize);
//         let part = format!("{}{}", spaces, symbols);
//         result.push(part);
//     }
//     for n in (1..i).rev() {
//         let spaces = " ".repeat(n as usize);
//         let symbols = v.repeat(n as usize);
//         let part = format!("{}{}", spaces, symbols);
//         result.push(part);
//     }
//     result
// }

// // 10-prev_prime
// pub fn prev_prime(nbr: u64) -> u64 {
//     if nbr < 2 {
//         return 2; // optimus prime
//     }
//     let mut result = nbr;
//     while !is_prime(result) {
//         result -= 1;
//     }
//     result
// }

// // 11-next_prime
// pub fn next_prime(nbr: u64) -> u64 {
//     if nbr < 2 {
//         return 2; // optimus prime
//     }
//     let mut result = nbr;
//     while !is_prime(result) {
//         result += 1;
//     }
//     result
// }

// fn is_prime(nbr: u64) -> bool {
//     if nbr == 2 || nbr == 3 {
//         return true;
//     }
//     if nbr % 2 == 0 {
//         return false;
//     }
//     for i in (3..=((nbr as f64).sqrt() as u64)).step_by(2) {
//         if nbr % i == 0 {
//             return false;
//         }
//     }
//     true

// nbr == 2 || (nbr > 2 && (2..=(nbr as f64).sqrt() as u64).all(|i| nbr % i != 0))
// }

// // 12-profanity_filter
// pub struct Message {
//     content: String,
//     user: String,
// }

// impl Message {
//     pub fn new(ms: String, u: String) -> Message {
//         Message { content: ms, user: u }
//     }

//     pub fn send_ms(&self) -> Option<&str> {
//         if self.content.is_empty() || self.content.contains("stupid") {
//             return None;
//         }
//         Some(&self.content)
//     }
// }

// pub fn check_ms(message: &Message) -> (bool, &str) {
//     match message.send_ms() {
//         None => (false, "ERROR: illegal"),
//         Some(msg) => (true, msg),
//     }
// }

// // 13-prime_checker
// #[derive(PartialEq, Eq, Debug)]
// pub enum PrimeErr {
//     Even,
//     Divider(u32),
// }

// pub fn prime_checker(nb: u32) -> Option<Result<u32, PrimeErr>> {
//     if nb < 2 {
//         return None
//     }
//     if nb == 2 || nb == 3 {
//         return Some(Ok(nb))
//     }
//     if nb % 2 == 0 {
//         return Some(Err(PrimeErr::Even));
//     }
    
//     for divider in (3..=(nb as f64).sqrt() as u32).step_by(2) {
//         if nb % divider == 0 {
//             return Some(Err(PrimeErr::Divider(divider)))
//         }
//     } 
//     Some(Ok(nb))
// }

// pub fn insertion_sort(slice: &mut [i32], step: usize) {
//     // let len = slice.len();
//     // let step = step.min(len.saturating_sub(1));
//     for current in 1..step {
//         let key = slice[current];
//         let mut previous = current;

//         while previous > 0 && slice[previous - 1] > key {
//             slice[previous] = slice[previous - 1];
//             previous -= 1;
//         }
//         slice[previous] = key
//     }
// }

// pub fn rpn(expr: &str) {
//     let mut stack: Vec<i64> = Vec::new();
    
//     for token in expr.split_whitespace() {
//         match token {
//             "+" | "-" | "*" | "/" | "%" => {
//                 if stack.len() < 2 {
//                     println!("Error");
//                     return;
//                 }
                
//                 let b = stack.pop().unwrap();
//                 let a = stack.pop().unwrap();
                
//                 let result = match token {
//                     "+" => a + b,
//                     "-" => a - b,
//                     "*" => a * b,
//                     "/" => {
//                         if b == 0 {
//                             println!("Error");
//                             return;
//                         }
//                         a / b
//                     },
//                     "%" => {
//                         if b == 0 {
//                             println!("Error");
//                             return;
//                         }
//                         a % b
//                     },
//                     _ => unreachable!(),
//                 };
                
//                 stack.push(result);
//             },
//             _ => {
//                 // Try to parse the token as a number
//                 match token.parse::<i64>() {
//                     Ok(num) => stack.push(num),
//                     Err(_) => {
//                         println!("Error");
//                         return;
//                     }
//                 }
//             }
//         }
//     }
    
//     // After processing all tokens, there should be exactly one value on the stack
//     if stack.len() != 1 {
//         println!("Error");
//         return;
//     }
    
//     println!("{}", stack[0]);
// }

// pub fn rot21(input: &str) -> String {
//     input
//         .chars()
//         .map(|c| {
//             if c.is_ascii_lowercase() {
//                 // For lowercase letters (a-z)
//                 let base = b'a';
//                 let rotated = (((c as u8 - base) + 21) % 26) + base;
//                 rotated as char
//             } else if c.is_ascii_uppercase() {
//                 // For uppercase letters (A-Z)
//                 let base = b'A';
//                 let rotated = (((c as u8 - base) + 21) % 26) + base;
//                 rotated as char
//             } else {
//                 // For non-alphabetic characters
//                 c
//             }
//         })
//         .collect()
// }

// pub fn matrix_determinant(matrix: [[isize; 3]; 3]) -> isize {
//     matrix[0][0] * (matrix[1][1] * matrix[2][2] - matrix[2][1] * matrix[1][2])
//     - matrix[0][1] * (matrix[1][0] * matrix[2][2] - matrix[2][0] * matrix[1][2])
//     + matrix[0][2] * (matrix[1][0] * matrix[2][1] - matrix[2][0] * matrix[1][1])
// }

// #[derive(Debug, PartialEq, Eq)]
// pub struct OfficeWorker {
//     name: String,
//     age: u32,
//     role: WorkerRole,
// }

// #[derive(Debug, PartialEq, Eq)]
// pub enum WorkerRole {
//     Admin,
//     User,
//     Guest,
// }

// impl From<&str> for OfficeWorker {
//     fn from(s: &str) -> Self {
//         let parts: Vec<&str> = s.split(',').collect();
//         let name = parts[0].to_string();
//         let age: u32 = parts[1].parse().unwrap();
//         let role = WorkerRole::from(parts[2]);
//         OfficeWorker {
//             name,
//             age,
//             role,
//         }
//     }
// }

// impl From<&str> for WorkerRole {
//     fn from(s: &str) -> Self {
//         match s {
//             "admin" => WorkerRole::Admin,
//             "user" => WorkerRole::User,
//             "guest" => WorkerRole::Guest,
//             _ => panic!("Invalid Role!")
//         }
//     }
// }
